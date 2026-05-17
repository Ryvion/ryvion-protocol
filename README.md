# Ryvion Protocol

What’s here
- unified_node_protocol.proto: messages and services for Orchestrator and Marketplace
- buf.yaml and buf.gen.yaml for linting and Go codegen targets

V7 alpha
- v7alpha/: additive experimental V7 protocol messages for gradual rollout
- unified_node_protocol.proto remains the v1/current protocol
- gen/go/: canonical generated Go module consumed by hub and node

Dev
- Lint: `buf lint .`
- Generate canonical Go stubs: `buf generate .`

Repo structure
- Do not generate protobuf stubs into `ryvion-hub/internal/genproto` or `ryvion-node/internal/genproto`.
- Hub and node should import `github.com/Ryvion/ryvion-protocol/gen/go/...`.
- Generated Go under `gen/go` is committed as the canonical shared module so hub and node use identical contract types.
