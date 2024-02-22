package api

import (
	"context"
	"fmt"
	"sync"
	"time"

	// This is 1/2 of the Dependency Issue (Import Cycle)
	//msg "GoImportCycle/services/agent/message"
)

// Define interface for injecting MSGServer
type MSGServer interface {
	InboundFromAPIServer()
}

type APIServer struct {
	Name string
	Run  bool
	Msg MSGServer
}

/**
func NewApiServer(name string) *APIServer {
	return &APIServer{
		Name: name,
		Run:  true,
	}
}
*/

func (p *APIServer) Start(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	//m2 := msg.NewMsgServer("Msg Server #2") // Create API Server using Dependency
	for p.Run != false {
		fmt.Println(p.Name)         // Print Server Name
		time.Sleep(time.Second * 1) // Sleep for 1 Second
		p.Msg.InboundFromAPIServer()  // Try and use Dependency from API Package
	}
	return
}

func (p *APIServer) InboundFromMessageServer() {
	fmt.Println("Message Server Reaches API Server..! Hello!")
}
