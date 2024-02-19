package main

import (
	"context"

	agent "GoImportCycle/services/agent/agent"
)

func main() {

	ctx := context.Background()
	newAgent := agent.NewAgent(ctx)
	newAgent.Start()

	return
}
