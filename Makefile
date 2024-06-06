server:
	go run ./cmd/api/main.go

evans:
	evans .\api\proto\chat.proto -p 8080
gen:
	protoc -I=api/proto/ --go_out=./ --go-grpc_out=./ api/proto/*.proto
