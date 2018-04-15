FROM golang:1.10

RUN mkdir -p /go/src/github.com/gtfx/microservices/frontend

COPY ./frontend/ /go/src/github.com/gtfx/microservices/frontend/
WORKDIR /go/src/github.com/gtfx/microservices/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/frontend ./frontend

FROM alpine
#RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=0 /go/src/github.com/gtfx/microservices/bin/frontend .
CMD ["./frontend"]