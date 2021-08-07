# file-transfer
A Go program that provides a server and client to send files from the server to all connected client instances.
It uses a tcp connection to send the file to the client.

## Usage
### Server
Go to the directory of the executable and run.
```
 server.exe -f <Path to file> -ip <ip of server> -p <port>
```
When no file path is given, the program will wait for you to type a path and press enter before opening the tcp connection.


### Client
```
 client.exe -ip <ip of server> -p <port>
```
The default ip and port are: localhost:5000

You can also change the default ip and port of the client or server by modifying the values in the main routine and building a new executable.
