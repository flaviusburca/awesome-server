# AwesomeServer

This project demonstrates a client-server application in Go using RabbitMQ for message queuing. 
The server handles commands sent by clients to manipulate an in-memory ordered map and dump the values to a file if needed.

## Project Structure

- `client/`: Contains client code
- `server/`: Contains server code
- `common/`: Contains common code shared between client and server
- `cmd/`: Contains the main entry points for the server and client
- `build/`: Contains the binaries for the server and client

## Assumptions
- The ordered map data structure provides O(1) complexity for add, delete, and get operations.
- The project is designed to scale by adding more clients and handling commands in parallel on the server.

## Design Considerations
- The server processes messages from the queue in parallel using goroutines to maximize throughput. Some tuning can be done using the GOMAXPROCS parameter.
- The ordered map is implemented using a combination of a map and a slice to maintain insertion order.
- The client reads commands from standard input and sends them to the server via RabbitMQ.
- The server creates files to store the output of the `getItem` and `getAllItems` commands if needed with the `StandardCommandHandle`
- The `LoggingCommandHandler` logs the commands received by the server and will only log the output of `getItem` and `getAllItems` commands

# Build and Run the project

## Requirements

- Docker
- Go 1.19
- RabbitMQ

## Run the Application

1. Clone the repository and build the project. This will create binaries for `amd64` and `arm64` architectures in the `build` folder.
   ```shell
   git clone
   cd awesome-project
   make
   ```
   
2. Start the RabbitMQ server:
3. Run AwesomeServer:
    ```shell
   ./build/server-amd64
    ```
   
4. Run AwesomeClient and send some commands to the server:

    ```shell
   ./build/client-amd64
   addItem key1 value1
   addItem key2 value2
   addItem key3 value3
   getItem key1
   deleteItem key1
   getAllItems
    ```
   
# Run tests and benchmarks

```shell
cd server
go test .
go test -bench .
```

