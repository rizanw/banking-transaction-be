# banking-transaction-be

# Overview

Backend restful API to manage bank transaction with 2 roles user including authentication using jwt+session and
authorization based on their roles.
`maker` role can create the transaction but can not `approve` the transaction. the `approver` role with review and
approve or reject the transaction.
beside that the registration is also using otp verification code send by email. and many more features inside! kindly
check!

## Features

endpoint list:

- auth:
    - send-otp
    - register
    - login
    - logout
- transaction:
    - download-template
    - upload
    - transactions
    - transaction/{id}
    - transaction/{id}/audit

further endpoint description, kindly check postman collection.

# Postman Collection

Check the postman collection here:
https://www.postman.com/navigation-candidate-18708542/workspace/banking-transaction/overview

# Database Design
![alt text](./schema/Screenshot%202024-06-18%20at%206.19.32%E2%80%AFPM.png)

# Local Development

## Prerequisites

Make sure you have installed all the following prerequisites on your development machine:

* go version : [1.19](https://golang.org/dl/)

## Local Run Guides:

To clone this repo:

```bash
git clone https://github.com/rizanw/banking-transaction-be.git
```

To build and start the apps:

- build the binaries:

```bash 
make build
```

- start the app:

```bash 
make run
```

- to use email feature please update the config.yaml or add env vars
```bash
export EMAIL=test@mail.com 
export EMAIL_PASSWORD=secret
```

## Unit Test

To run unit test

```bash
make test
```

# Project Structure

- `bin/` is directory for compiled binary
- `cmd/` is the main program directory
- `files/` contains app files (including db & config)
    - `file/etc/tx-bank` contains app config files
- `internal/` contains the whole logic of the app
    - `internal/common` contains helper functions
        - `internal/common/middleware` is for http middleware for client
        - `internal/common/session` is the session manager for auth
    - `internal/config` is the config of the app, has relation to files directory
    - `internal/handler` is application logic interface between this app with client
    - `internal/model` is model business design
    - `internal/repo` is the repositories to fetch/store data of this app
    - `internal/usecase` is main business logic
- `go.mod` the golang dependencies list