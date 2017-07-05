package network

import (
	"errors"
	"fmt"
	"sync"
)

type Route struct {
	sync.Mutex
	Default  func(*DataStruct)
	Table    map[uint16]func(*DataStruct)
	Callback map[uint16]func(*DataStruct)
}

func (this *Route) Route() {
	this.Table = make(map[uint16]func(*DataStruct))
	this.Callback = make(map[uint16]func(*DataStruct))
}
func (this *Route) Dispose() {
	if this != nil {
		this.Default = nil
		this.Table = nil
		this.Callback = nil
	}
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
			this.Table[title] = handle
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
func (this *Route) addCallback(index uint16, handle func(*DataStruct)) error {
	var err error
	if this != nil {
		this.Lock()
		defer this.Unlock()
		if _, ok := this.Callback[index]; ok {
			err = errors.New("Callback handle is not null")
		} else {
			this.Callback[index] = handle
		}
	} else {
		err = errors.New("AddCallback function route is null")
	}
	return err
}
func (this *Route) onData(ds *DataStruct) {
	if this != nil {
		if fun, ok := this.Callback[ds.ReceIndex]; ok {
			go fun(ds)
			delete(this.Callback, ds.ReceIndex)
			return
		}
		if fun, ok := this.Table[ds.Title]; ok {
			go fun(ds)
			return
		} else {
			if this.Default != nil {
				go this.Default(ds)
			} else {
				fmt.Println("Default route handle function is null")
				fmt.Println("Route code is: " + fmt.Sprintf("%d", ds.Title))
			}
		}
	}
}
