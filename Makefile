OUT_DIR=frontend/js
proto:
	protoc -I api/rpc/ --go_out=api/rpc --go-grpc_out=api/rpc api/rpc/rpc.proto

proto_js:
	protoc -I api/rpc/ --js_out=import_style=commonjs,binary:${OUT_DIR} --grpc-web_out=import_style=commonjs,mode=grpcwebtext:${OUT_DIR} api/rpc/rpc.proto

build:
	npm install
	docker build -t yt-dl-rpc .

server: build
	mkdir -p $HOME/download-yt
	docker run -d -p 3033:3033 -v $HOME/download-yt:/downloads yt-dl-rpc

go-client:
	go run client/main.go

js-client:
	node frontend/js/test.js
