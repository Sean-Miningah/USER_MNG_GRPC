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