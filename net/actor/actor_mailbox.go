package cherryActor

import (
	"strings"

	cconst "github.com/cherry-game/cherry/const"
	creflect "github.com/cherry-game/cherry/extend/reflect"
	ctime "github.com/cherry-game/cherry/extend/time"
	cfacade "github.com/cherry-game/cherry/facade"
	clog "github.com/cherry-game/cherry/logger"
)

type mailbox struct {
	queue                                 // queue
	name    string                        // 邮箱名
	funcMap map[string]*creflect.FuncInfo // 已注册的函数
}

func newMailbox(name string) mailbox {
	return mailbox{
		queue:   newQueue(),
		name:    name,
		funcMap: make(map[string]*creflect.FuncInfo),
	}
}

func getLastSegment(funcName string) string {
	parts := strings.Split(funcName, cconst.DOT)
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return funcName
}

func (p *mailbox) Register(funcName string, fn interface{}) {
	if funcName == "" || len(funcName) < 1 {
		clog.Errorf("[%s] Func name is empty.", fn)
		return
	}

	funcName = getLastSegment(funcName)
	funcInfo, err := creflect.GetFuncInfo(fn)
	if err != nil {
		clog.Errorf("funcName = %s, err = %v", funcName, err)
		return
	}

	if _, found := p.funcMap[funcName]; found {
		clog.Errorf("funcName = %s, already exists.", funcName)
		return
	}

	p.funcMap[funcName] = &funcInfo
}

func (p *mailbox) GetFuncInfo(funcName string) (*creflect.FuncInfo, bool) {
	funcInfo, found := p.funcMap[funcName]
	return funcInfo, found
}

func (p *mailbox) Pop() *cfacade.Message {
	v := p.queue.Pop()
	if v == nil {
		return nil
	}

	msg, ok := v.(*cfacade.Message)
	if !ok {
		clog.Warnf("Convert to *Message fail. v = %+v", v)
		return nil
	}

	return msg
}

func (p *mailbox) Push(m *cfacade.Message) {
	if m != nil {
		m.PostTime = ctime.Now().ToMillisecond()
		p.queue.Push(m)
	}
}

func (p *mailbox) onStop() {
	for key := range p.funcMap {
		delete(p.funcMap, key)
	}

	p.queue.Destroy()
}
