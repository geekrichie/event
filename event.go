package event


type Subscriber struct{
	callback func(Event)
	arguments []interface{}
}

type Event struct{
	data interface{}
}

func (e *Event) SetData(data interface{}) {
	e.data = data
}

func (e *Event) GetData() interface{}{
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
	subscribers map[string][]*Subscriber
}

func NewDispatcher() *Dispatcher{
	return  &Dispatcher{
		subscribers: make(map[string][]*Subscriber),
	}
}


func (d *Dispatcher) ExistEvent(name string) bool{
	_,ok := d.subscribers[name]
	return ok
}

func (d *Dispatcher) Subscribe(name string, callback func(Event)) {
	if !d.ExistEvent(name){
		d.subscribers[name] = make([]*Subscriber,0)
	}
	d.subscribers[name]  =  append(d.subscribers[name], NewSubscriber(callback))
}

func (d *Dispatcher) TriggerEvent(name string, event Event) {
	if !d.ExistEvent(name) {
		panic("event not exist")
	}
	for _,subscribers := range d.subscribers[name] {
		subscribers.call(event)
	}
}