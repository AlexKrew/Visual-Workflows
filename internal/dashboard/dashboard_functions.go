package dashboard

import "workflows/internal/workflows"

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
		if node.Type == "UICanvas" {
			return node, true
		}
	}

	return workflows.Node{}, false
}

func toUI(node workflows.Node, workflow workflows.Workflow) any {

	childrenPortIDs, hasChildren := childrenPortIDs(node)

	if hasChildren {

		connectedUIElements := []interface{}{}

		for _, edge := range workflow.Edges {
			if edge.Origin.NodeID == node.ID && contains(childrenPortIDs, edge.Origin.PortID) {

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

func childrenPortIDs(node workflows.Node) ([]string, bool) {
	childrenPortIds := []string{}
	hasChildren := false

	for _, port := range node.Ports {
		if port.Identifier == "children" {
			childrenPortIds = append(childrenPortIds, port.ID)
			hasChildren = true
		}
	}

	return childrenPortIds, hasChildren
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
