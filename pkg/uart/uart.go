package uart

import (
	"errors"
	"fmt"

	"go.bug.st/serial"
	"go.bug.st/serial/enumerator"

	"runtime"
)

type Uart struct {
	Port           serial.Port
	SendChannel    chan []byte
	ReceiveChannel chan []byte
	ErrorHandler   func(error) // error handler
	Connection     bool        // 是否连接
}

func GetPortsList() []string {

	if runtime.GOOS == "windows" {
		ports, err := serial.GetPortsList()
		if err != nil {
			fmt.Println("Error listing ports: ", err)
			return nil
		}
		return ports
	}

	if runtime.GOOS == "linux" {
		// get usb ports
		usbPorts, err := enumerator.GetDetailedPortsList()
		if err != nil {
			fmt.Println("Error listing ports: ", err)
			return nil
		}
		portNames := make([]string, 0)

		for _, port := range usbPorts {
			// if port.IsUSB {
			// fmt.Printf("Port: %v\n", port)
			portNames = append(portNames, port.Name)
			// }
		}
		return portNames
	}

	return nil

}

func OpenPort(portName string, baudRate int) (*Uart, error) {

	mode := &serial.Mode{
		BaudRate: baudRate,
	}

	port, err := serial.Open(portName, mode)
	if err != nil {
		fmt.Println("Error opening port: ", err)
		return nil, err
	}

	uart := &Uart{
		Port:           port,
		SendChannel:    make(chan []byte),
		ReceiveChannel: make(chan []byte),
		ErrorHandler:   nil,
	}

	go uart.SendRoutine()
	go uart.ReceiveRoutine()
	uart.Connection = true

	return uart, nil
}

func (u *Uart) SetErrorHandler(handler func(error)) {
	u.ErrorHandler = handler
}

func (uart *Uart) SendRoutine() {

	for {
		data := <-uart.SendChannel
		_, err := uart.Port.Write(data)
		if err != nil {
			fmt.Println("Error writing to port: ", err)
			if uart.ErrorHandler != nil {
				uart.ErrorHandler(errors.New("连接已断开"))
				uart.Close()
				return
			}

		}
	}
}

func (uart *Uart) ReceiveRoutine() {

	for {
		data := make([]byte, 128)
		n, err := uart.Port.Read(data)
		if err != nil {
			fmt.Println("Error reading from port: ", err)
			if uart.ErrorHandler != nil {
				uart.ErrorHandler(errors.New("连接已断开"))
				uart.Close()
				return
			}
		}
		uart.ReceiveChannel <- data[:n]
	}
}

func (uart *Uart) Send(data []byte) {
	uart.SendChannel <- data
}

func (uart *Uart) Receive() []byte {
	return <-uart.ReceiveChannel
}

func (uart *Uart) Close() {
	uart.Connection = false
	uart.Port.Close()
}
