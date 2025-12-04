# Temporal Schedules Demo

Demonstrates how to create Temporal schedules that start immediately from creation time, rather than using clock boundaries.

## Problem

Temporal interval schedules align to the Unix epoch. A 10-minute schedule created at 12:02 runs at 12:10, 12:20, etc.

## Solution

Use the `Offset` field in `ScheduleIntervalSpec`:

```go
phase := time.Now().UTC().Sub(time.Now().UTC().Truncate(interval))
```

See: [scheduler/main.go](./scheduler/main.go)

## Usage

**1. Run Temporal:**

```bash
docker-compose up -d
```

**2. Start the worker:**
```bash
go run ./worker
```

**3. Create the schedule:**
```bash
go run ./scheduler
```

**4. View in Temporal UI:** http://localhost:5580/namespaces/default/schedules
