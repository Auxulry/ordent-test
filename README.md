# Mini E-Commerce API
## Features
- Thorough documentation: Written with the same care as Go docs.
- Guaranteed consistency: Opinionated linting for Go integrated into Visual Studio Code and run against staged files on pre-commit.
## Table Of Contents
- [Features](#features)
- [Getting Started](#getting-started)
- [Environment](#environment)
- [Available Scripts](#available-scripts)
- [Architecture](#architecture)
- [Linting & Formating](#linting--formatting)
- [Customize Configuration](#customize-configuration)
## Getting Started
Make sure you have the following installed:
- Go v1.19 or higher
- golangci-lint v1.50 or higher
- google wire

Import the `database.sql` files to your local database.
```bash
# 1. Clone the repository.
git clone https://github.com/MochamadAkbar/ordent-testt.git
# 2. Enter your cloned directory.
cd ordent-test
# 3. Install dependencies. 
make dependency
# 4. Run on your local.
# This command is a default to run development mode,
# and wil be listen http://localhost:5000
make debug
```

## Environment
For the first time you must create environment in root directory. This List Environment must be create in this project:
- `.env` or `.env.local`

## Available Scripts
```
# 1. Running project in your local machine
make debug
# 2. Build project the output build must be in deploy directory
make build
# 3. Download dependecies used in this source
make dependency
# 4. Check Lint
make lint
# 5. Running all of unit tests
make test
# 6. Running go mod tidy
make tidy
# 7. Wire generate use google wire dependency injection
make wire_gen
```
## Architecture
```
└── api/
└── cmd/
    └── server/
        └── main.go
└── common/
└── config/
└── entity/
└── handler/
└── injector/
└── middleware/
└── repository/
└── usecase/
└── Makefile
```
### api
all of contract http response built in struct
### cmd/server/main.go
main server running
### common
all common utilities in this project
### config
all of configuration for this project
### entity
entity is describe for table identity
### handler
handler is describe for http method deliver
### injector
location for generate file from google wire
### middleware
all of middleware used in this project
### repository
repository is describe all functionality to action database
### usecase
usecase is describe all of business logic in this project

## Linting & Formatting
For this project use golangci-lint to check standard style code
## Customize Configuration
- See [Go Documentation](https://go.dev/doc/)
