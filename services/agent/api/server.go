package api

import (
	"context"
	"fmt"
	"sync"
	"time"

	msg "GoImportCycle/services/agent/message"
)

type APIServer struct {
	Name string
	Run  bool
}

func NewApiServer(name string) *APIServer {
	return &APIServer{
		Name: name,
		Run:  true,
	}
}

func (p *APIServer) Start(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	m2 := msg.NewMsgServer("Msg Server #2")
	for p.Run != false {
		fmt.Println(p.Name)
		time.Sleep(time.Second * 1)
		m2.InboundFromAPIServer()
	}
	return
}

func (p *APIServer) InboundFromMessageServer() {
	fmt.Println("Message Server Reaches API Server..! Hello!")
}
