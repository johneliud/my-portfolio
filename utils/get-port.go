package utils

import (
	"log"
	"os"
	"strconv"
)

func GetPort() string {
	switch len(os.Args) {
	case 1:
		return ":8080"
	case 2:
		port := os.Args[1]
		portNum, err := strconv.Atoi(port[1:])
		if err != nil {
			log.Printf("Failed converting %v to an int: %v\n", port[1:], err)
			os.Exit(1)
		}

		if portNum < 1024 || portNum > 65535 {
			log.Println("Invalid port number. Must be in the range of 1024-65535")
			os.Exit(1)
		}
		return ":" + strconv.Itoa(portNum)
	default:
		return ":8080"
	}
}
