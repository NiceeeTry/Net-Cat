package server

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

func GetName(connection net.Conn) string {
	f, err := os.Open("Welcome/welcome.txt")
	r, err := ioutil.ReadAll(f)

	fmt.Fprintf(connection, string(r))

	if err != nil {
		fmt.Fprintf(connection, "Welcome logo cannot be presented(\n")
		fmt.Fprintf(connection, "[ENTER YOUR NAME]:")
	}

	reader := bufio.NewReader(connection)
	username, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(connection, "There is an issue in reading your message(\n")
		GetName(connection)
	}

	username = strings.Trim(username, " \r\n")

	if username == "" || len(username) == 0 {
		fmt.Fprintf(connection, "The username is the necessary condition to enter the chat\n")
		return GetName(connection)
	}

	for _, simbol := range username {
		if simbol < 32 || simbol > 127 {
			fmt.Fprintln(connection, "Incorrect input\n")
			// fmt.Fprintf(connection, "[%s][%s]:", time, username)
			return GetName(connection)
		}
	}

	for names, _ := range openConnections {
		if username == names {
			fmt.Fprintln(connection, "Username is already taken\n")
			return GetName(connection)
		}
	}
	return username
}
