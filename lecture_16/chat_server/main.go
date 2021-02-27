package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

//dism /online /Enable-Feature /FeatureName:TelnetClient
/*
type client struct {
	name     string
	conn     net.Conn
	commands chan<- command
}

type command struct {
	id     string
	client *client
	args   []string
}

type server struct {
	commands chan command
}

func newServer() *server {

	return &server{
		commands: make(chan command),
	}

}

func (s *server) run() {

	for cmd := range s.commands {

		switch cmd.id {
		case "setup": //signup
			s.setup(cmd.client, cmd.args[1])
		}
	}
}

func (s *server) newClient(conn net.Conn) {

	c := &client{
		name:     "anonymous",
		conn:     conn,
		commands: s.commands,
	}

	c.readInput()
}

func (c *client) readInput() {

	for {

	}
}

func (s *server) setup(c *client, name string) {

	c.name = name

}

func waitForCmd() {

	ch := make(chan string)

	for v := range ch {
		fmt.Println(v)
		//v := <-ch
	}

}
*/

type user struct {
	name       string
	connection net.Conn
}

func checkSignup(users []user, c net.Conn) bool {

	for _, u := range users {

		if u.connection == c {
			return true
		}
	}
	return false
}

func getIndex(users []user, c net.Conn) int {

	for i, u := range users {

		if u.connection == c {
			return i
		}
	}
	return -1
}

func broadcast(users []user, c net.Conn, msg string) {

	for _, u := range users {
		if u.connection != c {
			indx := getIndex(users, c)
			s := users[indx]
			message := fmt.Sprintf(`%s > %s`, s.name, msg)
			u.connection.Write([]byte(message + "\n"))
			fmt.Println(u.name, u.connection.RemoteAddr().String())
		}
	}

}

func handleMSg(c net.Conn) {

	for {
		msg, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			return
		}
		if err == io.EOF {
			// Connection closed, deregister client
			fmt.Println("connection close")
		}
		fmt.Println(msg)
	}
	//fmt.Println("end of the message")

}

func main() {

	users := make([]user, 0)
	nl, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer nl.Close()
	log.Printf("server started on :8888")

	for {

		conn, err := nl.Accept()
		if err != nil {
			fmt.Println(err.Error())
			//continue
		}

		fmt.Println(conn.RemoteAddr().String(), users)
		bs := make([]byte, 1024)
		n, e := conn.Read(bs)
		if e != nil {
			fmt.Println(e.Error())
		}

		// msg, err := bufio.NewReader(conn).ReadString('\n')
		// if err != nil {
		// 	return
		// }

		//msg = strings.Trim(msg, "\r\n")
		//go handleMSg(conn)

		fmt.Println(n)
		//fmt.Println(bs)
		reqstr := string(bs[:n])
		fmt.Println(reqstr)
		msg := fmt.Sprintf(`Your message: %s`, reqstr)
		conn.Write([]byte(msg))

		//if checkSignup(users, conn) == false {
		//conn.Write([]byte("signup format: signup <username>\n"))
		//continue
		//}

		// switch cmd {

		// case "signup":
		// 	//fmt.Println("signup")
		// 	//members added into map
		// 	msg = fmt.Sprintf(`thank you %s for signup`, args[1])
		// 	users = append(users, user{name: args[1], connection: conn})
		// 	fmt.Println(len(users), users)
		// 	//wait for the command
		// 	conn.Write([]byte(msg + "\n"))

		// default:
		// 	fmt.Println("default")
		// 	//broadcast(users, conn, reqstr)
		// 	conn.Write([]byte("unknown command" + "\n"))
		// }

	}

	//https://developer.mozilla.org/en-US/docs/Web/HTTP/Messages
}
