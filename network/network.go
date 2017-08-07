package network

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/liuqi0826/seven/events"
	"github.com/liuqi0826/seven/utils"

	"github.com/gorilla/websocket"
)

type Buffer struct {
	headRead  bool
	bufferLen int32
	buffer    []byte
	reader    io.Reader
}

type DataStruct struct {
	Title     uint16
	Label     uint16
	SendIndex uint16
	ReceIndex uint16
	Data      []byte
}

//==================== local ====================

type Connection struct {
	Alive       bool
	Channel     chan []byte
	BindingConn *Connection
}

func (this *Connection) Connection() {
	this.Alive = true
	this.Channel = make(chan []byte, 10)
}
func (this *Connection) Read(buf []byte) (int, error) {
	var err error
	if this.BindingConn == nil {
		for {
			time.Sleep(time.Millisecond * 100)
			if this.BindingConn != nil {
				break
			}
		}
	}
	if msg, ok := <-this.BindingConn.Channel; ok {
		length := len(msg)
		stream := bytes.NewBuffer([]byte{})
		binary.Write(stream, binary.BigEndian, length)
		binary.Write(stream, binary.BigEndian, msg)
		copy(buf, stream.Bytes())
		return length, err
	} else {
		err = errors.New("Bingding connect channel closed")
	}
	return 0, err
}
func (this *Connection) Write(p []byte) (int, error) {
	var err error
	if this.Alive {
		this.Channel <- p
	} else {
		err = errors.New("Connection channel closed")
	}
	return len(p), err
}
func (this *Connection) Close() error {
	if !this.Alive {
		return nil
	}
	this.Alive = false
	close(this.Channel)
	return nil
}

//==================== websocket ====================

type WSConnction struct {
	Alive      bool
	ws         *websocket.Conn
	writeBuff  chan []byte
	errorCount int
}

func (this *WSConnction) WSConnction(ws *websocket.Conn) {
	this.Alive = true
	this.ws = ws
	this.writeBuff = make(chan []byte, 32)
	this.errorCount = 0
	fmt.Println("Websocket connect from: " + fmt.Sprintf("%s", this.ws.RemoteAddr()))

	go this.writeHandle()
}
func (this *WSConnction) Read(buf []byte) (int, error) {
	var err error
	if this.Alive {
		mt, msg, err := this.ws.ReadMessage()
		if err == nil {
			switch mt {
			default:
				length := len(msg)
				if len(buf) < length {
					buf = make([]byte, length)
				}
				copy(buf, msg)
				return length, err
			}
		} else {
			this.errorCount++
			if this.errorCount >= 10 {
				err = this.Close()
			}
			err = errors.New("Websocket read error")
		}
	} else {
		err = errors.New("Websocket connection closed")
	}
	return 0, err
}
func (this *WSConnction) Write(p []byte) (int, error) {
	var err error
	if this.Alive {
		this.writeBuff <- p
	} else {
		err = errors.New("Websocket connection closed")
	}
	return 0, err
}
func (this *WSConnction) Close() error {
	var err error
	if !this.Alive {
		return nil
	}
	this.Alive = false
	err = this.ws.Close()
	close(this.writeBuff)
	return err
}
func (this *WSConnction) writeHandle() {
	for {
		select {
		case message, ok := <-this.writeBuff:
			if ok {
				err := this.ws.WriteMessage(websocket.BinaryMessage, message)
				if err != nil {
					this.errorCount++
					if this.errorCount >= 10 {
						err = this.Close()
					}
					err = errors.New("Websocket write error")
				}
			} else {
				this.ws.WriteMessage(websocket.CloseMessage, []byte{})
				this.Alive = false
				return
			}
		}
	}
}

//==================== tcp socket ====================

//==================== network ====================

type Network struct {
	Buffer
	events.EventDispatcher
	Alive       bool
	Index       uint16
	CreateTime  int64
	Deadline    int64
	Router      *Route
	connectType string
	connect     io.ReadWriteCloser
	markTime    int64
	errorCount  int
}

