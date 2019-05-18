proto:
	protoc -I=$(GOPATH)/src/mymicor/mymicor-user/proto --go_out=plugins=micro:$(GOPATH)/src/mymicor/mymicor-user/proto \
		$(GOPATH)/src/mymicor/mymicor-user/proto/user.proto

docker:
	docker build -t mymicor-user .

run:
	docker run -d -p 50053:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns mymicor-user

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
