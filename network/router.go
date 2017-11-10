package network

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type CallBack struct {
	CreateTime int64
	Handle     func(*DataStruct)
}

type Route struct {
	sync.Mutex

	Default  func(*DataStruct)
	Table    map[uint16]*CallBack
	Callback map[uint16]*CallBack

	ticker *time.Ticker
}

func (this *Route) Route() {
	this.Table = make(map[uint16]*CallBack)
	this.Callback = make(map[uint16]*CallBack)

	this.ticker = time.NewTicker(time.Minute * 10)
	go this.check()
}
func (this *Route) Dispose() error {
	var err error
	if this != nil {
		this.Default = nil
		this.Table = nil
		this.Callback = nil
	} else {
		err = errors.New("Route is disposed.")
	}
	return err
}
func (this *Route) setDefaultHandle(defun func(*DataStruct)) {
	if this != nil {
		this.Default = defun
	}
}
func (this *Route) addHandle(title uint16, handle func(*DataStruct)) error {
	var err error
	if this != nil {
		this.Lock()
		defer this.Unlock()
		if _, ok := this.Table[title]; ok {
			err = errors.New("Route handle is not null")
		} else {
			cb := new(CallBack)
			cb.CreateTime = time.Now().Unix()
			cb.Handle = handle
			this.Table[title] = cb
		}
	} else {
		err = errors.New("AddHandle function route is null")
	}
	return err
}
func (this *Route) removeHandle(title uint16) error {
	var err error
	if this != nil {
		this.Lock()
		defer this.Unlock()
		if _, ok := this.Table[title]; ok {
			delete(this.Table, title)
		} else {
			err = errors.New("Remove handle is null")
		}
	} else {
		err = errors.New("RemoveHandle function route is null")
	}
	return err
}
func (this *Route) removeAllHandle() error {
	var err error
	if this != nil {
		this.Lock()
		defer this.Unlock()
		this.Table = make(map[uint16]*CallBack)
	} else {
		err = errors.New("RemoveHandle function route is null")
	}
	return err
}
func (this *Route) addCallback(index uint16, handle func(*DataStruct)) error {
	var err error
	if this != nil {
		this.Lock()
		defer this.Unlock()
		if _, ok := this.Callback[index]; ok {
			err = errors.New("Callback handle is not null")
		} else {
			cb := new(CallBack)
			cb.CreateTime = time.Now().Unix()
			cb.Handle = handle
			this.Callback[index] = cb
		}
	} else {
		err = errors.New("AddCallback function route is null")
	}
	return err
}
func (this *Route) onData(ds *DataStruct) {
	if this != nil && ds != nil {
		if cb, ok := this.Callback[ds.ReceIndex]; ok {
			go cb.Handle(ds)
			delete(this.Callback, ds.ReceIndex)
			return
		}
		if cb, ok := this.Table[ds.Title]; ok {
			go cb.Handle(ds)
			return
		} else {
			if this.Default != nil {
				go this.Default(ds)
			} else {
				fmt.Println("Default route handle function is null")
				fmt.Println("Route code is: " + fmt.Sprintf("%d", ds.Title))
			}
		}
	} else {
		fmt.Println("router is nil")
	}
}
func (this *Route) check() {
	for this != nil {
		select {
		case <-this.ticker.C:
			now := time.Now().Unix()
			for k, v := range this.Callback {
				if now-v.CreateTime > 600 {
					go v.Handle(nil)
					delete(this.Callback, k)
				}
			}
		}
	}
}
