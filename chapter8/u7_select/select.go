package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	abort := make(chan struct{})
	go func() {
		_, _ = os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("commencing countdown.")

	select {
	case <-time.After(10 * time.Second):

	case <-abort:
		fmt.Println("launch aborted!")
		return

	}
	fmt.Println("launch!!!")
}
