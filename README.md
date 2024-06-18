# banking-transaction-be

# Overview

Backend restful API to manage bank transaction with 2 roles user including authentication using jwt+session and
authorization based on their roles.
`maker` role can create the transaction but can not `approve` the transaction. the `approver` role with review and
approve or reject the transaction.
beside that the registration is also using otp verification code send by email. and many more features inside! kindly
check!

for the frontend please check: https://github.com/rizanw/banking-transaction-fe

## Features

endpoint list:

- **auth**:
    - **send-otp**: to request otp code
    - **register**: to register new user
    - **login**: to login into the app
    - **logout**: to clean the trace
- **transaction**:
    - **download-template**: to download transaction template in csv
    - **upload**: to upload and submit the transaction data
    - **transactions**: to get all data transactions and a lot of filter & pagination features here.
    - **transaction/{id}**: to get the detail data transaction
    - **transaction/{id}/audit**: to audit the transaction by approver
- **utils**:
    - **corporate** : helper endpoint for registration

further endpoint description, kindly check postman collection.

# Postman Collection

Check the postman collection here:
https://www.postman.com/navigation-candidate-18708542/workspace/banking-transaction/overview

# Database Design

![alt text](./schema/Screenshot%202024-06-18%20at%206.19.32%E2%80%AFPM.png)

# Getting Started

## Prerequisites

Make sure you have installed all the following prerequisites on your development machine:

* go version : [1.19](https://golang.org/dl/)
* docker: https://docs.docker.com/get-docker/

## Local Run Guides:

To clone this repo:

```bash
git clone https://github.com/rizanw/banking-transaction-be.git
```

To setup the environtment:

- to run docker-compose (we need it for local postgres)

```bash 
make docker-up 
```

- to initiate db schema (please make sure the postgres is running properly)

```bash 
./schema/setup.sh
```

- to use email feature please update the config.yaml or add env vars

```bash
export EMAIL=test@mail.com 
export EMAIL_PASSWORD=secret
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
