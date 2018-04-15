FROM golang:1.10

ARG app
ENV APP=$app
ENV GODIR=/go/src/github.com/gtfx/go-microservices/

RUN mkdir -p $GODIR/$APP

COPY glide.* $GODIR
RUN go get github.com/Masterminds/glide

WORKDIR $GODIR
RUN glide install

COPY . $GODIR
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/$APP ./$APP

FROM alpine
ENV GODIR=/go/src/github.com/gtfx/go-microservices/

WORKDIR /root
COPY --from=0 $GODIR/bin/$APP .
