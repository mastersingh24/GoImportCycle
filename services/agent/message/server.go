package message

import (
	api "GoImportCycle/services/agent/api"
	"context"
	"fmt"
	"sync"
	"time"
)

type MSGServer struct {
	Name string
	Run  bool
}

func NewMsgServer(name string) *MSGServer {
	return &MSGServer{
		Name: name,
		Run:  true,
	}
}

func (p *MSGServer) Start(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	a2 := api.NewApiServer("API Server #2")
	for p.Run != false {
		fmt.Println(p.Name)
		time.Sleep(time.Second * 1)
		a2.InboundFromMessageServer()
	}
	return
}

func (p *MSGServer) InboundFromAPIServer() {
	fmt.Println("API Server Reaches Message Server..! Hello!")
}
