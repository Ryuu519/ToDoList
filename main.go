package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("folosire: ./main.exe *comanda*")
		os.Exit(1)
	}
	switch args[1] {
	case "add":
		fmt.Println("adaugrea unui task")
	case "del":
		fmt.Println("stergerea unui task")
	case "help":
		fmt.Println("afisarea comenzilor ce pot fi apelate de aplicatie")
	case "edit":
		fmt.Println("editarea unui task")
	default:
		fmt.Println("argument invalid, pentru listarea comenzilor folosit ./todo help")
	}
}
