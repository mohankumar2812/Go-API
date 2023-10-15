## Description

To get started with Gin, you'll typically create a new Go project, import the Gin package, define your routes, middleware, and handlers, and then run the application.

I develope the user registration for api for login and update the user details with using database mongoDB.

## GO Installation

[GO Install] -- Follow this link to install golang in machine.

## Run 

Run our go API:

```sh
go run main.go
```

** To build and run: ** 

Build the script:

```sh
    go build -o build/bin 
```

Run the build file:

```sh
    ./build/bin
```

## APIs

| Usage | API |        
| -------- | -------- |
| creat | http://localhost:8080/createUser |
| Update | https://localhost:8080/updateUser |
| Delete | https://localhost:8080/deleteUser |
| Login | https://localhost:8080/login |
| GetUser | https://localhost:8080/getUser|