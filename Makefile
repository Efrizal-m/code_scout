proto:
	protoc --go_out=pb --go-grpc_out=pb --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative internal/lottery/lottery.proto

run-server:
	go run cmd/server/main.go

run-client:
	go run cmd/client/main.go