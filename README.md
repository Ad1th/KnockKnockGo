# KnockKnockGo

A simple TCP port scanner written in Go.

## What It Does

- Scans TCP ports 1-10000
- Uses `localhost` by default
- Accepts a custom host with `-server`
- Shows response time for each open port
- Detects many common services by port
- Prints open ports and a final open-port count

## Run

```bash
go run main.go
```

Use a custom host:

```bash
go run main.go -server scanme.nmap.org
```

Scan a custom range with a worker pool:

```bash
go run main.go -server scanme.nmap.org -start 1 -end 2000 -workers 200 -timeout-ms 400
```

Flags:

- `-server`: target host (default `localhost`)
- `-start`: start port (default `1`)
- `-end`: end port (default `10000`)
- `-workers`: number of concurrent workers (default `runtime.NumCPU()*8`)
- `-timeout-ms`: TCP dial timeout per port in milliseconds (default `500`)

## Notes

- This is for learning and authorized testing only.
- Large scans can take time depending on host/network conditions.
