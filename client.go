package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/rpc"
	"os"
	"strings"
)

type File struct {
	Sender  string
	Name    string
	Content string
}

func printMenu() {
	fmt.Println("0) Salir")
	fmt.Println("1) Enviar Mensaje")
	fmt.Println("2) Enviar Archivo")
	fmt.Println("3) Mostrar Chat")
	fmt.Println("4) Mostrar Archivos")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Nickname: ")
	var nickname string
	fmt.Scanln(&nickname)

	c, err := rpc.Dial("tcp", "127.0.0.1:4242")
	if err != nil {
		panic(err)
	}

	var opc int64
	for {
		printMenu()
		fmt.Scanln(&opc)

		switch opc {
		case 1:
			var msg string
			fmt.Print("Mensaje: ")
			scanner.Scan()

			msg = nickname + ": " + scanner.Text()

			var result int
			err = c.Call("Server.SendMsg", msg, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}
		case 2:
			fmt.Print("Path: ")
			scanner.Scan()
			path := scanner.Text()
			dat, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}

			strs := strings.Split(path, "/")

			file := File{
				Sender:  nickname,
				Name:    strs[len(strs)-1],
				Content: string(dat),
			}

			var result int
			err = c.Call("Server.SendFile", file, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}
		case 3:
			var result string
			err = c.Call("Server.GetChat", "", &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Chat:\n", result)
			}
		case 4:
			var result string
			err = c.Call("Server.GetFiles", "", &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("files:\n", result)
			}
		case 0:
			return
		default:
			fmt.Println("Opcion no valida")
		}

	}

	// Enviar Mensaje
	// Enviar Archivo
	// Mostrar Chat

}
