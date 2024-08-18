package main

import (
	"log"
	"net"
)

func main() {

	log.Println("Listening on port 6369")

	listener, err := net.Listen("tcp", ":6369")
	if err != nil {
		log.Println("Error listening at port 6369 ", err.Error())
		return; 
	}

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection: ", err.Error())
			break; 
		} 
		go handleConnection(connection); 
	}
}

func handleConnection(conn net.Conn){ 
	defer conn.Close(); 
	log.Println("Client connected from: ", conn.RemoteAddr().String())
	
	for {
		var buffer []byte = make([]byte, 100 * 1024); 
		
		_, err := conn.Read(buffer); 

		if err != nil { 
			log.Println("Client disconnected with: ", err.Error()); 
			conn.Close()
			break; 
		}
    	conn.Write([]byte( "+OK " + string(buffer) + "\r\n"))
	}
}
