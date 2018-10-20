#!/bin/bash
set -e

export SERIAL_PORT=$(awk 'NR==1 && NF>1 {print $NF; exit}' <(socat -d -d pty,raw,echo=0 'exec:tranx2sim -i 5000 -t 5 -j 2000 ,pty,raw,echo=0' 2>&1))

if [[ -z "${SERIAL_PORT}" ]]; then
    echo "could not set SERIAL_PORT"
    exit 1
fi

find ./ -type f -name '*.go' | entr -r go run cmd/lappy/main.go
