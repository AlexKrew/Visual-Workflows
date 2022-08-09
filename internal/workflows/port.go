package workflows

import (
	"errors"
	"fmt"
	"workflows/internal/utils"
)

type PortID = utils.UUID
type UniquePortID = string

type Ports = []Port

type Port struct {
	ID             PortID  `json:"id"`
	Identifier     string  `json:"identifier"`
	Label          string  `json:"label"`
	Datatype       string  `json:"datatype"`
	IsInputPort    bool    `json:"is_input"`
	IsTrigger      bool    `json:"is_trigger"`
	DefaultMessage Message `json:"default_value"`
}

func PortByID(id PortID, ports []Port) (Port, bool) {
	for _, port := range ports {
		if port.ID == id {
			return port, true
		}
	}

	return Port{}, false
}

func InputPortByID(id PortID, ports []Port) (Port, bool) {
	port, exists := PortByID(id, ports)

	if !exists || !port.IsInputPort {
		return Port{}, false
	}

	return port, true
}

func OutputPortByID(id PortID, ports []Port) (Port, bool) {
	port, exists := PortByID(id, ports)

	if !exists || port.IsInputPort {
		return Port{}, false
	}

	return port, true
}

func (port *Port) IsCompatibleWith(other Port) (bool, []error) {

	errs := []error{}

	// only one input and one output port
	isInputOutput := (port.IsInputPort && !other.IsInputPort) || (!port.IsInputPort && other.IsInputPort)
	if !isInputOutput {
		errs = append(errs, errors.New("both ports are of the same type (input or output)"))
	}

	sameDatatype := port.Datatype == other.Datatype
	if !sameDatatype {
		errs = append(errs, errors.New("different datatypes"))
	}

	compatible := isInputOutput && sameDatatype

	return compatible, errs
}

type PortAddress struct {
	NodeID NodeID `json:"node_id"`
	PortID PortID `json:"port_id"`
}

func (pa *PortAddress) UniquePortID() UniquePortID {
	return fmt.Sprintf("%s--%s", pa.NodeID, pa.PortID)
}
