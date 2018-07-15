package main

import (
	"fmt"
	"time"
	"github.com/btittelbach/go-bbhw"
)

type (
	GpioInputInfo struct {
		port uint
		gpio *bbhw.MMappedGPIO
		state  bool
	}
)

var (
	input_ports = []uint {
		// P8
		38, 39,
		34, 35,
		66, 67,
		69, 68,
		45, 44,
		23, 26,
		27, 65,
		22, 63,
		62, 37,
		36, 33,
		32, 61,
		86, 88,
		87, 89,
		10, 11,
		9, 81,
		8, 80,
		78, 79,
		76, 77,
		74, 75,
		72, 73,
		70, 71,
		// P9
		30, 60,
		31, 50,
		48, 51,
		5, 4,
		3, 2,
		49, 15,
		117, 14,
		115, 113,
		111, 112,
		31,
		41, 42,
	}
)


func main() {
	var gpio_inputs []*GpioInputInfo;

	for _, port := range input_ports {
		gpio := bbhw.NewMMappedGPIO(port, bbhw.IN)
		state, err := gpio.GetState()
		if err != nil {
			fmt.Printf("%v\n", err);
		}
		gpio_inputs = append(gpio_inputs, &GpioInputInfo {
			port: port,
			gpio: gpio,
			state: state,
		})
	}

	defer func() {
		fmt.Printf("Closing gpio ports\n");
		for _, gpio := range gpio_inputs {
			gpio.gpio.Close()
		}
		fmt.Printf("Terminated..\n");
	}()


	for {
		for _, gpio := range gpio_inputs {
			state, err := gpio.gpio.GetState()
			if err != nil {
				fmt.Printf("%v\n", err)
			}

			if state != gpio.state {
				fmt.Printf("Port: %v State: %v\n", gpio.port, state)
				gpio.state = state
			}
		}
		time.Sleep(100 * time.Millisecond)
	}

}


