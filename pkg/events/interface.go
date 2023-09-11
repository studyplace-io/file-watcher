package events

import "github.com/fsnotify/fsnotify"

// IEventProcessor 处理变更事件接口
type IEventProcessor interface {
	// SendEvent 处理文件事件源逻辑
	SendEvent(ee fsnotify.Event) error
}

// EventProcessor 事件处理对象
type EventProcessor struct {
	IEventProcessor
}

func NewEventProcessor(IEventProcessor IEventProcessor) *EventProcessor {
	return &EventProcessor{IEventProcessor: IEventProcessor}
}


func (ep EventProcessor) SendEvent(ee fsnotify.Event) error {
	return ep.SendEvent(ee)
}

// K8sEventMode k8s事件模式
func K8sEventMode() IEventProcessor {
	return NewGenerator()
}