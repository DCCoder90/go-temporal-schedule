package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"schedules-demo/scheduled"
)

func main() {
	c, err := client.Dial(client.Options{
		HostPort: "localhost:5533",
	})
	if err != nil {
		log.Fatalln("Unable to create Temporal client:", err)
	}
	defer c.Close()

	w := worker.New(c, scheduled.TaskQueue, worker.Options{})
	w.RegisterWorkflow(scheduled.ScheduledWorkflow)

	log.Println("Starting worker on task queue:", scheduled.TaskQueue)
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker:", err)
	}
}
