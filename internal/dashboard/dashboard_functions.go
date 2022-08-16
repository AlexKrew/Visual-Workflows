package dashboard

import (
	"workflows/internal/workflows"
)

func ConfigFromWorkflow(workflow workflows.Workflow) (UIElement, bool) {

	canvasNode, exists := findCanvasNode(workflow.Nodes)
	if !exists {
		return UIElement{}, false
	}

	uiElements := toUI(canvasNode, workflow).(UIElement)
	return uiElements, true

}

func findCanvasNode(nodes []workflows.Node) (workflows.Node, bool) {
	for _, node := range nodes {
		if node.Type == "Canvas" {
			return node, true
		}
	}

	return workflows.Node{}, false
}

func toUI(node workflows.Node, workflow workflows.Workflow) any {

	childrenPortID, hasChildren := childrenPortID(node)

	if hasChildren {

		connectedUIElements := []interface{}{}

		for _, edge := range workflow.Edges {
			if edge.Origin.NodeID == node.ID && edge.Origin.PortID == childrenPortID {

				// recursivly generate ui config
				targetNode, _ := workflow.NodeByID(edge.Target.NodeID)
				connectedUIElements = append(connectedUIElements, toUI(targetNode, workflow))
			}
		}

		ret := uiElementFromNode(node)
		ret.Children = append(ret.Children, connectedUIElements...)
		return ret

	} else {
		return uiElementLeafFromNode(node)
	}
}

func childrenPortID(node workflows.Node) (string, bool) {
	for _, port := range node.Ports {
		if port.Identifier == "children" {
			return port.ID, true
		}
	}

	return "", false
}
