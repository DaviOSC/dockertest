package main

import (
    "bufio"
    "fmt"
    "net"
)

func main() {
    endereco := ":8080" // Alterado para escutar em todas as interfaces
    listener, _ := net.Listen("tcp", endereco)
    defer listener.Close()
    fmt.Println("Servidor TCP rodando na porta", endereco,"...")

    for {
        conn, _ := listener.Accept()
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    fmt.Fprintln(conn, "Conexão estabelecida com sucesso!")
	fmt.Println("Cliente conectado:", conn.RemoteAddr())

    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        text := scanner.Text()
		if text == "exit" {
			fmt.Println("Cliente",conn.RemoteAddr(), "desconectado")
			fmt.Fprintln(conn, "Desconectando...")
			break
		}
		if text != "" {
        	fmt.Println("Recebido do cliente",conn.RemoteAddr(),":", text)
        	fmt.Fprintln(conn, "Servidor recebeu:", text)
		}
    }
}