FROM golang:1.11

RUN apt-get update && apt-get install -y socat entr gawk
RUN go get github.com/bitbrewers/tranx2/cmd/tranx2sim

WORKDIR /lappy
COPY etc/entrypoint.sh /entrypoint.sh

COPY go.mod go.sum ./
COPY vendor vendor
COPY . ./

ENTRYPOINT [ "/entrypoint.sh" ]
