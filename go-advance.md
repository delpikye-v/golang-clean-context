1. Concurrency & Parallelism

Goroutines: lightweight thread → hàng trăm nghìn goroutine trong 1 process.

Channel Patterns:

Fan-in / Fan-out

Worker pool

Pipeline

Context: cancel propagation, timeout, deadline.

sync vs atomic: lock, RWMutex, Once, Cond vs atomic ops.

🔹 2. Architecture & Design

Clean Architecture / Hexagonal Architecture: tách domain, usecase, infra.

Dependency Injection (DI): với wire
 hoặc fx.

Module Split: chia theo bounded context (DDD).

Generics (Go 1.18+): viết repository/service generic, collection utils.

🔹 3. Data & Persistence

Database ORM/driver:

GORM
 (popular)

sqlx
 (lightweight)

Migration: golang-migrate

Advanced DB Pattern: transaction context, optimistic/pessimistic locking.

Caching: Redis client (go-redis), pattern cache-aside, write-through.

🔹 4. Networking & Communication

gRPC: hiệu năng cao, contract-first với Protobuf.

gRPC Gateway: expose REST ↔ gRPC.

Message Queue: Kafka, SQS, NATS → publish/subscribe, event sourcing.

Websocket / SSE: real-time event stream.

🔹 5. Performance & Scaling

pprof / tracing: profile CPU, memory, goroutine leak.

Benchmarking: testing.B (sub-bench, allocs).

Tuning GC: hiểu GOGC, memory pooling (sync.Pool).

Zero-copy I/O: io.Reader, io.Writer, bufio.

🔹 6. Security & Auth

JWT, OAuth2: middleware auth, refresh token rotation.

mTLS: gRPC mutual TLS.

Secret manager: AWS Secrets Manager / Vault.

🔹 7. Deployment & Cloud

Dockerize Go app (multi-stage build → image < 20MB).

Serverless: Go Lambda (AWS).

K8s operator pattern: build controller bằng Go.

CI/CD: GitHub Actions, GitLab CI.

🔹 8. Advanced Patterns

Plugin system: dùng go plugin hoặc dynamic module.

Code generation: go generate, stringer, protobuf.

Reflection & unsafe: khi cần dynamic (ORM, JSON parser).

Event-driven workflow: Saga / Orchestration pattern (AWS Step Functions, Temporal).
