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
	return &Agent{
		ctx: ctx,
		API: api.NewApiServer("API Server"),
		MSG: msg.NewMsgServer("Message Server"),
	}
}

func (a *Agent) Start() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go a.API.Start(a.ctx, &wg)
	go a.MSG.Start(a.ctx, &wg)
	wg.Wait()

}
