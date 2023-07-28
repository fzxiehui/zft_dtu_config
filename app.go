package main

import (
	"changeme/pkg/uart"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	go func() {
		for {
			time.Sleep(1 * time.Second)
			// fmt.Println("Hello")
			// get now time %H:%M:%S
			now := time.Now()
			// emit an event
			runtime.EventsEmit(a.ctx, "test", now.Format("15:04:05"))
		}
	}()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	fmt.Print(uart.GetPortsList())
	portList := uart.GetPortsList()
	if len(portList) == 0 {
		return "No ports found"
	}
	result := strings.Join(portList, ", ")
	return result
}

// Shutdown is called when the app is shutting down
func (a *App) shutdown() {
	fmt.Println("Goodbye")
}
