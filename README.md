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

## Notes

- This is for learning and authorized testing only.
- Large scans can take time depending on host/network conditions.
