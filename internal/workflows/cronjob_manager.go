package workflows

import "time"

type CronJobManager struct {
	eventStream *EventStream
	workflow    *Workflow
	cronJobs    []CronJob
	triggerID   int
}

func NewCronJobManager(eventStream *EventStream, workflow *Workflow) *CronJobManager {
	return &CronJobManager{
		eventStream: eventStream,
		workflow:    workflow,
		triggerID:   0,
	}
}

func (manager *CronJobManager) addCronJob(cronJob CronJob) {
	manager.cronJobs = append(manager.cronJobs, cronJob)
}

func (manager *CronJobManager) startCronJobs() {
	manager.stopCronJobs()

	for _, cronjob := range manager.cronJobs {
		go manager.executeCronJob(cronjob)
	}
}

func (manager *CronJobManager) executeCronJob(job CronJob) {

	// stop
	if job.TriggerID != manager.triggerID {
		return
	}

	triggerEvent := NewCronTriggerEvent(CronTriggerBody{
		WorkflowID: manager.workflow.ID,
	})
	manager.eventStream.AddEvent(triggerEvent)

	// re-call itself
	time.Sleep(time.Second * time.Duration(job.Interval))
	manager.executeCronJob(job)
}

func (manager *CronJobManager) stopCronJobs() {
	manager.triggerID++
}
