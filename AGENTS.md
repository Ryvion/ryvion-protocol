# Codex Guidance

This repository is protocol-only. Keep changes focused on protobuf contracts and local protocol documentation.

## Protocol Rules

- Do not break the existing v1 protocol.
- V7 protocol changes must be additive.
- Do not remove or rename existing fields.
- Do not generate language-specific stubs into this repo unless the current task explicitly asks for it.
- Keep language-specific generated code in downstream repos.
- Prefer multiple small proto files for V7 readability.
- Do not modify generated files, build scripts, `buf.yaml`, or `buf.gen.yaml` unless the current task explicitly allows it.

## V7 Tasks

For V7 tasks, always read:

- `V7_REPO_CONTEXT.md`
- `tasks/v7/TASK_INDEX.md`
- the current task file

Use versioned packages for new V7 contracts and keep messages safe for gradual rollout across `hub-orch`, `node-agent`, web/SDKs, and future POP/relay/cache services.

## Validation

Run these commands when the required tools and Buf configuration are available:

- `buf lint .`
- `buf generate .`
