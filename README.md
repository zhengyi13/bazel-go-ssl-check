# SSL Cert Expiration Checker

This is a simple go tool to connect to a configured list of host:port pairs via SSL, and print out the expiration date of each cert found.

The assumption is that you'd like to monitor a server for cert expirations, and you'll use this tool to probe the machine(s) in question for an arbitrary list of ports, and print the output.

## Usage

### Configuration

Create a simple YAML config file such as the one found here. Create a top-level key called "probes", and the value of that will be a list of strings. Each string will be a bare, colon-separate "host:port" pair.

```yaml
probes:
  - "localhost:8443"
  - "acme.com:443"
  - "internal.foo:1234"
```

### Invocation

`go run ssl_expiry.go`

## Notes

Inspired by a simple use case at work.
