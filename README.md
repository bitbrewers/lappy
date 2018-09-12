# Lappy

Lappy is a backend for AMB TranX-2 timing system.

[![Build Status](https://travis-ci.com/bitbrewers/lappy.svg?branch=master)](https://travis-ci.com/bitbrewers/lappy)
[![GoDoc](https://godoc.org/github.com/bitbrewers/lappy?status.svg)](https://godoc.org/github.com/bitbrewers/lappy)
[![codecov](https://codecov.io/gh/bitbrewers/lappy/branch/master/graph/badge.svg)](https://codecov.io/gh/bitbrewers/lappy)
[![Go Report Card](https://goreportcard.com/badge/github.com/bitbrewers/lappy)](https://goreportcard.com/report/github.com/bitbrewers/lappy)

## Features
- [x] Read passing and noise events from serial device
- [x] Store passing events in sqlite database with race id
- [x] Publish noise and passing events as [server-sent events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events)
- [ ] Associate metadata with race records
- [ ] Systemd files
- [ ] HTTP API
    - [x] Starting and stopping recoring and publishing passing events
    - [x] Subscribe to SSE
    - [ ] CRUD for metadata like drivers names and car numbers
    - [ ] Fetch records per race
    - [ ] Fetch top records


## Configuration

Required environment variables:

- `LOG_LEVEL` : debug|info|error
- `DSN` : sqlite3 DSN eg. `file:test.db`
- `LISTEN_ADDR` : address to listen eg. `0.0.0.0:8000`
- `SERIAL_PORT` : path to serial device eg. `/dev/usbTTY11`


## Running with docker using local serial device

Create tunnel with socat

```bash
socat -d -d PTY,raw,echo=0 PTY,raw,echo=0
2018/09/12 13:33:11 socat[31301] N PTY is /dev/pts/4
2018/09/12 13:33:11 socat[31301] N PTY is /dev/pts/5
2018/09/12 13:33:11 socat[31301] N starting data transfer loop with FDs [5,5] and [7,7]
```

Start container and mount /dev/pts to access just created `/dev/pts/4`

```bash
docker run --rm -it \
    -v /dev/pts:/dev/pts \
    -p 8000:8000 \
    -e LOG_LEVEL='debug' \
    -e DSN='file:test.db' \
    -e LISTEN_ADDR=':8000' \
    -e SERIAL_PORT='/dev/pts/4' \
    bitbrewers/lappy:latest
```

Install simulator and start it to create trafic

```bash
go get github.com/bitbrewers/tranx2/cmd/tranx2sim
tranx2sim -i 10000 -t 5 -j 2000 > /dev/pts/5
```

## Running locally with simulator

Instructions for basic linux setup

### Requirements

- socat
- go
- tranx2sim

### Steps

1. Install socat and go from repos
2. Install tranx2sim: `go get github.com/bitbrewers/tranx2/cmd/tranx2sim`
3. Use socat for piping serial data: `socat PTY,raw,link=ttyS10,echo=0 PTY,raw,link=ttyS11,echo=0`
4. Start simulator `tranx2sim -i 10000 -t 5 -j 2000 > ./ttyS10`
5. Start lappy with `SERIAL_PORT=./ttyS11`

## Development

Clone repo and run `make dev`.
For more check the Makefile.
