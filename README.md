# SchedSim: distributed CPU scheduling simulator (Go)

| Weeks | Status | Where |
|---|---|---|
| 1-2 | working | `internal/scheduler`, `cmd/schedsim-cli`, `profiling/` |
| 3-4 | working (v1) | `internal/lsm` (WAL, memtable, SSTables, bloom, compaction) |
| 5-6 | working (v1) | `internal/concurrent` (PQ, ShardedMap) + benchmarks |
| 7-8 | scaffold | `api/proto`, `internal/server` |
| 9-10 | scaffold | `internal/raft` |
| 11-12 | scaffold | `deploy/`, `chaos/` |

## Run (VS Code: open this folder, install the Go extension)
    go run ./cmd/schedsim-cli -algo all
    go run ./cmd/schedsim-cli -algo rr -quantum 3 -file workloads/sample.csv -store data
    go test -race ./...
    go test -bench=. -benchmem ./internal/...
    go run ./cmd/schedsim-server   # then: curl localhost:8080/healthz

Press F5 in VS Code and pick a launch configuration.

## Profiling baseline (Linux)
    perf stat -e context-switches,cpu-migrations,task-clock <cmd>
    pidstat -w 1      # voluntary / involuntary context switches
    perf record -g <cmd> && perf report
Save results to `profiling/baseline.md`.
