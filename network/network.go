package network

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/liuqi0826/seven/events"
	"github.com/liuqi0826/seven/utils"
)

const MAX_NETWORK_PACKAGE_LENGTH = 2048
const DELAY = time.Millisecond * 10

type DataStruct struct {
	Title     uint16
	Label     uint16
	SendIndex uint16
	ReceIndex uint16
	Data      []byte
}

//==================== local ====================

type ChanConnection struct {
	sync.Mutex

	alive   bool
	channel chan []byte
	binding *ChanConnection
}

func (this *ChanConnection) ChanConnection(connect *ChanConnection) {
	this.alive = true
	this.channel = make(chan []byte, 32)
	if connect != nil {
		this.binding = connect
		connect.binding = this
	}
}
func (this *ChanConnection) Read(buf []byte) (int, error) {
	var err error
	if this != nil && this.alive {
		for this.binding == nil {
			time.Sleep(time.Millisecond * 100)
		}
		if msg, ok := <-this.binding.channel; ok {
			length := len(msg)
			stream := bytes.NewBuffer([]byte{})
			binary.Write(stream, binary.BigEndian, length)
			binary.Write(stream, binary.BigEndian, msg)
			copy(buf, stream.Bytes())
			return length, err
		} else {
			err = errors.New("Bingding connect channel closed")
		}
	} else {
		err = errors.New("ChanConnection closed")
	}
	return 0, err
}
func (this *ChanConnection) Write(msg []byte) (int, error) {
	var err error
	if this != nil && this.alive {
		if len(this.channel) < cap(this.channel) {
			this.channel <- msg
		} else {
			err = errors.New("ChanConnection connect is full.")
			return 0, err
		}
	} else {
		err = errors.New("ChanConnection channel closed")
		return 0, err
	}
	return len(msg), err
}
func (this *ChanConnection) Close() error {
	var err error
	if this != nil && this.alive {
		this.Lock()
		defer this.Unlock()
		this.alive = false
		this.binding = nil
		close(this.channel)
	} else {
		err = errors.New("ChanConnection channel closed")
		return err
	}
	return err
}

//==================== websocket ====================

type WSConnction struct {
	sync.Mutex

	alive    bool
	connect  *websocket.Conn
	readBuf  chan []byte
	writeBuf chan []byte
}

func (this *WSConnction) WSConnction(connect *websocket.Conn) {
	this.alive = true
	this.connect = connect
	this.readBuf = make(chan []byte, 32)
	this.writeBuf = make(chan []byte, 32)
	fmt.Println("Websocket connect from: " + fmt.Sprintf("%s", this.connect.RemoteAddr()))

	go this.readListen()
	go this.writeListen()
}
func (this *WSConnction) Read(buf []byte) (int, error) {
	var err error
	if this != nil && this.alive {
		if msg, ok := <-this.readBuf; ok {
			length := len(msg)
			if length <= MAX_NETWORK_PACKAGE_LENGTH {
				copy(buf, msg)
				return length, err
			} else {
				err = errors.New("Data package length is out of range.")
			}
		} else {
			err = this.Close()
		}
	} else {
		err = errors.New("Websocket connection closed")
	}
	return 0, err
}
func (this *WSConnction) Write(msg []byte) (int, error) {
	var err error
	if this != nil && this.alive {
		if len(this.writeBuf) < cap(this.writeBuf) {
			this.writeBuf <- msg
		} else {
			err = errors.New("Websocket connect is full.")
			return 0, err
		}
	} else {
		err = errors.New("Websocket connection closed.")
		return 0, err
	}
	return len(msg), err
}
func (this *WSConnction) Close() error {
	var err error
	if this != nil && this.alive {
		this.Lock()
		defer this.Unlock()
		this.alive = false
		close(this.readBuf)
		close(this.writeBuf)
		fmt.Println("Websocket close: " + fmt.Sprintf("%s", this.connect.RemoteAddr()))
		err = this.connect.Close()
	} else {
		err = errors.New("Websocket has been closed.")
	}
	return err
}
func (this *WSConnction) readListen() {
	for this != nil && this.alive {
		_, msg, err := this.connect.ReadMessage()
		if err == nil {
			this.readBuf <- msg
		} else {
			err = this.Close()
		}
	}
}
func (this *WSConnction) writeListen() {
	for this != nil && this.alive {
		if message, ok := <-this.writeBuf; ok {
			err := this.connect.WriteMessage(websocket.BinaryMessage, message)
			if err != nil {
				this.Close()
			}
		} else {
			this.Close()
		}
	}
}

//==================== tcp socket ====================

type TCPConnection struct {
	sync.Mutex

	alive      bool
	headRead   bool
	packageLen uint32
	buffer     []byte
	connect    net.TCPConn
	readBuf    chan []byte
	writeBuf   chan []byte
}