func (this *Network) Network() {
	this.EventDispatcher.EventDispatcher()
	this.Router = new(Route)
	this.Router.Route()
	this.Router.addHandle(uint16(0), this.pong)
	this.Deadline = 60
}
func (this *Network) Create(conn io.ReadWriteCloser) {
	this.Alive = true
	this.Index = 0
	this.connect = conn
	this.reader = bufio.NewReader(this.connect)
	this.CreateTime = time.Now().Unix()
	this.listen()
	evt := new(events.Event)
	evt.Type = events.CONNECT
	this.DispatchEvent(evt)
}
func (this *Network) SetHeartBeat(interval time.Duration) {
	go func() {
		for {
			if this != nil && this.Alive && this.connect != nil {
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
}
func (this *Network) GetConnectType() string {
	return this.connectType
}
func (this *Network) Send(data *DataStruct) {
	if this != nil {
		this.Index += 1
		data.SendIndex = this.Index
		go this.write(data)
	} else {
		fmt.Println("network is nil")
	}
}
func (this *Network) SendWithCallback(data *DataStruct, handle func(*DataStruct)) {
	this.Index += 1
	data.SendIndex = this.Index
	this.Router.addCallback(this.Index, handle)
	go this.write(data)
}
func (this *Network) AddHandle(title uint16, handle func(*DataStruct)) error {
	return this.Router.addHandle(title, handle)
}
func (this *Network) SetDefaultHandle(defun func(*DataStruct)) {
	this.Router.setDefaultHandle(defun)
}
func (this *Network) RemoveHandle(title uint16) error {
	return this.Router.removeHandle(title)
}
func (this *Network) Close() error {
	var err error
	if this != nil && this.Alive {
		this.Alive = false
		this.Router.Dispose()
		this.Router = nil
		err = this.connect.Close()
		evt := new(events.Event)
		evt.Type = events.CLOSE
		this.DispatchEvent(evt)
	} else {
		err = errors.New("Network is closed!")
	}
	return err
}
func (this *Network) listen() {
	go this.read()
}
func (this *Network) pong(message *DataStruct) {
	if this.markTime == 0 {
		go func() {
			for {
				if this.Alive {
					if time.Now().Unix()-this.markTime > this.Deadline {
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
func (this *Network) write(data *DataStruct) {
	length := int32(len(data.Data) + 8)
	stream := bytes.NewBuffer([]byte{})
	binary.Write(stream, binary.BigEndian, length)
	binary.Write(stream, binary.BigEndian, data.Title)
	binary.Write(stream, binary.BigEndian, data.Label)
	binary.Write(stream, binary.BigEndian, data.SendIndex)
	binary.Write(stream, binary.BigEndian, data.ReceIndex)
	binary.Write(stream, binary.BigEndian, data.Data)
	this.connect.Write(stream.Bytes())
}
func (this *Network) read() {
	for {
		if this.Alive {
			buf := make([]byte, 256)
			length, err := this.reader.Read(buf)
			if err == nil {
				this.buffer = append(this.buffer, buf[:length]...)
				for len(this.buffer) > 0 {
					if this.headRead {
						if len(this.buffer) >= int(this.bufferLen) {
							this.headRead = false
							recvData := new(DataStruct)
							titleBuffer := bytes.NewBuffer(this.buffer[0:2])
							err = binary.Read(titleBuffer, binary.BigEndian, &recvData.Title)
							utils.ErrorHandle("Read title", err)
							labelBuffer := bytes.NewBuffer(this.buffer[2:4])
							err = binary.Read(labelBuffer, binary.BigEndian, &recvData.Label)
							utils.ErrorHandle("Read label", err)
							sendIndexBuffer := bytes.NewBuffer(this.buffer[4:6])
							err = binary.Read(sendIndexBuffer, binary.BigEndian, &recvData.SendIndex)
							utils.ErrorHandle("Read send index", err)
							receIndexBuffer := bytes.NewBuffer(this.buffer[6:8])
							err = binary.Read(receIndexBuffer, binary.BigEndian, &recvData.ReceIndex)
							utils.ErrorHandle("Read rece index", err)
							
							if this.bufferLen > 8 {
								recvData.Data = this.buffer[8:this.bufferLen]
							}
							
							this.buffer = this.buffer[this.bufferLen:len(this.buffer)]
							this.Router.onData(recvData)
						} else {
							break
						}
					} else {
						if len(this.buffer) >= 4 {
							lenBuffer := bytes.NewBuffer(this.buffer[0:4])
							err = binary.Read(lenBuffer, binary.BigEndian, &this.bufferLen)
							utils.ErrorHandle("Read len", err)
							this.headRead = true
							this.buffer = this.buffer[4:len(this.buffer)]
						}
					}
				}
			} else {
				this.errorCount++
				evt := new(events.Event)
				evt.Type = events.ERROR
				evt.Data = err
				this.DispatchEvent(evt)
				if this.errorCount >= 10 {
					this.Close()
					break
				}
			}
		} else {
			break
		}
	}
}
