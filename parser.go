package main

import "net"

type Command struct { 
	Name string
	Args []string
}


func ParseRESP(connection net.Conn, inputBuf []byte) ([]Command, error) {
	return nil, nil 
}