package main

import (
	"context"
	"fmt"
)

// App struct
type PlanBus struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewPlan() *PlanBus {
	return &PlanBus{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *PlanBus) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *PlanBus) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
