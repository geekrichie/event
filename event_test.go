package event

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

func test(event Event)  {
	fmt.Println(" enter the start event ")
	data := event.GetData()
	if content, ok := data.(map[string]interface{}); ok {
		d := content["dispatcher"].(*Dispatcher)
		d.TriggerEvent("end", event)
	}
}
func sample(event Event) {
	fmt.Println(event.GetData())
}

func onEntrance(event Event){
	fmt.Println("saving to the database", event.GetData())
}

func TestCall(t *testing.T) {
	var e = &SimpleEvent{}
	e.SetData(map[string]interface{}{
		"easy" : 12,
	})
	s := NewSubscriber(onEntrance)
	s.call(e)
}

func TestDispatcher_Subscribe(t *testing.T) {
	d := NewDispatcher()
	d.Subscribe("start", sample)
	d.Subscribe("start", onEntrance)
	var e = &SimpleEvent{}
	e.SetData(map[string]interface{}{
		"easy" : 12,
	})
	d.TriggerEvent("start",e)
}

func TestDispatcher_Subscribe1(t *testing.T) {
	d := NewDispatcher()
	d.Subscribe("start", test)
	d.Subscribe("end", onEntrance)
	var e = &SimpleEvent{}
	e.SetData(map[string]interface{}{
		"easy" : 12,
		"dispatcher":d,
	})
	d.TriggerEvent("start",e)
}

func TestFunc(t *testing.T) {
	v := reflect.ValueOf(test)
	if v.Kind() == reflect.Func {
		name := runtime.FuncForPC(v.Pointer()).Name()
		fmt.Println("Name of function : " + name)
	}
}

func TestDispatcher_UnSubscribe(t *testing.T) {
	d := NewDispatcher()
	d.Subscribe("start", test)
	d.Subscribe("end", onEntrance)
	d.Subscribe("end",sample)
	var e = &SimpleEvent{}
	e.SetData(map[string]interface{}{
		"easy" : 12,
		"dispatcher":d,
	})
	d.UnSubscribe("end", onEntrance)
	d.TriggerEvent("start",e)
}