
domain:
	protoc --proto_path=proto --go_out=. --go-grpc_out=. --go_opt=module=github.com/shiva0612/grpc-proto --go-grpc_opt=module=github.com/shiva0612/grpc-proto  proto/*.proto
test:
	go test ./bench/... -run=Benchmark -bench=. -benchtime=2s -count=5 -benchmem
