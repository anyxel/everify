<p align="center">
    <img src="https://cdn.shouts.dev/media/146/real-email-verify.png" alt="logo" width="100"/><br>
    <b>eVerify</b><br>
    Real Email Verify
</p>

## Features

* Syntax checker
* Domain verification
* MX (Mail Exchange records) verification

## Installation

To install `eVerify`, simply run:

```bash
go get github.com/anyxel/everify
```

Run

```bash
go run -ldflags "-X main.version=v1.0.0 -X main.build=31122023" ./app
```

## Build & Run

Linux:

```bash
go build -o everify -ldflags "-X main.version=v1.0.0 -X 'main.build=$(date)'" ./app

./everify
```

Windows:

```bash
go build -o everify.exe -ldflags "-X main.version=v1.0.0 -X 'main.build=$(date)'" ./app

./everify
```

Available commands

```bash
# linux
./everify -h

# windows
./everify.exe -h
```