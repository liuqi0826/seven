package network

import (
	"bufio"
	"net"
	"time"
)

type CrossdomainListener struct {
	Alive       bool
	tcpAddress  *net.TCPAddr
	tcpListener *net.TCPListener
}

func (this *CrossdomainListener) CrossdomainListener() {

}
func (this *CrossdomainListener) Listen() error {
	var err error
	this.Alive = true
	this.tcpAddress, err = net.ResolveTCPAddr("tcp", ":843")
	this.tcpListener, err = net.ListenTCP("tcp", this.tcpAddress)
	go this.listen()
	return err
}
func (this *CrossdomainListener) listen() {
	for {
		if this.Alive {
			connect, err := this.tcpListener.Accept()
			if err == nil {
				cd := new(Crossdomain)
				cd.Crossdomain(connect)
			}
		} else {
			break
		}
	}
}

type Crossdomain struct {
	Alive   bool
	connect net.Conn
	ticker  *time.Ticker
}

func (this *Crossdomain) Crossdomain(connect net.Conn) {
	this.Alive = true
	this.connect = connect
	this.ticker = time.NewTicker(time.Second)
	policy := `<?xml version="1.0"?>
			<!DOCTYPE cross-domain-policy SYSTEM "/xml/dtds/cross-domain-policy.dtd">
			<cross-domain-policy>
				<site-control permitted-cross-domain-policies="master-only"/>
				<allow-access-from domain="*" to-ports="*" />
			</cross-domain-policy>`
	writer := bufio.NewWriter(this.connect)
	writer.WriteString(policy)
	writer.Flush()
	go this.wait()
}
func (this *Crossdomain) wait() {
	<-this.ticker.C
	this.dispose()
}
func (this *Crossdomain) dispose() {
	this.Alive = false
	this.ticker.Stop()
	this.connect.Close()
	this.connect = nil
}
