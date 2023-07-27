package uart

import (
	"fmt"

	"go.bug.st/serial"
)

type Uart struct {
	Port           serial.Port
	SendChannel    chan []byte
	ReceiveChannel chan []byte
}

func GetPortsList() []string {

	ports, err := serial.GetPortsList()
	if err != nil {
		fmt.Println("Error listing ports: ", err)
		return nil
	}

	return ports
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
	}

	go uart.SendRoutine()
	go uart.ReceiveRoutine()

	return uart, nil
}

func (uart *Uart) SendRoutine() {

	for {
		data := <-uart.SendChannel
		_, err := uart.Port.Write(data)
		if err != nil {
			fmt.Println("Error writing to port: ", err)
		}
	}
}

func (uart *Uart) ReceiveRoutine() {

	for {
		data := make([]byte, 128)
		n, err := uart.Port.Read(data)
		if err != nil {
			fmt.Println("Error reading from port: ", err)
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
	uart.Port.Close()
}
