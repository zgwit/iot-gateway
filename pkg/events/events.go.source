package events

import "github.com/asaskevich/EventBus"

//EventInterface Events接口
type EventInterface interface {
	Emit(event string, data ...interface{})
	On(event string, fn interface{})
	Once(event string, fn interface{})
	Off(event string, fn interface{})
}

//EventEmitter 引入Emitter函数
type EventEmitter struct {
	events EventBus.Bus
}

//Emit 发送消息
func (e *EventEmitter) Emit(event string, data ...interface{}) {
	if e.events == nil {
		return
	}
	e.events.Publish(event, data...)
}

//On 监听
func (e *EventEmitter) On(event string, fn interface{}) {
	if e.events == nil {
		e.events = EventBus.New()
	}
	_ = e.events.Subscribe(event, fn)
}

//Once 监听一次
func (e *EventEmitter) Once(event string, fn interface{}) {
	if e.events == nil {
		e.events = EventBus.New()
	}
	_ = e.events.SubscribeOnce(event, fn)
}

//Off 取消监听
func (e *EventEmitter) Off(event string, fn interface{}) {
	if e.events == nil {
		return
	}
	_ = e.events.Unsubscribe(event, fn)
}
