# Go Learning Path

A structured, hands-on repo for learning Go, organized into three tiers.
Every file is a **runnable, self-contained demo** with heavy comments —
open it, read top to bottom, then run it.

## How to run any file

Each `.go` file has its own `package main` and `func main()`, so run files
**one at a time**:

```bash
go run "basic overveiw/Loops.go"
go run intermediate/01_Pointers.go
go run advanced/01_Generics.go
```

For the testing example (which is a real package, not a standalone file):

```bash
cd "intermediate/08_Testing"
go test -v ./...
go test -bench=.
```

---

##  Tier 1 — `basic overveiw/` (Fundamentals)

| Topic | File |
|---|---|
| Hello World | `hello.go` |
| Variables & constants | `Variables.go` |
| If / Else | `If-Else.go` |
| Switch | `Switch.go` |
| Loops | `Loops.go`, `main.go` |
| Arrays & Slices | `Arrays & Slices.go` |
| Maps | `Maps.go` |
| Functions | `Functions.go` |
| Structs | `Structs.go` |
| Methods | `Methods.go` |
| Interfaces | `Interfaces.go` |
| Error Handling | `ErrorHandling.go` |
| Goroutines | `Goroutines.go` |
| Channels | `Channels.go` |

##  Tier 2 — `intermediate/` (Real-world building blocks)

| Topic | File |
|---|---|
| Pointers | `01_Pointers.go` |
| Closures | `02_Closures.go` |
| defer / panic / recover | `03_DeferPanicRecover.go` |
| Struct embedding & composition | `04_StructEmbedding.go` |
| Strings, runes & unicode | `05_StringsAndRunes.go` |
| File I/O & JSON (encoding/json) | `06_FileAndJSON.go` |
| Error wrapping & context.Context | `07_ErrorWrappingAndContext.go` |
| Testing (`go test`, table-driven tests, benchmarks) | `08_Testing/` |

##  Tier 3 — `advanced/` (Pro-level Go)

| Topic | File |
|---|---|
| Generics (type parameters, constraints) | `01_Generics.go` |
| sync.RWMutex, sync.Once, sync.Map, atomic | `02_SyncPrimitives.go` |
| select, fan-in, worker pools | `03_SelectAndWorkerPool.go` |
| net/http server + client, middleware | `04_HTTPServerAndClient.go` |
| reflect + Functional Options pattern | `05_ReflectionAndFunctionalOptions.go` |

---

##  Where to go next (topics beyond this repo)

Once you're comfortable with everything above, these are the next things a
professional Go developer typically learns:

- **Modules & tooling**: `go mod init/tidy`, semantic versioning, private modules
- **Project layout**: the community `cmd/`, `internal/`, `pkg/` convention
- **Concurrency patterns**: pipelines, `errgroup`, rate limiting, context propagation across services
- **Databases**: `database/sql`, connection pooling, migrations, an ORM like `sqlc` or `gorm`
- **gRPC & Protocol Buffers**: for service-to-service APIs
- **Web frameworks**: `chi`, `gin`, or `echo` once you understand `net/http`
- **Dependency injection & clean architecture** in Go
- **Structured logging** (`slog`), metrics (Prometheus), tracing (OpenTelemetry)
- **Docker + deployment**: multi-stage builds for tiny Go binaries
- **CI/CD**: `golangci-lint`, `go vet`, GitHub Actions
- **Profiling & performance**: `pprof`, benchmarking, escape analysis
- **Advanced testing**: mocking (`gomock`, `testify`), fuzzing (`go test -fuzz`), integration tests
- **Building CLIs**: `cobra` / `flag` package
- **Building CLIs & tools that others `go install`**

##  Recommended free resources

- [A Tour of Go](https://go.dev/tour/) — official interactive tutorial
- [Go by Example](https://gobyexample.com/) — short, practical snippets per topic
- [Effective Go](https://go.dev/doc/effective_go) — official style/idioms guide
- [100 Go Mistakes and How to Avoid Them](https://100go.co/) — great once you're past basics
