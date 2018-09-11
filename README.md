# Lappy

Lappy is a backend for AMB TranX-2 timing system.

[![Build Status](https://travis-ci.com/bitbrewers/lappy.svg?branch=master)](https://travis-ci.com/bitbrewers/lappy)
[![GoDoc](https://godoc.org/github.com/bitbrewers/lappy?status.svg)](https://godoc.org/github.com/bitbrewers/lappy)
[![codecov](https://codecov.io/gh/bitbrewers/lappy/branch/master/graph/badge.svg)](https://codecov.io/gh/bitbrewers/lappy)
[![Go Report Card](https://goreportcard.com/badge/github.com/bitbrewers/lappy)](https://goreportcard.com/report/github.com/bitbrewers/lappy)

## Features
- Read passing and noise events from serial device
- Store passing events in sqlite database with race id
- Publish noise and passing events as [server-sent events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events)
- HTTP API
    - Starting and stopping recoring and publishing passing events
    - Subscribe to SSE

- TODO:
    - CRUD for metadata like drivers name and car number
    - Fetch records per race
    - Associate metadate to race records
    - Systemd files

## Running

Required environment variables:

- `LOG_LEVEL` : debug|info|error
- `DSN` : sqlite3 DSN eg. `file:test.db`
- `LISTEN_ADDR` : address to listen eg. `0.0.0.0:8000`
- `SERIAL_PORT` : path to serial device eg. `/dev/usbTTY11`
