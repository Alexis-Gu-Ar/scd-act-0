package main

import (
	"fmt"
	"net"
	"net/rpc"
	"os"
)

/*
	"errors"
	"fmt"
	"net"
	"net/rpc"
*/

type File struct {
	Sender  string
	Name    string
	Content string
}

type Message struct {
	sender  string
	content string
}

type Client struct {
	nickname string
}

/*

	// TODO: ask to store in a file

	fmt.Println("Do you want to store the msg and filenames? s/n")
	fmt.Scan(&input)

	if input == "s" {
		file, err := os.Create("./logs.txt")
		defer file.Close()
		if err != nil {
			panic(err)
		}
		fmt.Println(server.log)
		file.WriteString(server.log)
	}
*/
type Server struct {
	clients  []Client
	files    []File
	messages []string
	log      string
	file     *os.File
}

func (s *Server) SendMsg(msg string, reply *int) error {
	fmt.Println(msg)
	*reply = 200
	s.file.WriteString(msg + "\n")
	s.messages = append(s.messages, msg)
	return nil
}

func (s *Server) SendFile(file File, reply *int) error {
	fmt.Println(file.Sender + " sent the file: " + file.Name)
	fmt.Println("with content: ")
	fmt.Println(file.Content)
	*reply = 200
	s.file.WriteString("file: " + file.Name + "\n")
	s.files = append(s.files, file)
	return nil
}

func (s *Server) GetFiles(_ string, reply *string) error {
	result := ""

	for _, file := range s.files {
		result += file.Name + ": \n" + file.Content + "\n\n"
	}
	*reply = result
	return nil
}

func (s *Server) GetChat(_ string, reply *string) error {
	result := ""

	for _, msg := range s.messages {
		result += msg + "\n"
	}
	*reply = result
	return nil
}

func start(server Server) {
	file, err := os.Create("./logs.txt")
	server.file = file

	rpc.Register(&server)
	ln, err := net.Listen("tcp", ":4242")

	if err != nil {
		panic(err)
	}

	for {
		c, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go rpc.ServeConn(c)
	}

}

func main() {
	server := Server{
		clients:  make([]Client, 0),
		files:    make([]File, 0),
		log:      "",
		messages: make([]string, 0),
	}

	go start(server)

	var input string
	fmt.Scanln(&input)
	server.file.Close()
	/*
		1. Mostrar los mensajes/nombre de los archivos enviados. (mostrar al instante y reenviar a clientes)
		2. Opción para respaldar en un archivo de texto los mensajes/nombre de los archivos enviados. (Preguntar antes de terminar)
		3. Terminar servidor. (con enter)
	*/

}
