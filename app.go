package main

import (
	"changeme/pkg/uart"
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx  context.Context
	uart *uart.Uart
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

// 获取串口列表
func (a *App) GetPortsList() []string {

	ports := uart.GetPortsList()
	if len(ports) == 0 {
		return []string{"No ports found"}
	}
	return ports
}

func (a *App) OpenPort(port string, baud int) string {
	// fmt.Println("OpenPort")
	uart, err := uart.OpenPort(port, baud)
	if err != nil {
		return "打开串口失败"
	}
	a.uart = uart

	a.uart.SetErrorHandler(func(err error) {
		runtime.EventsEmit(a.ctx, "error", err.Error())
	})

	// read data
	go func() {
		for {
			data := a.uart.Receive()
			// emit an event
			fmt.Println("data", string(data))
			runtime.EventsEmit(a.ctx, "read", base64.StdEncoding.EncodeToString(data))

		}
	}()

	return "打开串口成功"
}

func (a *App) ClosePort() string {
	if a.uart == nil {
		return "关闭失败"
	}
	a.uart.Close()
	return "已成功关闭串口"
}

// func (a *App) Send(data []byte) string {
// 	if a.uart == nil {
// 		return "发送失败"
// 	}
// 	a.uart.Send(data)
// 	return "发送成功"
// }

// func (a *App) Send(data string) string {
// 	if a.uart == nil {
// 		return "发送失败"
// 	}
// 	a.uart.Send([]byte(data))
// 	return "发送成功"
// }

func (a *App) Send(dataBase64 string) string {
	if a.uart == nil {
		return "发送失败"
	}
	data, err := base64.StdEncoding.DecodeString(dataBase64)
	if err != nil {
		return "发送失败"
	}
	a.uart.Send(data)
	return "发送成功"
}
