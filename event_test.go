package event

import (
	"fmt"
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

func onEntrance(event Event){
	fmt.Println("saving to the database", event.GetData())
}

func TestCall(t *testing.T) {
	var e = Event{
		data:12,
	}
	s := NewSubscriber(test)
	s.call(e)
}

func TestDispatcher_Subscribe(t *testing.T) {
	d := NewDispatcher()
	d.Subscribe("start", test)
	d.Subscribe("start", onEntrance)
	var e = Event{}
	e.SetData(map[string]interface{}{
		"easy" : 12,
	})
	d.TriggerEvent("start",e)
}

func TestDispatcher_Subscribe1(t *testing.T) {
	d := NewDispatcher()
	d.Subscribe("start", test)
	d.Subscribe("end", onEntrance)
	var e = Event{}
	e.SetData(map[string]interface{}{
		"easy" : 12,
		"dispatcher":d,
	})
	d.TriggerEvent("start",e)
}