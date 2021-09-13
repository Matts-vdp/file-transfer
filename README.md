# file-transfer
A Go program that provides a server and client to send files from the server to all connected client instances.
It uses a tcp connection to send the file to the client.

## Usage
### Config
Edit the ip and port in the config file.

### Server
Go to the directory of the executable and run.
```
 file-transfer -s -f <Path to file>
```
When no file path is given, the program will wait for you to type a path and press enter before opening the tcp connection.


### Client
```
 file-transfer
```
