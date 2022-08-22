package workflows

import (
	"errors"
)

type CronJob struct {
	NodeID    NodeID
	Interval  int // interval in seconds
	TriggerID int
}

func NewCronJob(node Node) (CronJob, error) {
	if node.Type != "CronJob" {
		return CronJob{}, errors.New("node is not of type 'CronJob'")
	}

	var intervalPort *Port = nil
	for _, port := range node.Ports {
		if port.Identifier == "interval" {
			intervalPort = &port
		}
	}

	if intervalPort == nil {
		return CronJob{}, errors.New("cronjob node is missing interval port")
	}

	intervalInSeconds := int(intervalPort.DefaultMessage.Value.(float64))
	if intervalInSeconds <= 0 {
		return CronJob{}, errors.New("cronjob interval should be > 0")
	}

	return CronJob{
		NodeID:   node.ID,
		Interval: intervalInSeconds,
	}, nil
}
