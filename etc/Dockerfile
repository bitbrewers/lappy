FROM golang:1.11

RUN apt-get update && apt-get install -y socat entr gawk
RUN go get github.com/bitbrewers/tranx2/cmd/tranx2sim

WORKDIR /lappy

COPY cmd cmd
COPY go.mod *.go ./
RUN go mod vendor

COPY etc/entrypoint.sh /entrypoint.sh

ENTRYPOINT [ "/entrypoint.sh" ]
