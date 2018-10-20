FROM golang:1.11-alpine as builder
RUN apk --no-cache add make git gcc musl-dev

WORKDIR /lappy

COPY cmd cmd
COPY Makefile *.go go.mod go.sum ./
RUN make build

FROM gcr.io/distroless/base
COPY --from=builder /lappy/builds/lappy /lappy

ENTRYPOINT [ "/lappy" ]
