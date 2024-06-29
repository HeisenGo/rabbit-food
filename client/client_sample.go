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

	//_, err = conn.Write([]byte(`{"location":"auth/register","header":{"method":"POST"},"data":{"phone":"09337317958","email":"jasem111@yahoo.com","password":"Jaf@123"}}`))
	//_, err = conn.Write([]byte(`{"location":"wallets/cards","header":{"method":"POST","Authorization": "Bearer "},"data":{"card_number":"112233444"}}`))
	//_, err = conn.Write([]byte(`{"location":"wallets/deposit","header":{"method":"POST","Authorization": "Bearer "},"data":{"card_number":"112233444", "amount":10000}}`))
	//_, err = conn.Write([]byte(`{"location":"wallets/withdraw","header":{"method":"POST","Authorization": "Bearer "},"data":{"card_number":"112233444", "amount":8000}}`))
	//_, err = conn.Write([]byte(`{"location":"wallets/cards","header":{"method":"GET","Authorization": "Bearer MTkzMjAxNTgsIlVzZXJJRCI6MywiSXNBZG1pbiI6ZmFsc2V9."}}`))

	//_, err = conn.Write([]byte(`{"location":"users/addresses","header":{"method":"POST","Authorization": "Bearer "},"data":{"address_line":"borje milad","coordinates":{"lat":35.7448, "lng":51.3755},"city":"tehran"}}`))
	//_, err = conn.Write([]byte(`{"location":"restaurants/menus","header":{"method":"POST","Authorization": "Bearer "}, "data":{"name":"burgurs", "restaurant_id":3}}`))
	//_, err = conn.Write([]byte(`{"location":"restaurants/menus","header":{"method":"GET","Authorization": "Bearer "}, "data":{"restaurant_id":3}}`))
	//_, err = conn.Write([]byte(`{"location":"restaurants/menu-items","header":{"method":"POST","Authorization": "Bearer "}, "data":{"menu_id":1, "price":200000, "name":"kabab koobide", "preparation_minutes": 40, "cancellation_penalty_percentage": 9}}`))
	//_, err = conn.Write([]byte(`{"location":"restaurants/menu-items","header":{"method":"GET","Authorization": "Bearer "}, "data":{"menu_id":1}}`))
	//_, err = conn.Write([]byte(`{"location":"wallets","header":{"method":"GET","Authorization": "Bearer "}}`))
	_, err = conn.Write([]byte(`{"location":"restaurants","header":{"method":"POST","Authorization": "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTk2NzI1OTAsIlVzZXJJRCI6OCwiSXNBZG1pbiI6ZmFsc2V9.0wvaFlvyxQYulpqOZuY1Y9gtUpgrcU9orkVhxolDO-ffWGEVs0xmGEjaD9735xg7ns15jChtL79_2NwKhTVPKw"}, "data":{"name":"NewRest","phone":"4254578", "address":{"address_line":"borje milad", "coordinates":{"lat":35.7448, "lng":51.3755},"city":"tehran"}}}`))
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
