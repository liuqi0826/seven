package network

import (
	//"fmt"

	"net"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/liuqi0826/seven/utils"
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
	pattern     string
	wsConfig    *WebSocketConfig
}

func (this *Listener) Listener() {
	this.networkChan = make(chan *Network, 0)
}
func (this *Listener) Listen(netType string, address string, extra interface{}) error {
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
		if cfg, ok := extra.(*WebSocketConfig); ok {
			this.wsConfig = cfg
			if this.wsConfig.URL == "" {
				this.pattern = "/"
			} else if strings.Index(this.wsConfig.URL, "/") < 0 {
				this.pattern = "/" + this.wsConfig.URL
			} else {
				this.pattern = this.wsConfig.URL
			}
		} else {
			this.pattern = "/"
		}
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
				conn, err := this.tcpListener.Accept()
				if err == nil {
					connect := new(TCPConnection)
					connect.TCPConnection(conn)
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
		http.HandleFunc(this.pattern, this.onWebsocket)
		if this.wsConfig != nil && this.wsConfig.SSLCRT != "" && this.wsConfig.SSLKey != "" {
			http.ListenAndServeTLS(this.address, this.wsConfig.SSLCRT, this.wsConfig.SSLKey, nil)
		} else {
			http.ListenAndServe(this.address, nil)
		}
	}
}
func (this *Listener) onWebsocket(w http.ResponseWriter, r *http.Request) {
	ws, err := this.upgrader.Upgrade(w, r, nil)
	if err == nil {
		wsc := new(WSConnction)
		wsc.WSConnction(ws)
		err = r.ParseForm()
		if err == nil {
			wsc.Values = r.Form
		}
		nw := new(Network)
		nw.Network()
		nw.Create(wsc)
		nw.connectType = NETWORK_TYPE_WEB_SOCKET
		go func() {
			this.networkChan <- nw
		}()
	}
}
