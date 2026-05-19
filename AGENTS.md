# Codex Agent Instructions - ryvion-protocol

## Scope

`ryvion-protocol` is the canonical protobuf and generated Go module shared by
`ryvion-hub` and `ryvion-node`.

Active protocol work should focus on:

- node registration, heartbeat, work lease, abort, and receipt transport
- local AI/llama.cpp job contracts
- metering and evidence references
- gRPC `NodeGateway`

## Rules

- Do not generate protobuf stubs in hub or node repos.
- Keep generated Go under `gen/go` in this repo when protocol files change.
- Do not add new V7/V8/Foresight/mesh/plane names to active contracts.
- Archive inactive planning docs, render-farm contracts, and experimental mesh
  contracts instead of expanding them.
- Do not break existing generated packages that hub/node still import; remove
  old contracts only after downstream imports are removed.

## Validation

- `buf lint .`
- `buf generate .`
- `go test ./...`