func (this *TCPConnection) TCPConnection(connect net.TCPConn) {
	this.alive = true
	this.buffer = make([]byte, 2048)
	this.connect = connect
	this.readBuf = make(chan []byte, 32)
	this.writeBuf = make(chan []byte, 32)
	fmt.Println("TCP connect from: " + fmt.Sprintf("%s", this.connect.RemoteAddr()))

	go this.readListen()
	go this.writeListen()
}
func (this *TCPConnection) Read(buf []byte) (int, error) {
	var err error
	if this != nil && this.alive {
		if msg, ok := <-this.readBuf; ok {
			length := len(msg)
			if length <= MAX_NETWORK_PACKAGE_LENGTH {
				copy(buf, msg)
				return length, err
			} else {
				err = errors.New("Data package length is out of range.")
			}
		} else {
			err = errors.New("TCP read chan closed")
		}
	} else {
		err = errors.New("TCP connect closed")
	}
	return 0, err
}
func (this *TCPConnection) Write(msg []byte) (int, error) {
	var err error
	if this != nil && this.alive {
		if len(this.writeBuf) < cap(this.writeBuf) {
			length := uint32(len(msg))
			stream := bytes.NewBuffer([]byte{})
			binary.Write(stream, binary.BigEndian, length)
			binary.Write(stream, binary.BigEndian, msg)
			this.writeBuf <- stream.Bytes()
		} else {
			err = errors.New("TCPsocket connect is full.")
			return 0, err
		}
	} else {
		err = errors.New("TCPsocket connect closed.")
		return 0, err
	}
	return len(msg), err
}
func (this *TCPConnection) Close() error {
	var err error
	this.Lock()
	defer this.Unlock()
	this.alive = false
	close(this.readBuf)
	close(this.writeBuf)
	fmt.Println("TCP close: " + fmt.Sprintf("%s", this.connect.RemoteAddr()))
	err = this.connect.Close()
	return err
}
func (this *TCPConnection) readListen() {
	for this != nil && this.alive {
		buffer := make([]byte, 2048)
		length, err := this.connect.Read(buffer)
		if err == nil {
			this.Lock()
			defer this.Unlock()
			this.buffer = append(this.buffer, buffer[:length]...)
			if this.headRead {
				if len(this.buffer) >= int(this.packageLen) {
					this.headRead = false
					data := this.buffer[:this.packageLen]
					this.buffer = this.buffer[this.packageLen:len(this.buffer)]
					this.writeBuf <- data
				} else {
					break
				}
			} else {
				if len(this.buffer) >= 4 {
					lenBuffer := bytes.NewBuffer(this.buffer[0:4])
					err = binary.Read(lenBuffer, binary.BigEndian, &this.packageLen)
					utils.ErrorHandle("Read len", err)
					this.headRead = true
					this.buffer = this.buffer[4:len(this.buffer)]
				}
			}
		} else {
			this.Close()
		}
	}
}
func (this *TCPConnection) writeListen() {
	for this != nil && this.alive {
		if msg, ok := <-this.writeBuf; ok {
			go this.connect.Write(msg)
		} else {
			this.Close()
		}
	}
}

//==================== udp socket ====================

//==================== network ====================

type Network struct {
	sync.Mutex
	events.EventDispatcher

	alive       bool
	index       uint16
	connectType string
	createTime  int64
	deadline    int64
	markTime    int64

	router  *Route
	connect io.ReadWriteCloser
}

