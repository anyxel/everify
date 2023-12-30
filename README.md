<p align="center">
    <img src="https://cdn.shouts.dev/media/146/real-email-verify.png" alt="logo" width="100"/><br>
    <b>eVerify</b><br>
    Real Email Verify
</p>

## Features

* Syntax checker
* Domain verification
* MX (Mail Exchange records) verification

## Use in Go Project

To install `eVerify`, simply run:

```bash
go get github.com/anyxel/everify

import (
  "github.com/anyxel/everify"
)

everify.run('email@exampl.com')
```

## Git Clone & Run

```bash
git clone https://github.com/anyxel/everify.git
cd everify
go run -ldflags "-X main.version=v1.0.0 -X main.build=31122023" ./app
```

## Build & Run

All possible options:

```bash
go tool dist list
```

Example Linux/amd64:

```bash
env GOOS=linux GOARCH=amd64 go build -o everify -ldflags "-X main.version=v1.0.0 -X 'main.build=$(date)'" ./app

./everify -h
```

Example Windows/amd64:

```bash
env GOOS=windows GOARCH=amd64 go build -o everify.exe -ldflags "-X main.version=v1.0.0 -X 'main.build=$(date)'" ./app

./everify.exe -h
```