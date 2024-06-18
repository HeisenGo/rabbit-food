package main

import (
	"bufio"
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
	_, err = conn.Write([]byte(`1{"phone": "09127827281", "email": "fuhu@hfhf.fh", "password": "Kiarash@123"}`))
	fmt.Println("Data has been sent!")
	if err != nil {
		fmt.Println("Error writing to server:", err)
		return
	}

	// Read the response from the server
	response, err := bufio.NewReader(conn).ReadString(' ')
	if err != nil {
		fmt.Println("Error reading from server:", err)
		return
	}
	fmt.Printf("Server response: %s", response)
}
