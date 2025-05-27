package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/reagin/double-ratchet/server"
	"github.com/reagin/double-ratchet/utils"
)

func main() {
	go startServiceServer()
	defer stopServiceServer()

	fmt.Println("press [ h | help ] to show help information")

	for {
		choice, err := utils.ScanfString()
		if err != nil {
			choice = "quit"
			fmt.Printf("Unknow Error: %s\n", err.Error())
		}

		switch choice {
		case "q", "quit":
			return
		case "h", "help":
			fmt.Println("press [ q | quit ] to quit server")
		default:
			fmt.Println("press [ h | help ] to show help information")
		}
	}
}

func startServiceServer() {
	log.Println("Service server is running...")

	if err := server.ServiceServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("Start Service Server Error: %s\n", err)
	}
}

func stopServiceServer() {
	log.Println("Service server is stopping...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := server.ServiceServer.Shutdown(ctx); err != nil {
		log.Printf("Stop Service Server Error: %s\n", err)
	}
}
