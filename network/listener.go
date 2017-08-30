package network

import (
	"github.com/liuqi0826/seven/utils"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)

type Listener struct {
	alive       bool
	tcpAddress  *net.TCPAddr
	tcpListener *net.TCPListener
	upgrader    *websocket.Upgrader
	localChan   chan *ChanConnection
	networkChan chan *Network
	netType     string
	address     string
}

func (this *Listener) Listener() {
	this.networkChan = make(chan *Network, 0)
}
func (this *Listener) Listen(netType string, address string) error {
	var err error
	this.alive = true
	this.netType = netType
	this.address = address
	this.networkChan = make(chan *Network, 0)
	switch this.netType {
	case NETWORK_TYPE_LOCAL:
		this.localChan = make(chan *ChanConnection, 32)
	case NETWORK_TYPE_TCP:
		this.tcpAddress, err = net.ResolveTCPAddr("tcp", this.address)
		utils.ErrorHandle("Listener ResolveTCPAddr...", err)
		this.tcpListener, err = net.ListenTCP("tcp", this.tcpAddress)
		utils.ErrorHandle("Listener ListenTCP...", err)
	case NETWORK_TYPE_UDP:
	case NETWORK_TYPE_WEB_SOCKET:
		this.upgrader = new(websocket.Upgrader)
		this.upgrader.ReadBufferSize = 2048
		this.upgrader.WriteBufferSize = 2048
		this.upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	}
	go this.listen()
	return err
}
func (this *Listener) Dispose() error {
	var err error
	this.alive = false
	close(this.networkChan)
	switch this.netType {
	case NETWORK_TYPE_LOCAL:
		close(this.localChan)
	case NETWORK_TYPE_TCP:
	case NETWORK_TYPE_UDP:
	case NETWORK_TYPE_WEB_SOCKET:
	}
	return err
}
func (this *Listener) GetNetType() string {
	return this.netType
}
func (this *Listener) Accept() (*Network, error) {
	var err error
	network := <-this.networkChan
	return network, err
}
func (this *Listener) listen() {
	switch this.netType {
	case NETWORK_TYPE_LOCAL:
		for {
			if this.alive {
				for c := range this.localChan {
					connect := new(ChanConnection)
					connect.ChanConnection(c)
					nw := new(Network)
					nw.Network()
					nw.Create(connect)
					nw.connectType = NETWORK_TYPE_LOCAL
					go func() {
						this.networkChan <- nw
					}()
				}
			} else {
				break
			}
		}
	case NETWORK_TYPE_TCP:
		for {
			if this.alive {
				connect, err := this.tcpListener.Accept()
				if err == nil {
					nw := new(Network)
					nw.Network()
					nw.Create(connect)
					nw.connectType = NETWORK_TYPE_TCP
					go func() {
						this.networkChan <- nw
					}()
				}
			} else {
				break
			}
		}
	case NETWORK_TYPE_UDP:
	case NETWORK_TYPE_WEB_SOCKET:
		http.HandleFunc("/", this.onWebsocket)
		http.ListenAndServe(this.address, nil)
	}
}
func (this *Listener) onWebsocket(w http.ResponseWriter, r *http.Request) {
	ws, err := this.upgrader.Upgrade(w, r, nil)
	if err == nil {
		wsc := new(WSConnction)
		wsc.WSConnction(ws)
		nw := new(Network)
		nw.Network()
		nw.Create(wsc)
		nw.connectType = NETWORK_TYPE_WEB_SOCKET
		go func() {
			this.networkChan <- nw
		}()
	}
}
