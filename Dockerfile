FROM golang:1.10

ARG app
ENV APP=$app


#RUN mkdir -p /go/src/github.com/gtfx/microservices/frontend
RUN mkdir -p /go/src/github.com/gtfx/microservices/$APP

#COPY ./frontend/ /go/src/github.com/gtfx/microservices/frontend/
COPY ./$APP/ /go/src/github.com/gtfx/microservices/$APP/

WORKDIR /go/src/github.com/gtfx/microservices/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/$APP ./$APP

FROM alpine
#RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=0 /go/src/github.com/gtfx/microservices/bin/$APP .
