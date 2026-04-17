# KnockKnockGo

A simple TCP port scanner written in Go.

## What It Does

- Scans TCP ports 1-10000
- Prompts you for a target host
- Prints open ports and a final open-port count

## Run

```bash
go run main.go
```

Example input:

```text
Enter host (default localhost): scanme.nmap.org
```

If you just press Enter, it scans `localhost`.

## Notes

- This is for learning and authorized testing only.
- Large scans can take time depending on host/network conditions.
