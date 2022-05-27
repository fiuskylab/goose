start-env:
	go env -w GOOS=js GOARCH=wasm

clean-env:
	go env -w GOOS=linux GOARCH=amd64

run:
	GOOS=js GOARCH=wasm go build -o main.wasm
	cp \"$(go env GOROOT)/misc/wasm/wasm_exec.js\" .
	go run cmd/server.go

