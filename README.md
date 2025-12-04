# Temporal Schedules Demo

Demonstrates how to create Temporal schedules that start immediately from creation time, rather than using clock boundaries.

## Problem

Temporal interval schedules align to the Unix epoch. A 10-minute schedule created at 12:02 runs at 12:10, 12:20, etc.

## Solution

Use `TriggerImmediately` and `Offset`:

```go
phase := time.Now().UTC().Sub(time.Now().UTC().Truncate(interval))

c.ScheduleClient().Create(ctx, client.ScheduleOptions{
    ID:                 scheduleID,
    TriggerImmediately: true,   // Execute immediately on creation
    Spec: client.ScheduleSpec{
        Intervals: []client.ScheduleIntervalSpec{{
            Every:  interval,
            Offset: phase,      // Align subsequent runs to creation time
        }},
    },
    ...
})
```

- **`TriggerImmediately`**: Runs the workflow once when the schedule is created
- **`Offset`**: Shifts all subsequent runs to align with your creation time

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
