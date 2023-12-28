package game

import (
	"github.com/0990/goserver/network"
	"github.com/golang/protobuf/proto"
)

type EventType int

const (
	PlayerDisconnect EventType = iota
	PlayerReconnect
	ClientMsg_ReqClientReady
	ClientMsg_ReqInput
	ClientMsg_ReqGMCommand
	ClientMsg_ReqGameOver        = 5
	ClientMsg_ReportFrameKeyData = 6
	ClientMsg_ReqBroadcast       = 7
	ClientMsg_ReportSeatData     = 8
	ClientMsg_Ping               = 9

	PlayerLogin    = 20
	PlayerLogout   = 21
	AgonesAllocate = 22
)

type PlayerScore struct {
	userId uint64
	score  int32
}

type EventAgonesAllocateData struct {
	users []string
}

type EventPlayerLoginData struct {
	userId   uint64
	nickname string
	session  network.Session
}

type Sender interface {
	SendMsg(msg proto.Message) error
}

type Event struct {
	Type   EventType
	UserId uint64
	Data   interface{}
	Sender network.Session
}

type EventProcessor struct {
	events chan Event
}

func NewEventProcessor() *EventProcessor {
	return &EventProcessor{
		events: make(chan Event, 100),
	}
}

func (ep *EventProcessor) AddEvent(event Event) {
	ep.events <- event
}

func (ep *EventProcessor) Event() chan Event {
	return ep.events
}
