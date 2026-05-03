Unified Node Protocol (Protobufs)

What’s here
- unified_node_protocol.proto: messages and services for Orchestrator and Marketplace
- buf.yaml and buf.gen.yaml for linting and Go codegen targets

V7 alpha
- v7alpha/: additive experimental V7 protocol messages for gradual rollout
- unified_node_protocol.proto remains the v1/current protocol
- generated code should be consumed from downstream repos, not committed here

Dev
- Lint: `buf lint .`
- Generate (Go example): `buf generate .`

Repo structure
- Keep language-specific stubs in each downstream repo (hub-orch/internal/genproto, mobile, etc.)
- This repo is protocol-only to simplify versioning and CI
