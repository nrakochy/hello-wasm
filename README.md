# Wasm

To run the server, `cd web && go run server.go` -> `localhost:8333`

To recompile the wasm, `tinygo build -o ./web/public/threader.wasm -target wasm ./main.go`