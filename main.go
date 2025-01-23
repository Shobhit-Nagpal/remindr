package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func main() {
	fmt.Println("Hello from Remindr")

	ticker := time.NewTicker(5 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				notifyCmd := exec.Command("notify-send", "get up, bhai. kya ho raha hai")
				if err := notifyCmd.Run(); err != nil {
					log.Fatal(err)
				}

			}
		}
	}()

  time.Sleep(30 * time.Second)
}
