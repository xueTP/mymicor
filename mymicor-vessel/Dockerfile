# vessel-service/Dockerfile
FROM golang:1.9.0 as builder

WORKDIR /go/src/mymicor/mymicor-vessel

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/mymicor/mymicor-vessel/mymicor-vessel .

CMD ["./mymicor-vessel"]