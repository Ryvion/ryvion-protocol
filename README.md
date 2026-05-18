# ryvion-protocol

Canonical protobuf definitions and generated Go packages for Ryvion.

Active direction:

- `ryvion/node/v1`: node-to-hub gRPC control stream
- `gen/go`: committed generated Go packages consumed by hub and node

Legacy packages remain only while downstream repos still import them. Do not
add new product work to `v7alpha`, speculative, or experiment contracts.

Development:

```bash
buf lint .
buf generate .
go test ./...
```
