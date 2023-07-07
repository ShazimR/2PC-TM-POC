# 2PC-TM-POC

Take Home Assignment for Parametric Research

## Prerequisites

This project requires Go version 1.18 to run the server. If you want to recompile the `transaction.proto` file, then you will also need the `protoc` compiler.

## Usage

To build the project use the command:

```bash
$ go mod download
$ make build
```

To run it, an executable called `2pc-tm-poc.exe` should have been created after building. Run it with:

```bash
$ ./2pc-tm-poc.exe
```

## System Description

TODO

### Components

TODO

### Handling Transactions

General Flow of Service Calls and Messages

![sequence diagram](./TM_sequence_diagram.drawio.png)

#### Success Scenarios

TODO

#### Failure Scenarios

TODO

## Technical Debt and Future Work

TODO
