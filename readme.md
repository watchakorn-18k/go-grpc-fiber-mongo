# GRPC

- `proto` for proto file
- `server` for connect to database
- `client` for connect to server with grpc normal
- `client-fiber` for connect to server with fiber
- `restful` for connect to server with restful


### `proto` build
```
proto/gen.cmd
```

### `server` run
```
cd server
go mod 
go run main.go
```

## Clients

### `client` run
```
cd client
go mod 
go run main.go
```

### `client-fiber` run
```
cd client-fiber
go mod 
go run main.go
```
GET http://localhost:3000/user/test_1234

### `restful` run
```
cd restful
go mod 
go run main.go
```
GET http://localhost:3000/user/test_1234