FROM golang:1.11

RUN apt-get update && apt-get install -y socat entr
RUN go get github.com/golang/dep/cmd/dep github.com/bitbrewers/tranx2/cmd/tranx2sim

WORKDIR /go/src/github.com/bitbrewers/lappy

ADD Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only

COPY cmd cmd
COPY Makefile *.go ./
RUN go install

CMD [ "bash", "-c", "socat -d -d pty,raw,echo=0 'exec:tranx2sim -i 5000 -t 5 -j 2000 ,pty,raw,echo=0' & find ./ -type f -name '*.go' | entr -r go run cmd/lappy/main.go"]
