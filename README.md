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

This system models a Transaction Manager (TM) for a distributed system that implements the two-phase commit (2PC) protocol to coordinate stateful changes between participants with interdependencies.

### Components

- Transaction Manager (TM): A gRPC server that implements one procedure, PerformOperation. This takes a string `test` as input and returns a boolean, `success`, and a string `message`. In the real system, any data related to the procedure should be passed and returned to and from the system.

- Participants: The services A, B, and C are the participants of the PerformOperation procedure. They are represented as function calls that simulate a call to a service. In the real system, the participants will vary depending on the procedure being called.

- Database: Services A, B, and C will read and write to the database. This is not implemented in this system, however, the services will output messages to the standard output stream describing what they are simulating (ie. `Service A is accessing the database`).

- Client: The client is external to the system an will invoke a remote procedure call. An application like `Kreya` can be used to simulate a client.

### Handling Transactions

**General Flow of Service Calls and Messages**

![sequence diagram](./TM_sequence_diagram.drawio.png)

Steps:

1. Request received.
2. Phase 1: Process request and message all participants to prepare.
3. Wait for votes: All participants should return a `yes` or `no`.
4. Phase 2: Message participants to commit or abort their operation depending on their votes.
5. Wait for acknowledgement and return response.

_The actions performed on the database by the participants (services A, B, and C) after receiving a commit/abort message from the TM are ommited as they would crowd the diagram. It should be similar to the earlier calls to the database._

#### Success Scenarios

All participants vote `yes`: If all participants return a vote of `yes`, the TM will proceed with messaging all participants to `commit`. Once an acknowledgement is received from all participants, return the success response.

#### Failure Scenarios

At least one participant votes `no`: If one or more of the participants returns a vote of `no`, the TM will proceed with messaging all participants to `abort`. Once an acknowledgement is recieved from all participants, return the failure response.

## Testing and Verification

In order to test the correctness of the system, the PerformOperation procedure can be called, passing it a string of 3 bits (eg. "010", "111", "100", etc). The first character represents the outcome for service A, the second character for service B and the third for service C. Having a "1" indicates that the service will succed, whereas a "0" indicates that the service will fail.

    Example inputs and expected response:

    "010" --> fail: service A votes `no`, service B votes `yes`, service C votes `no`
    "111" --> success: service A votes `yes`, service B votes `yes`, service C votes `yes`

Now it is easy to see that given a request of "111" is the only success scenario (ie. all participants vote `yes`). In testing the system, all valid inputs were tested and a successful response was returned only for the "111" case, validating that the system functions correctly (however, I omitted handling invalid inputs).

## Technical Debt and Future Work

This is hard to determine without a deeper understanding of the current implementation and architecture being used. However, the 2PC protocol is very simple, and should be easy to implement.

### Advantages of 2PC

- Guarantees atomicity as either all commit or all abort.
- Provides fault tolerance if any participant fails.
- A very simple and straightforward solution.

### Drawbacks of 2PC

- Due to the 2 phases, if one of the service is delayed or fails, it can lead to blocking other participants and requests. This may lead to a decrease in performance (decrease in throughout of handled request).
- Increased overhead due to requiring a TM to coordinate participants. Once again, this may lead to a decrease in performance.
- Since the TM corrdinates all participants, the TM acts as a single point of failure. If the TM fails, the entire protocol becomes unavailable. This could lead to transaction failures or indefinite waiting times.
- Scalability may be hindered, since a procedure with many more than 3 participants will require that each participant vote and acknowledge in each of the two phases. The coordination overhead and blocking behaviour will be drastically increased as the number of participants in a procedure increases.

### Alternative Protocols

3-Phase Commit: Introduces an additional "prepare-to-commit" phase between the prepare and commit phases. This extra phase allows for a node to send a "yes" or "no" message, indicating its ability to commit. It helps in reducing the blocking behavior by introducing a timeout mechanism, ensuring progress even in the presence of failures.
