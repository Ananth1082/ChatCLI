
# ChatCLI

A Command Line Program where you can create rooms and chat with people.

The application takes advantage of Go's concurrency to handle many requests at a time to create a robust chat app.

## Pre-req

Need to install netcat to connect to the server

## How to use

- Clone the repo
    ``` 
    git clone https://github.com/Ananth1082/ChatCLI.git
    ```

- Run the program
    ```
    cd main
    go run main.go
    ```
- Connect to the server by using its local IP address
    ``` 
    netcat <local_ip_address> <port>
    ```
- Start chatting :)

Please free to point out bugs or suggest improvments to the app.
