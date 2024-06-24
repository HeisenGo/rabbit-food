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
	//_, err = conn.Write([]byte(`{"location":"auth/register","header":{"method":"POST"},"data":{"phone":"09391063672","email":"kih123@gk.co","password":"AliJaf@123"}}`))
	//_, err = conn.Write([]byte(`{"location":"wallets/cards","header":{"method":"POST","Authorization": "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTkzMjAxNTgsIlVzZXJJRCI6MywiSXNBZG1pbiI6ZmFsc2V9.7dT6MIKbKUOizMbn31yXXFmfc8MKJ3qg4wqmD9QMsuw6dSWtgaJCSFQTQp0UbuiHyRh3IuD0NqegTGFcLuWimQ"},"data":{"card_number":"112233444"}}`))
	//_, err = conn.Write([]byte(`{"location":"wallets/deposit","header":{"method":"POST","Authorization": "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTkzMjAxNTgsIlVzZXJJRCI6MywiSXNBZG1pbiI6ZmFsc2V9.7dT6MIKbKUOizMbn31yXXFmfc8MKJ3qg4wqmD9QMsuw6dSWtgaJCSFQTQp0UbuiHyRh3IuD0NqegTGFcLuWimQ"},"data":{"card_number":"112233444", "amount":10000}}`))
	//_, err = conn.Write([]byte(`{"location":"wallets/withdraw","header":{"method":"POST","Authorization": "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTkzMjAxNTgsIlVzZXJJRCI6MywiSXNBZG1pbiI6ZmFsc2V9.7dT6MIKbKUOizMbn31yXXFmfc8MKJ3qg4wqmD9QMsuw6dSWtgaJCSFQTQp0UbuiHyRh3IuD0NqegTGFcLuWimQ"},"data":{"card_number":"112233444", "amount":8000}}`))
	_, err = conn.Write([]byte(`{"location":"wallets/cards","header":{"method":"GET","Authorization": "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTkzMjAxNTgsIlVzZXJJRCI6MywiSXNBZG1pbiI6ZmFsc2V9.7dT6MIKbKUOizMbn31yXXFmfc8MKJ3qg4wqmD9QMsuw6dSWtgaJCSFQTQp0UbuiHyRh3IuD0NqegTGFcLuWimQ"}}`))
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
