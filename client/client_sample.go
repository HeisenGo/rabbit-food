package main

import (
	"fmt"
	"net"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Running...")

	// Read a message from the console
	//fmt.Print("Enter message: ")
	//reader := bufio.NewReader(os.Stdin)
	//message, _ := reader.ReadString('\n')

	// Send the message to the server
	//_, err = conn.Write([]byte(`{"location":"auth/register","header":{"method":"POST"},"data":{"phone":"09187567111","email":"kia55raa1sh@gk.co","password":"AliJaf@123"}}`))
	_, err = conn.Write([]byte(`{"location":"wallets/cards","header":{"method":"POST","Authorization": "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTkyNjkxODksIlVzZXJJRCI6MiwiSXNBZG1pbiI6ZmFsc2V9.yrZVongFbbqzzmlO7MkelKoo-5LR85MYmnMMyoII7bI_YOlf2kr7-kPve9_bpnJg-RCNiAFt-9glDsTYPvfLIA"},"data":{"card_number":"123456"}}`))
	fmt.Println("Data has been sent!")
	if err != nil {
		fmt.Println("Error writing to server:", err)
		return
	}

	// Read the response from the server
	// Read data from the connection
	buffer := make([]byte, 4096)
	bytesRead, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	fmt.Println("Received:", string(buffer[:bytesRead]))
	if err != nil {
		fmt.Println("Error reading from server:", err)
		return
	}
}
