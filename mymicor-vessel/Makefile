proto:
	protoc -I=$(GOPATH)/src/mymicor/mymicor-vessel/proto --go_out=plugins=micro:$(GOPATH)/src/mymicor/mymicor-vessel/proto \
		$(GOPATH)/src/mymicor/mymicor-vessel/proto/vessel.proto

docker:
	docker build -t mymicor-vessel .

run:
	docker run -d -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns mymicor-vessel

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
