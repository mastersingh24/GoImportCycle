package message

import (
	// This is 1/2 of the Dependency Issue (Import Cycle)
	api "GoImportCycle/services/api"
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
	a2 := api.NewApiServer("API Server #2") // Create API Server using Dependency
	for p.Run != false {
		fmt.Println(p.Name)           // Print Server Name
		time.Sleep(time.Second * 1)   // Sleep for 1 Second
		a2.InboundFromMessageServer() // Try and use Dependency from API Package
	}
	return
}

func (p *MSGServer) InboundFromAPIServer() {
	fmt.Println("API Server Reaches Message Server..! Hello!")
}
