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

	//_, err = conn.Write([]byte(`{"location":"auth/register","header":{"method":"POST"},"data":{"phone":"09537317958","email":"jassem111@yahoo.com","password":"Jaf@123"}}`))
	//_, err = conn.Write([]byte(`{"location":"wallets/cards","header":{"method":"POST","Authorization": "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTkzMjAxNTgsIlVzZXJJRCI6MywiSXNBZG1pbiI6ZmFsc2V9.7dT6MIKbKUOizMbn31yXXFmfc8MKJ3qg4wqmD9QMsuw6dSWtgaJCSFQTQp0UbuiHyRh3IuD0NqegTGFcLuWimQ"},"data":{"card_number":"112233444"}}`))
	//_, err = conn.Write([]byte(`{"location":"wallets/deposit","header":{"method":"POST","Authorization": "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTkzMjAxNTgsIlVzZXJJRCI6MywiSXNBZG1pbiI6ZmFsc2V9.7dT6MIKbKUOizMbn31yXXFmfc8MKJ3qg4wqmD9QMsuw6dSWtgaJCSFQTQp0UbuiHyRh3IuD0NqegTGFcLuWimQ"},"data":{"card_number":"112233444", "amount":10000}}`))
	//_, err = conn.Write([]byte(`{"location":"wallets/withdraw","header":{"method":"POST","Authorization": "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTkzMjAxNTgsIlVzZXJJRCI6MywiSXNBZG1pbiI6ZmFsc2V9.7dT6MIKbKUOizMbn31yXXFmfc8MKJ3qg4wqmD9QMsuw6dSWtgaJCSFQTQp0UbuiHyRh3IuD0NqegTGFcLuWimQ"},"data":{"card_number":"112233444", "amount":8000}}`))
	//_, err = conn.Write([]byte(`{"location":"wallets/cards","header":{"method":"GET","Authorization": "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTkzMjAxNTgsIlVzZXJJRCI6MywiSXNBZG1pbiI6ZmFsc2V9.7dT6MIKbKUOizMbn31yXXFmfc8MKJ3qg4wqmD9QMsuw6dSWtgaJCSFQTQp0UbuiHyRh3IuD0NqegTGFcLuWimQ"}}`))

	//_, err = conn.Write([]byte(`{"location":"users/addresses","header":{"method":"POST","Authorization": "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTk2MjE0OTEsIlVzZXJJRCI6MSwiSXNBZG1pbiI6ZmFsc2V9.-B7_215-0RgnsEpj3l6hCln0XHMkHhp_oXTRKFvBXP6z86bhwaEJ4481SfewHwQ2SORo0ol32yQMdJFvt00png"},"data":{"address_line":"borje milad","coordinates":{"lat":35.7448, "lng":51.3755},"city":"tehran"}}`))
	//_, err = conn.Write([]byte(`{"location":"restaurants","header":{"method":"POST","Authorization": "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTk1Mjk4NjksIlVzZXJJRCI6MTksIklzQWRtaW4iOmZhbHNlfQ.M_8r5XeYn4HGbH2Xvo0KGxH-Ee1XMhO4sT46H2XbBFpiftLHpmsp-akZydZotxWsjoFeOfh0Dso52V72AF5a9w"}, "data":{"name":"sag paz", "phone":"02351", "city":"Tehran", "address":"gkkhdkgh jdgkhgkhdk fgkhkfhg", "coordinates":"(3,4)"}}`))
	//_, err = conn.Write([]byte(`{"location":"restaurants/menus","header":{"method":"POST","Authorization": "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTk1Mjk4NjksIlVzZXJJRCI6MTksIklzQWRtaW4iOmZhbHNlfQ.M_8r5XeYn4HGbH2Xvo0KGxH-Ee1XMhO4sT46H2XbBFpiftLHpmsp-akZydZotxWsjoFeOfh0Dso52V72AF5a9w"}, "data":{"name":"burgurs", "restaurant_id":3}}`))
	//_, err = conn.Write([]byte(`{"location":"restaurants/menus","header":{"method":"GET","Authorization": "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTk1Mjk4NjksIlVzZXJJRCI6MTksIklzQWRtaW4iOmZhbHNlfQ.M_8r5XeYn4HGbH2Xvo0KGxH-Ee1XMhO4sT46H2XbBFpiftLHpmsp-akZydZotxWsjoFeOfh0Dso52V72AF5a9w"}, "data":{"restaurant_id":3}}`))
	//_, err = conn.Write([]byte(`{"location":"restaurants/menu-items","header":{"method":"POST","Authorization": "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTk1Mjk4NjksIlVzZXJJRCI6MTksIklzQWRtaW4iOmZhbHNlfQ.M_8r5XeYn4HGbH2Xvo0KGxH-Ee1XMhO4sT46H2XbBFpiftLHpmsp-akZydZotxWsjoFeOfh0Dso52V72AF5a9w"}, "data":{"menu_id":1, "price":200000, "name":"kabab koobide", "preparation_minutes": 40, "cancellation_penalty_percentage": 9}}`))
	//_, err = conn.Write([]byte(`{"location":"restaurants/menu-items","header":{"method":"GET","Authorization": "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTk1Mjk4NjksIlVzZXJJRCI6MTksIklzQWRtaW4iOmZhbHNlfQ.M_8r5XeYn4HGbH2Xvo0KGxH-Ee1XMhO4sT46H2XbBFpiftLHpmsp-akZydZotxWsjoFeOfh0Dso52V72AF5a9w"}, "data":{"menu_id":1}}`))
	//_, err = conn.Write([]byte(`{"location":"wallets","header":{"method":"GET","Authorization": "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTk1Mjk4NjksIlVzZXJJRCI6MTksIklzQWRtaW4iOmZhbHNlfQ.M_8r5XeYn4HGbH2Xvo0KGxH-Ee1XMhO4sT46H2XbBFpiftLHpmsp-akZydZotxWsjoFeOfh0Dso52V72AF5a9w"}}`))
	//_, err = conn.Write([]byte(`{"location":"restaurants/categories","header":{"method":"POST","Authorization": "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTk2NzQ1MzEsIlVzZXJJRCI6MywiSXNBZG1pbiI6ZmFsc2V9.g0TciO0ehNn2q8S5a_lptkzBlTtGqGK3xoCqFF4zsvhQ-CxPco5E0UGrqFsh-OgWZBGL-STmj9KiqP0HwoD--Q"}, "data":{"restaurant_id":1, "category_ids":[1,2]}}`))
	_, err = conn.Write([]byte(`{"location":"restaurants/categories","header":{"method":"GET","Authorization": "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTk2NzQ1MzEsIlVzZXJJRCI6MywiSXNBZG1pbiI6ZmFsc2V9.g0TciO0ehNn2q8S5a_lptkzBlTtGqGK3xoCqFF4zsvhQ-CxPco5E0UGrqFsh-OgWZBGL-STmj9KiqP0HwoD--Q"}, "data":{"restaurant_id": 1}}`))

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
