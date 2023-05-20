
# Net-Cat 💬 📨 📲

The project consists of recreating the NetCat in a Server-Client Architecture that can run in server mode on a specified port listening for incoming connections, and it can be used in client mode, trying to connect to a specified port and transmitting information to the server.

The project must work in a similar way that the original NetCat works, in other words, it is necessary to create a group chat. The project has the following features :

## Features 🖊 📤 🫂    
- TCP connection between the server and multiple clients (relation of 1 to many).
- A name requirement for the client.
- Control connections quantity.
- Clients must be able to send messages to the chat 📩.
- Do not broadcast EMPTY messages from a client.
- Messages sent must be identified by the time that were sent and the user name of who sent the message, for example: `[2020-01-20 15:48:41][client. name]:[client.message]`
- If a Client joins the chat, all the previous messages sent to the chat must be uploaded to the new Client.
- If a Client connects to the server, the rest of the Clients must be informed by the server that the Client joined the group.
- If a Client exits the chat, the rest of the Clients must be informed by the server that the Client left.
- All Clients must receive the messages sent by other Clients.
- If a Client leaves the chat, the rest of the Clients must not disconnect.
- If there is no port specified, then set it as default port 8989. Otherwise, the program must respond with a usage message:`[USAGE]: ./TCPChat $port`

## The usage
To `run a server` the following command has to be entered ▶:
```bash
$ go run ./TCPChat/main.go
```
The response has to be this:
`Listening on the port :8989`

To connect to the server as aclient (listener), this command has to be entered:
```bash
nc localhost 8989
```

## The web application 
`Client 1: Yenlik`

![App Screenshot](https://github.com/NiceeeTry/Net-Cat/assets/120025832/d3aaafc5-cddc-4636-a043-14010b082194)

`Client 2: Lee`

![App Screenshot](https://github.com/NiceeeTry/Net-Cat/assets/120025832/d793b6e2-92ea-485e-a505-1be062758351)

`The maximum number of users connected to the TCP chat is 10 users however this number can be adjusted.`



