
# Quickly
It is a stub application that can easily prepare a stub environment.
Response information to return can be set in config.json.

## Steps
- `go run main.go`

## Change response
- ./config.json
---
`{
  "routes": [
    {
      "url": "/abc/def/aaa",
      "response": "{\" response \ ": [aaa]}}"
    },
    ...
  ]
} `
