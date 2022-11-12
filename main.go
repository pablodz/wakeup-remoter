package main

import (
	"fmt"
	"os"

	"github.com/pablodz/wakeup-remoter/internal/server"
)

func main() {
	fmt.Println("uid", os.Geteuid())
	server.ServeWakeUpRemoter()
}
