package server

import (
	"fmt"
	"time"
)

func Broadcaster() {
	for {
		select {
		case msg := <-joinConnection:
			mutex.Lock()
			for _, user := range openConnections {
				if user.UserName != msg.UserName {
					fmt.Fprintf(user.conn, "\n%s %s\n[%s][%s]", msg.UserName, msg.msg, msg.time, user.UserName)
				} else {
					for _, previous := range history {
						// fmt.Fprintf(user.conn, "%s %s\n[%s][%s]:\n", msg.UserName, msg.msg, msg.time, user.UserName)

						fmt.Fprintf(user.conn, previous)
						fmt.Fprintf(user.conn, "\n")
					}
					fmt.Fprintf(user.conn, "[%s][%s]:", time.Now().Format("2006-01-02 15:04:05"), user.UserName)
					// fmt.Fprintf(user.conn, "\n")
					// fmt.Fprintf(user.conn, "\n[%s][%s]:", msg.time, user.UserName)
				}
			}
			mutex.Unlock()
			//
		case message := <-Connection:
			mutex.Lock()
			for _, user := range openConnections {
				if user.UserName != message.UserName {
					fmt.Fprintf(user.conn, "\n[%s][%s]:%s\n", message.time, message.UserName, message.msg)
					// hst := fmt.Sprintf("\n[%s][%s]: %s\n", message.time, message.UserName, message.msg)
					// history = append(history, hst)
				}
				fmt.Fprintf(user.conn, "[%s][%s]:", message.time, user.UserName)
			}
			mutex.Unlock()
		case msg := <-deadConnection:
			mutex.Lock()
			for _, user := range openConnections {
				if user.UserName != msg.UserName {
					fmt.Fprintf(user.conn, "%s %s\n[%s][%s]:", msg.UserName, msg.msg, msg.time, user.UserName)
				}
			}
			mutex.Unlock()
		}
	}
}
