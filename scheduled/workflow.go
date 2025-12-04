package scheduled

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

const TaskQueue = "schedules-demo-task-queue"

func ScheduledWorkflow(ctx workflow.Context) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("Scheduled workflow executed", "time", workflow.Now(ctx).Format(time.RFC3339))
	return nil
}
