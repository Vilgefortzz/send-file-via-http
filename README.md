# send-file-via-http

Sending files via http based on server-client

Install Gorilla WebSocket:

    $  go get github.com/gorilla/websocket

To run the program, start the server with passing arguments:

    $ go run server.go -- arg1
    
- arg1(optional) - host name (e.g. 127.0.0.1)

<br>

Next, start the client with passing arguments:

    $ go run client.go -- arg1 arg2

- arg1 - name of the file located in send-files directory (e.g. hello.go)

- arg2(optional) - host name (e.g. 127.0.0.1)

<br>

To run tests from /tests (server must be running):

    $ go test -v
