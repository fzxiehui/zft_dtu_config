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

// domReady 在前端Dom加载完毕后调用
func (a *App) domReady(ctx context.Context) {
	// Add your action here
	// 在这里添加你的操作
	fmt.Println("Dom is ready")
}

// beforeClose在单击窗口关闭按钮或调用runtime.Quit即将退出应用程序时被调用.
// 返回 true 将导致应用程序继续，false 将继续正常关闭。
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	fmt.Println("beforeClose")
	return false
}

// 在应用程序终止时被调用
func (a *App) shutdown(ctx context.Context) {
	// 在此处做一些资源释放的操作
	fmt.Println("shutdown")
}
