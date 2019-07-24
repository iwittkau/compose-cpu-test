FROM golang:1.12-alpine

ARG GOPROXY

WORKDIR /build

ADD . /build

RUN GOPROXY=$GOPROXY CGO_ENABLED=0 go build -o load main.go

FROM scratch

COPY --from=0 /build/load /load

ENTRYPOINT [ "/load" ]
