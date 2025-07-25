package pomelo

import (
	"go.uber.org/zap/zapcore"

	cfacade "github.com/cherry-game/cherry/facade"
	clog "github.com/cherry-game/cherry/logger"
	cactor "github.com/cherry-game/cherry/net/actor"
	cproto "github.com/cherry-game/cherry/net/proto"
)

const (
	ResponseFuncName = "response"
	PushFuncName     = "push"
	KickFuncName     = "kick"
	BroadcastName    = "broadcast"
)

type ActorBase struct {
	cactor.Base
}

func (p *ActorBase) Response(session *cproto.Session, v interface{}) {
	Response(p, session.AgentPath, session.Sid, session.Mid, v)
}

func (p *ActorBase) ResponseCode(session *cproto.Session, statusCode int32) {
	ResponseCode(p, session.AgentPath, session.Sid, session.Mid, statusCode)
}

func (p *ActorBase) Push(session *cproto.Session, route string, v interface{}) {
	Push(p, session.AgentPath, session.Sid, session.Uid, route, v)
}

func (p *ActorBase) Kick(session *cproto.Session, reason interface{}, closed bool) {
	Kick(p, session.AgentPath, session.Sid, session.Uid, reason, closed)
}

func (p *ActorBase) Broadcast(agentPath string, uidList []int64, allUID bool, route string, v interface{}) {
	Broadcast(p, agentPath, uidList, allUID, route, v)
}

func Response(iActor cfacade.IActor, agentPath, sid string, mid uint32, v interface{}) {
	data, err := iActor.App().Serializer().Marshal(v)
	if err != nil {
		clog.Warnf("[Response] Marshal error. v = %+v", v)
		return
	}

	rsp := &cproto.PomeloResponse{
		Sid:  sid,
		Mid:  mid,
		Data: data,
	}

	iActor.Call(agentPath, ResponseFuncName, rsp)

	if clog.PrintLevel(zapcore.DebugLevel) {
		clog.Debugf("[Response] agentPath = %s, sid = %s, mid = %d, message = %+v",
			agentPath, sid, mid, v)
	}
}

func ResponseCode(iActor cfacade.IActor, agentPath, sid string, mid uint32, statusCode int32) {
	rsp := &cproto.PomeloResponse{
		Sid:  sid,
		Mid:  mid,
		Code: statusCode,
	}

	iActor.Call(agentPath, ResponseFuncName, rsp)

	if clog.PrintLevel(zapcore.DebugLevel) {
		clog.Debugf("[ResponseCode] agentPath = %s, sid = %s, mid = %d, statusCode = %d",
			agentPath, sid, mid, statusCode)
	}
}

func Push(iActor cfacade.IActor, agentPath, sid string, uid int64, route string, v interface{}) {
	if route == "" {
		clog.Warn("[Push] route value error.")
		return
	}

	data, err := iActor.App().Serializer().Marshal(v)
	if err != nil {
		clog.Warnf("[Push] Marshal error. route =%s, v = %+v", route, v)
		return
	}

	rsp := &cproto.PomeloPush{
		Sid:   sid,
		Route: route,
		Data:  data,
	}

	iActor.Call(agentPath, PushFuncName, rsp)

	if clog.PrintLevel(zapcore.DebugLevel) {
		clog.Debugf("[Push] agentPath = %s, sid = %s, uid = %d, route = %s, message = %+v",
			agentPath, sid, uid, route, v)
	}
}

func Kick(iActor cfacade.IActor, agentPath, sid string, uid int64, reason interface{}, closed bool) {
	data, err := iActor.App().Serializer().Marshal(reason)
	if err != nil {
		clog.Warnf("[Kick] Marshal error. reason = %+v", reason)
		return
	}

	rsp := &cproto.PomeloKick{
		Sid:    sid,
		Reason: data,
		Close:  closed,
	}

	iActor.Call(agentPath, KickFuncName, rsp)

	clog.Infof("[Kick] agentPath = %s, sid = %s, uid = %d, reason = %+v, closed = %t",
		agentPath, sid, uid, reason, closed)
}

func Broadcast(iActor cfacade.IActor, agentPath string, uidList []int64, allUID bool, route string, v interface{}) {
	if !allUID && len(uidList) < 1 {
		clog.Warn("[Broadcast] uidList value error.")
		return
	}

	if route == "" {
		clog.Warn("[Broadcast] route value error.")
		return
	}

	data, err := iActor.App().Serializer().Marshal(v)
	if err != nil {
		clog.Warnf("[Broadcast] Marshal error. route = %s, v = %+v", route, v)
		return
	}

	rsp := &cproto.PomeloBroadcastPush{
		UidList: uidList,
		AllUID:  allUID,
		Route:   route,
		Data:    data,
	}

	iActor.Call(agentPath, BroadcastName, rsp)

	if clog.PrintLevel(zapcore.DebugLevel) {
		clog.Debugf("[Broadcast] agentPath = %s, uidList = %+v, allUID = %t, route = %s, message = %+v",
			agentPath, uidList, allUID, route, v)
	}
}
