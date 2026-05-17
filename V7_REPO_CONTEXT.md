# V7 Repo Context

## Ryvion V7 Category

Ryvion is a verified execution fabric and compute object web, not just a GPU marketplace. V7 protocol work should support verifiable execution, object-addressed evidence, capability-aware routing, and gradual rollout across the Ryvion system.

## Proto Repo Role

This repository defines cross-service protocol contracts between:

- `ryvion-hub`
- `ryvion-node`
- web/SDKs
- future POP/relay/cache services

The proto repo only defines contracts. It does not implement business logic.

## Existing Protocol

`unified_node_protocol.proto` is the current/legacy v1 protocol. It includes node registration, heartbeat, work assignment, receipts, marketplace messages, and service messages.

The existing v1 protocol must not be broken. Do not remove or rename existing fields, messages, enum values, or RPCs unless an explicit migration task authorizes it.

## V7 Protocol Philosophy

- Additive only.
- Use versioned packages for V7 contracts.
- Do not include raw prompt or transcript fields in the evidence protocol.
- Use hashes, CIDs, and object IDs for evidence references.
- Do not put secrets into messages.
- Design V7 messages so they support gradual rollout and coexist with v1 clients.
- Prefer small, readable proto files grouped by V7 concern.

## Planned V7 Proto Files

- `v7alpha/common.proto`
- `v7alpha/capability.proto`
- `v7alpha/network.proto`
- `v7alpha/model_lease.proto`
- `v7alpha/evidence.proto`
- `v7alpha/objectcdn.proto`
- `v7alpha/heartbeat.proto`

## Ownership Boundaries

`ryvion-hub` owns:

- RCOG
- RYV3GraphReceipt
- ExecutionPlan / RoleSlot
- RiskGate
- auctions
- PathOracle
- ObjectCDN/FEC planning
- audit/settlement

`ryvion-node` owns:

- capability passport
- network profile
- model lease local state
- local CAS
- artifact manifest
- sandbox policy
- evidence payload
- proofrunner bridge

Proto messages define the contract between these owners. They should not encode `ryvion-hub` or `ryvion-node` implementation details beyond the data required for stable interoperability.
