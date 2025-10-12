package daemons

import (
	"fmt"
	"time"
)

func StartNotifier() {
	go func() {
		for {
			fmt.Println("🔔 Checking for due tasks...")
			time.Sleep(5 * time.Second)
		}
	}()
}
