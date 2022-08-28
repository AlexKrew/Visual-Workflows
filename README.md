# Visual-Workflows

A visual scripting tool to create workflows and applications by combining nodes.
This project is part of a master project by Alexander Krewinkel and Maurice Falk at RheinMain University of Applied Sciences (Wiesbaden, Germany) in Summer 2022.

## Setup

---

We are using a devcontainer for development, see `./devcontainer`. Open the container using docker and the [Remote Development](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack) extension in Visual Studio Code or follow the introduction [Developing inside a Container](https://code.visualstudio.com/docs/remote/containers).

Install the npm packages by running

```bash
$ cd frontend
$ npm install
```

The code was primarily tested in the Chrome Browser!

### Starting the frontend client

Run the following command inside the `/frontend` directory after installing the npm packages:

```bash
$ npm run serve
```

### Starting a Local Server Instance

For starting a local server instance + webserver, you simple can run following command inside the working directory:

```bash
$ go run main.go
```

### Starting a Remote Client

You can easily spawn an example remote client by running following command from the command line in the working directory:

```bash
$ go run clients/go/main.go
```

## Adding other Remote Clients

---

This repository contains an example client (using Golang) for executing remote services in the `clients/go` directory. If you don't want to use Golang for your client you can simply implement your own client.

The communication between the server and clients is using the [gRPC](https://grpc.io/) protocol and the corresponding `.proto` can be found in `gateway/gateway.proto`.

## File Structure

---

```
.
│
└───clients
│   │
│   └───go               // an example remote client written in golang
│
└───example_services
│   │
│   └───random_service   // an own example service for demonstrating the usage of remote services
│
└───frontend
│   │
│   └───src              // everything regading the editor and dashboard
│
└───gateway
│   │
│   └───gateway_grpc.pb.go   // auto-generated
│   │
│   └───gateway.pb.go        // auto-generated
│   │
│   └───gateway.proto        // protobuf file for generating grpc client/server stubs
│
└───internal            // code regarding the workflow server
│
│
└───shared              // shared code used by the internal packages and the go-client
│
│
└───webserver           // the webserver implementation using the go gin package
│
│
└───workflows           // the file-system storage of the available workflows
```
