Unified Node Protocol (Protobufs)

What’s here
- unified_node_protocol.proto: messages and services for Orchestrator and Marketplace
- buf.yaml and buf.gen.yaml for linting and Go codegen targets

Dev
- Lint: `buf lint proto`
- Generate (Go example): `buf generate proto`

Repo structure
- Keep language-specific stubs in each downstream repo (hub-orch/internal/genproto, mobile, etc.)
- This repo is protocol-only to simplify versioning and CI

