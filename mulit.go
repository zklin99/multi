package main

import (
	"fmt"
	"os/exec"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	numThreads := 2

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cmd := exec.Command("D:/programs/WeChat/WeChat.exe")
			err := cmd.Start()
			if err != nil {
				fmt.Println("Error opening WeChat:", err)
			}
		}()
	}

	wg.Wait()
}
