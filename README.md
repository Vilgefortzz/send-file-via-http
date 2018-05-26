# send-file-via-websocket

Sending files via websocket based on server-client

To run the program, start the server:

    $ go run server.go

Next, start the client with passing arguments:

    $ go run client.go -- arg1 arg2

arg1 - name of the file located in send-files directory (e.g. hello.go)

arg2(optional) - host name (e.g. 127.0.0.1)