# KnockKnockGo

A simple TCP port scanner written in Go.

## What It Does

- Scans TCP ports 1-30000 by default
- Supports full scan (1-65535) with `-full`
- Uses `localhost` by default
- Accepts a custom host with `-server`
- Shows response time for each open port
- Detects many common services by port
- Prints open ports and a final open-port count

## Run

### Install

```bash
go install github.com/Ad1th/KnockKnockGo@latest
mv "$(go env GOPATH)/bin/KnockKnockGo" "$(go env GOPATH)/bin/knock"
export PATH="$PATH:$(go env GOPATH)/bin"
```

After that, you can run:

```bash
knock
```

### Run from source

```bash
go run main.go
```

Full scan:

```bash
go run main.go -full
```

Use a custom host:

```bash
go run main.go -server scanme.nmap.org
```

Full scan for a custom host:

```bash
go run main.go -server scanme.nmap.org -full
```

Scan a custom range with a worker pool:

```bash
go run main.go -server scanme.nmap.org -start 1 -end 2000 -workers 200 -timeout-ms 400
```

Flags:

- `-server`: target host (default `localhost`)
- `-start`: start port (default `1`)
- `-end`: end port (default `30000`)
- `-full`: scan all ports (`1-65535`)
- `-workers`: number of concurrent workers (default `runtime.NumCPU()*8`)
- `-timeout-ms`: TCP dial timeout per port in milliseconds (default `500`)

## Notes

- This is for learning and authorized testing only.
- Large scans can take time depending on host/network conditions.
