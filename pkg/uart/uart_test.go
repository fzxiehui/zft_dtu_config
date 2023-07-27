package uart

import (
	"fmt"
	"testing"
)

func TestUart(t *testing.T) {
	fmt.Println("TestUart")

	ports := GetPortsList()
	if ports == nil {
		t.Error("No ports found")
	}

	fmt.Println("Ports found: ", ports)

	uart, err := OpenPort(ports[0], 115200)
	if err != nil {
		t.Error("Error opening port: ", err)
	}

	fmt.Println("Port opened: ", uart.Port)
}
