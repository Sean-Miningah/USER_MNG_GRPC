build: 
	go build -o bin/main main.go 

run: 
	go run main.go


protoc: 
	protoc --go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		usermgmt/usermgmt.proto

protoc_server: 
	go run server/server.go

protoc_client: 
	go run client/client.go