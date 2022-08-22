package workflows

import (
	"log"
	"time"
)

type CronJobManager struct {
	eventStream     *EventStream
	workflow        *Workflow
	cronJobs        []CronJob
	runningCronJobs []CronJob
	triggerID       int
}

func NewCronJobManager(eventStream *EventStream, workflow *Workflow) *CronJobManager {
	manager := CronJobManager{
		eventStream:     eventStream,
		workflow:        workflow,
		triggerID:       0,
		cronJobs:        []CronJob{},
		runningCronJobs: []CronJob{},
	}

	return &manager
}

func (manager *CronJobManager) addCronJob(cronJob CronJob) {

	// check if cronjob for node already registered
	for _, job := range manager.cronJobs {
		if job.NodeID == cronJob.NodeID {
			return
		}
	}

	manager.cronJobs = append(manager.cronJobs, cronJob)
}

func (manager *CronJobManager) startCronJobs() {
	manager.stopCronJobs()

	for _, cronjob := range manager.cronJobs {
		log.Println("EXEC CRONS", cronjob.NodeID)
		cronjobCopy := cronjob
		cronjobCopy.TriggerID = manager.triggerID
		log.Println("COPY", cronjob.TriggerID, cronjobCopy.TriggerID)
		manager.runningCronJobs = append(manager.runningCronJobs, cronjobCopy)
		go manager.executeCronJob(cronjobCopy)
	}
}

func (manager *CronJobManager) executeCronJob(job CronJob) {

	// stop
	if job.TriggerID != manager.triggerID {
		removed := manager.removeRunningCronJob(job)
		log.Println("STOP, Remove", removed, "Running:", len(manager.runningCronJobs))
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

func (manager *CronJobManager) removeRunningCronJob(cronJob CronJob) bool {
	for i, job := range manager.runningCronJobs {
		if job.NodeID == cronJob.NodeID && job.TriggerID == cronJob.TriggerID {
			// remove index
			manager.runningCronJobs = append(manager.runningCronJobs[:i], manager.runningCronJobs[i+1:]...)
			return true
		}
	}

	return false
}

func (manager *CronJobManager) stopCronJobs() {
	manager.triggerID++
}
