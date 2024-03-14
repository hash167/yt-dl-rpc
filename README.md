## YT-DL-RPC
YT-DL-RPC is a project that provides a gRPC interface for downloading videos from YouTube. It uses the popular youtube-dl command-line program under the hood and exposes its functionality over gRPC. This allows you to integrate YouTube downloading into your own applications in a language-agnostic way.

## Features
- Download individual videos using grpc clients in any programming language
- Get progress updates for ongoing downloads
- Query available video formats
- Manage multiple concurrent downloads
- Kill or clear downloads
- Check free disk space
- Update the youtube-dl executable

### Getting Started

Prerequisites

- Docker
- Go 1.16 or later
- Node.js 14 or later
- npm 6 or later
- Protocol Buffers compiler (protoc)

#### Building
To build the project, run the following command:

##### generate proto files for clients
This will compile the Protocol Buffers definitions, install the necessary npm packages, and build a Docker image for the server.

- make proto (golang)
- make proto_js (javascript)


#### Running the Server
To start the rpc server, run the following command:
- make build
- make server

##### server
This will start a Docker container with the rpc server. The server listens on port 3033.

#### Running the Clients
The project includes a Go client and a JavaScript client for testing the server.

To run the Go client, use the following command:
- make go-client

To run the JavaScript client, use the following command:
- make js-client

### API
The gRPC API is defined in the api/rpc/rpc.proto file. You can generate client stubs for your preferred language using the Protocol Buffers compiler.

Contributing
Contributions are welcome! Please feel free to submit a pull request.

License
This project is licensed under the MIT License. See the LICENSE file for details.

initialize
DownloadRe