func (this *Network) Network() {
	this.router = new(Route)
	this.router.Route()
	this.router.addHandle(uint16(0), this.pong)
	this.EventDispatcher.EventDispatcher()
	this.deadline = 60
}
func (this *Network) Create(conn io.ReadWriteCloser) {
	this.alive = true
	this.connect = conn
	this.index = 0
	this.createTime = time.Now().Unix()

	go this.listen()

	evt := new(events.Event)
	evt.Type = events.CONNECT
	this.DispatchEvent(evt)
}
func (this *Network) GetIndex() uint16 {
	if this != nil && this.alive && this.connect != nil {
		return this.index
	}
	return 0
}
func (this *Network) GetCreateTime() int64 {
	if this != nil && this.alive && this.connect != nil {
		return this.createTime
	}
	return 0
}
func (this *Network) GetConnectType() string {
	if this != nil && this.alive && this.connect != nil {
		return this.connectType
	}
	return ""
}
func (this *Network) SetHeartBeat(interval time.Duration) error {
	var err error
	if this != nil && this.alive && this.connect != nil {
		go func() {
			for {
				if this != nil && this.alive && this.connect != nil {
					p := new(DataStruct)
					p.Title = uint16(0)
					p.Label = uint16(0)
					this.Send(p)
					time.Sleep(interval)
				} else {
					break
				}
			}
		}()
	} else {
		err = errors.New("Network is closed!")
	}
	return err
}
func (this *Network) Send(data *DataStruct) error {
	var err error
	if this != nil && this.alive && this.connect != nil {
		go this.flush(data)
	} else {
		err = errors.New("Network is closed!")
	}
	return err
}
func (this *Network) SendWithCallback(data *DataStruct, handle func(*DataStruct)) error {
	var err error
	if this != nil && this.alive && this.connect != nil {
		this.Lock()
		defer this.Unlock()
		this.index++
		if this.index == 0 {
			this.index = 1
		}
		data.SendIndex = this.index
		this.router.addCallback(this.index, handle)
		this.SendSafely(data)
	} else {
		err = errors.New("Network is closed!")
	}
	return err
}
func (this *Network) SendSafely(data *DataStruct) error {
	var err error
	if this != nil && this.alive && this.connect != nil {
		go func() {
			var count int = 0
			for {
				length, err := this.flush(data)
				if length > 0 && err == nil {
					break
				}
				count++
				if count >= 100 {
					break
				}
				time.Sleep(DELAY)
			}
		}()
	} else {
		err = errors.New("Network is closed!")
	}
	return err
}
func (this *Network) AddHandle(title uint16, handle func(*DataStruct)) error {
	var err error
	if this != nil && this.alive && this.connect != nil {
		err = this.router.addHandle(title, handle)
	} else {
		err = errors.New("Network is closed!")
	}
	return err
}
func (this *Network) SetDefaultHandle(defun func(*DataStruct)) error {
	var err error
	if this != nil && this.alive && this.connect != nil {
		this.router.setDefaultHandle(defun)
	} else {
		err = errors.New("Network is closed!")
	}
	return err
}
func (this *Network) RemoveHandle(title uint16) error {
	var err error
	if this != nil && this.alive && this.connect != nil {
		err = this.router.removeHandle(title)
	} else {
		err = errors.New("Network is closed!")
	}
	return err
}
func (this *Network) RemoveAllHandle() error {
	var err error
	if this != nil && this.alive && this.connect != nil {
		err = this.router.removeAllHandle()
	} else {
		err = errors.New("Network is closed!")
	}
	return err
}
func (this *Network) Close() error {
	var err error
	if this != nil && this.alive && this.connect != nil {
		this.Lock()
		defer this.Unlock()
		this.alive = false
		err = this.router.Dispose()
		this.router = nil
		err = this.connect.Close()
		this.connect = nil

		evt := new(events.Event)
		evt.Type = events.CLOSE
		this.DispatchEvent(evt)
	} else {
		err = errors.New("Network is closed!")
	}
	return err
}
func (this *Network) listen() {
	for this != nil && this.alive {
		buffer := make([]byte, 2048)
		length, err := this.connect.Read(buffer)
		if err == nil {
			recvData := new(DataStruct)
			titleBuffer := bytes.NewBuffer(buffer[0:2])
			err = binary.Read(titleBuffer, binary.BigEndian, &recvData.Title)
			utils.ErrorHandle("Read title", err)
			labelBuffer := bytes.NewBuffer(buffer[2:4])
			err = binary.Read(labelBuffer, binary.BigEndian, &recvData.Label)
			utils.ErrorHandle("Read label", err)
			sendIndexBuffer := bytes.NewBuffer(buffer[4:6])
			err = binary.Read(sendIndexBuffer, binary.BigEndian, &recvData.SendIndex)
			utils.ErrorHandle("Read send index", err)
			receIndexBuffer := bytes.NewBuffer(buffer[6:8])
			err = binary.Read(receIndexBuffer, binary.BigEndian, &recvData.ReceIndex)
			utils.ErrorHandle("Read rece index", err)
			if length > 8 {
				recvData.Data = buffer[8:length]
			}
			this.router.onData(recvData)
		} else {
			evt := new(events.Event)
			evt.Type = events.ERROR
			evt.Data = err

			this.DispatchEvent(evt)
			this.Close()
			break
		}
	}
}
func (this *Network) flush(data *DataStruct) (n int, err error) {
	stream := bytes.NewBuffer([]byte{})
	binary.Write(stream, binary.BigEndian, data.Title)
	binary.Write(stream, binary.BigEndian, data.Label)
	binary.Write(stream, binary.BigEndian, data.SendIndex)
	binary.Write(stream, binary.BigEndian, data.ReceIndex)
	binary.Write(stream, binary.BigEndian, data.Data)
	if this.connect != nil {
		return this.connect.Write(stream.Bytes())
	} else {
		err := errors.New("Connect lost.")
		return 0, err
	}
}
func (this *Network) pong(message *DataStruct) {
	if this.markTime == 0 {
		go func() {
			for {
				if this != nil && this.alive && this.connect != nil {
					if time.Now().Unix()-this.markTime > this.deadline {
						this.Close()
					}
					time.Sleep(time.Second * 10)
				} else {
					break
				}
			}
		}()
	}
	this.markTime = time.Now().Unix()
}
