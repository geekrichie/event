package event

import (
	"errors"
	"reflect"
)

//type Event struct{
//	name string
//	callback func()
//}
type Subscriber struct{
	callback func([]interface{})
	arguments []interface{}
}

func NewSubscriber(callback func([]interface{}), arguments ...interface{}) *Subscriber{
	return &Subscriber{
		callback: callback,
		arguments: arguments,
	}
}

func (s *Subscriber) call() ([]reflect.Value, error){
	f := reflect.ValueOf(s.callback)
	if len(s.arguments) != f.Type().NumIn() {
		return nil, errors.New("the number of input params not match!")
	}
	in := make([]reflect.Value, len(s.arguments))
	for k, v := range s.arguments {
		in[k] = reflect.ValueOf(v)
	}
	return f.Call(in), nil
}



type Dispatcher struct {
	subscribers map[string][]*Subscriber
}

func New(name string) *Dispatcher{
	return  &Dispatcher{
		subscribers: make(map[string][]*Subscriber),
	}
}


func (d *Dispatcher) existEvent(name string) bool{
	_,ok := d.subscribers[name]
	return ok
}

func (d *Dispatcher) Subscribe(name string, callback func([]interface{}), auguments []interface{}) {
	if !d.existEvent(name){
		d.subscribers[name] = make([]*Subscriber,0)
	}
	d.subscribers[name]  =  append(d.subscribers[name], NewSubscriber(callback, auguments...))
}

func (d *Dispatcher) TriggerEvent(name string) {
	if !d.existEvent(name) {
		panic("event not exist")
	}
	for _,subscribers := range d.subscribers[name] {
		subscribers.call()
	}
}