# Wasm

To run the server, `cd web && go run server.go` -> `localhost:8333`

To recompile the wasm, `tinygo build -o ./web/public/threader.wasm -target wasm ./main.go`


(eventually) Exploring the [tickler pattern explained here](http://www.goldsborough.me/go/2020/12/06/12-24-24-non-blocking_parallelism_for_services_in_go/)