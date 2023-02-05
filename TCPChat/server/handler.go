package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func HandleConnections(connection net.Conn) {
	mutex.Lock()
	UserName := GetName(connection)
	mutex.Unlock()
	User := Clients{UserName, connection}
	mutex.Lock()
	openConnections[UserName] = User
	mutex.Unlock()
	mutex.Lock()
	if len(openConnections) > 10 {
		fmt.Fprintf(User.conn, "Chat is full\n")
		delete(openConnections, UserName)
		connection.Close()
		mutex.Unlock()
		return
	}
	mutex.Unlock()
	mutex.Lock()

	joined := Message{
		msg:      "has joined our chat...",
		time:     time.Now().Format("2006-01-02 15:04:05"),
		UserName: User.UserName,
	}
	joinConnection <- joined

	mutex.Unlock()
	// fmt.Fprintf(connection, "[%s][%s]:", time.Now().Format("2006-01-02 15:04:05"), UserName)
	input := bufio.NewScanner(connection)
	for input.Scan() {
		text := strings.Trim(input.Text(), " ")
		if !checkText(text) {
			fmt.Fprintln(connection, "The empty messages are prohibited")
			fmt.Fprintf(connection, "[%s][%s]:", time.Now().Format("2006-01-02 15:04:05"), UserName)
			continue
		}
		messageNew := Message{
			msg:      text,
			time:     time.Now().Format("2006-01-02 15:04:05"),
			UserName: User.UserName,
		}
		mutex.Lock()
		hst := fmt.Sprintf("[%s][%s]:%s", messageNew.time, messageNew.UserName, messageNew.msg)
		history = append(history, hst)
		Connection <- messageNew
		mutex.Unlock()
	}
	mutex.Lock()
	delete(openConnections, UserName)
	mutex.Unlock()
	text := Message{
		msg:      "has left our chat...",
		time:     time.Now().Format("2006-01-02 15:04:05"),
		UserName: User.UserName,
	}
	mutex.Lock()
	deadConnection <- text
	connection.Close()
	mutex.Unlock()
}
