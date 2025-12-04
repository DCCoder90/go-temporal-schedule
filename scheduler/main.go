package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"schedules-demo/scheduled"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{
		HostPort: "localhost:5533",
	})
	if err != nil {
		log.Fatalln("Unable to create Temporal client:", err)
	}
	defer c.Close()

	interval := 1 * time.Minute // Change this number to your interval (e.g., 10 * time.Minute)
	scheduleID := "immediate-start-schedule-2"

	// To start at the current time, we need: phase = currentTime % interval
	now := time.Now().UTC()
	phase := now.Sub(now.Truncate(interval))

	fmt.Printf("Current time: %s\n", now.Format(time.RFC3339))
	fmt.Printf("Interval: %s\n", interval)
	fmt.Printf("Phase offset: %s\n", phase)
	fmt.Printf("Next runs will be at: %s, %s, %s, ...\n",
		now.Format("15:04:05"),
		now.Add(interval).Format("15:04:05"),
		now.Add(2*interval).Format("15:04:05"),
	)

	// Create the schedule with the calculated phase offset
	_, err = c.ScheduleClient().Create(context.Background(), client.ScheduleOptions{
		ID:                 scheduleID,
		TriggerImmediately: true, // Run once on creation
		Spec: client.ScheduleSpec{
			Intervals: []client.ScheduleIntervalSpec{
				{
					Every:  interval,
					Offset: phase, // Subsequent runs align to creation time
				},
			},
		},
		Action: &client.ScheduleWorkflowAction{
			ID:        "scheduled-workflow",
			Workflow:  scheduled.ScheduledWorkflow,
			TaskQueue: scheduled.TaskQueue,
		},
	})

	if err != nil {
		log.Fatalln("Unable to create schedule:", err)
	}
}
