# ryvion-protocol

Canonical protobuf definitions and generated Go packages for Ryvion.

Active direction:

- `ryvion/node/v1`: node-to-hub gRPC control stream
- `gen/go`: committed generated Go packages consumed by hub and node

Legacy `v7alpha`, speculative, experiment, and DePIN-era contracts have been
removed from the active protocol module. New product contracts should stay
focused on managed render/media work orchestration.

Development:

```bash
buf lint .
buf generate .
go test ./...
```
