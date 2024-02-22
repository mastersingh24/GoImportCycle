package agent

import (
	"context"
	"sync"

	api "GoImportCycle/services/agent/api"
	msg "GoImportCycle/services/agent/message"
)

type Agent struct {
	ctx context.Context
	API *api.APIServer
	MSG *msg.MSGServer
}

func NewAgent(ctx context.Context) *Agent {


	a:= &api.APIServer{
		Name: "API Server",
		Run: true,
	}

	m:= &msg.MSGServer{
		Name: "Message Server",
		Run: true,
	}

	//a.InboundFromAPIServer = m.InboundFromAPIServer()
	//m.InboundFromMessageServer = a.InboundFromMessageServer()
	a.InboundFromAPIServer = func(){m.InboundFromAPIServer()}
	m.InboundFromMessageServer = func(){a.InboundFromMessageServer()}

	return &Agent{
		ctx: ctx,
		API: a,
		MSG: m,
	}
}

func (a *Agent) Start() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go a.API.Start(a.ctx, &wg)
	go a.MSG.Start(a.ctx, &wg)
	wg.Wait()

}
