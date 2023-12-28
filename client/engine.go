package client

import (
	"github.com/0990/goserver/network"
	"github.com/0990/goserver/util"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"reflect"
)

type Engine struct {
	processor  *network.Processor
	conn       *websocket.Conn
	targetAddr string
}

func NewEngine(targetAddr string) *Engine {

	return &Engine{
		targetAddr: targetAddr,
		processor:  network.NewProcessor(),
	}
}

func (p *Engine) Close() {
	p.conn.Close()
}

func (p *Engine) Run() error {
	c, _, err := websocket.DefaultDialer.Dial(p.targetAddr, nil)
	if err != nil {
		return err
	}
	p.conn = c
	go p.ReadLoop()
	return nil
}

func (p *Engine) ReadLoop() {
	for {
		_, data, err := p.conn.ReadMessage()
		if err != nil {
			logrus.WithError(err).Error("read message")
			break
		}

		msg, err := p.processor.Unmarshal(data)
		if err != nil {
			logrus.Errorf("unmarshal message error: %v", err)
			continue
		}
		p.processor.Handle(msg, p)
	}
}

func (p *Engine) SendMsg(msg proto.Message) {
	data, _ := p.processor.Marshal(msg)
	err := p.conn.WriteMessage(websocket.BinaryMessage, data)
	if err != nil {
		logrus.WithError(err).Error("write message")
	}
}

func (p *Engine) SendRawMsg(msgID uint32, data []byte) {
	err := p.conn.WriteMessage(websocket.BinaryMessage, data)
	if err != nil {
		logrus.WithError(err).Error("write message")
	}
}

func (p *Engine) ID() int32 {
	return 0
}

func (p *Engine) RegisterSessionMsgHandler(cb interface{}) {
	err, funValue, msgType := util.CheckArgs1MsgFun(cb)
	if err != nil {
		logrus.WithError(err).Error("RegisterServerMsgHandler")
		return
	}
	msg := reflect.New(msgType).Elem().Interface().(proto.Message)
	p.processor.RegisterSessionMsgHandler(msg, func(s network.Session, message proto.Message) {
		funValue.Call([]reflect.Value{reflect.ValueOf(s), reflect.ValueOf(message)})
	})
}
