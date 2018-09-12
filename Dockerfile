FROM golang:1.11 as builder

WORKDIR /go/src/github.com/bitbrewers/lappy

RUN go get github.com/golang/dep/cmd/dep
ADD Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only

COPY cmd cmd
COPY Makefile *.go ./
RUN make build

FROM gcr.io/distroless/base
COPY --from=builder /go/src/github.com/bitbrewers/lappy/builds/lappy /lappy

ENTRYPOINT [ "/lappy" ]
