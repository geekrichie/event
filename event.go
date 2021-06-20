package event

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
)

type Subscriber struct{
	callback func(Event)
	arguments []interface{}
}

type Event interface{
	SetData(data interface{})
	GetData() interface{}
}

type SimpleEvent struct{
	data interface{}
}

func (e *SimpleEvent) SetData(data interface{}) {
	e.data = data
}

func (e *SimpleEvent) GetData() interface{}{
	return e.data
}

func NewSubscriber(callback func(Event)) *Subscriber{
	return &Subscriber{
		callback: callback,
	}
}

func (s *Subscriber) call(event Event){
	s.callback(event)
}

type Dispatcher struct {
	subscribers map[string]map[string]*Subscriber
}

func NewDispatcher() *Dispatcher{
	return  &Dispatcher{
		subscribers: make(map[string]map[string]*Subscriber),
	}
}


func (d *Dispatcher) ExistEvent(name string) bool{
	_,ok := d.subscribers[name]
	return ok
}

func (d *Dispatcher) Subscribe(name string, callback func(Event)) error{
	if !d.ExistEvent(name){
		d.subscribers[name] = make(map[string]*Subscriber)
	}
	funcname := GetFuncName(callback)
	if !d.AlreadySubscribed(name, funcname) {
		d.subscribers[name][funcname] =  NewSubscriber(callback)
	}else {
		return errors.New(fmt.Sprintf(" the func %s has already subscribed the event %s", funcname, name))
	}
	return nil
}
func (d *Dispatcher) AlreadySubscribed(name string, funcname string) bool{
	if _, ok := d.subscribers[name]; ok {
		if _, ok2 := d.subscribers[name][funcname]; ok2 {
			return true
		}
	}
	return false
}

func (d *Dispatcher) UnSubscribe(name string, callback func(Event)) {
	funcname := GetFuncName(callback)
	if d.AlreadySubscribed(name, funcname) {
		delete(d.subscribers[name],funcname)
	}
	if len(d.subscribers[name]) == 0 {
		d.RemoveEvent(name)
	}
}

func (d *Dispatcher) RemoveEvent(name string) {
	if d.ExistEvent(name) {
		delete(d.subscribers,name)
	}
}

func (d *Dispatcher) TriggerEvent(name string, event Event) {
	if !d.ExistEvent(name) {
		panic("event not exist")
	}
	for _,subscribers := range d.subscribers[name] {
		subscribers.call(event)
	}
}

func GetFuncName(target interface{}) string{
	v := reflect.ValueOf(target)
	if v.Kind() != reflect.Func{
		panic("the param is not of func type")
	}
	name :=  runtime.FuncForPC(v.Pointer()).Name()
	return name
}