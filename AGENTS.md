# Codex Guidance

This repository is the canonical protocol module. Keep changes focused on protobuf contracts, generated Go stubs under `gen/go`, and local protocol documentation.

## Protocol Rules

- Do not break the existing v1 protocol.
- V7 protocol changes must be additive.
- Do not remove or rename existing fields.
- Current architecture keeps canonical Go protobuf output in `gen/go`; do not generate protobuf stubs into hub/node repos.
- Do not hand-edit generated files under `gen/go`; change `.proto` files and run `buf generate .`.
- Prefer multiple small proto files for V7 readability.
- Do not modify generated files, build scripts, `buf.yaml`, or `buf.gen.yaml` unless the current task explicitly allows it.

## V7 Tasks

For V7 tasks, always read:

- `V7_REPO_CONTEXT.md`
- `tasks/v7/TASK_INDEX.md`
- the current task file

Use versioned packages for new V7 contracts and keep messages safe for gradual rollout across `ryvion-hub`, `ryvion-node`, web/SDKs, and future POP/relay/cache services.

## Validation

Run these commands when the required tools and Buf configuration are available:

- `buf lint .`
- `buf generate .`
