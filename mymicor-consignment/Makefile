proto:
	protoc -I=$(GOPATH)/src/mymicor/mymicor-consignment/proto --go_out=plugins=micro:$(GOPATH)/src/mymicor/mymicor-consignment/proto \
		$(GOPATH)/src/mymicor/mymicor-consignment/proto/consignment.proto

docker:
	docker build -t mymicor-service .

run:
	docker run -d -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns mymicor-service

build-linux:
	set CGO_ENABLED = 1
	set GOARCH=amd64
	set GOOS=linux
	go build

build:
	set CGO_ENABLED = 0
	set GOARCH=amd64
	set GOOS=windows
	go build
