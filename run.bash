GOOS=js GOARCH=wasm go build -o main.wasm
echo "Running..."
go run server/server.go