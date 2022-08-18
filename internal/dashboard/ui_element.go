package dashboard

import "workflows/internal/workflows"

type UIElement struct {
	ID       string         `json:"id"`
	Type     string         `json:"type"`
	Fields   map[string]any `json:"fields"`
	Children []interface{}  `json:"children"` // either UIElements | UIElementLeaf
}

type UIElementLeaf struct {
	ID     string         `json:"id"`
	Type   string         `json:"type"`
	Fields map[string]any `json:"fields"`
}

func uiElementFromNode(node workflows.Node) UIElement {
	return UIElement{
		ID:       node.ID,
		Type:     node.Type,
		Children: []interface{}{},
		Fields:   fieldsFromNode(node),
	}
}

func uiElementLeafFromNode(node workflows.Node) UIElementLeaf {
	return UIElementLeaf{
		ID:     node.ID,
		Type:   node.Type,
		Fields: fieldsFromNode(node),
	}
}

func fieldsFromNode(node workflows.Node) map[string]any {
	fields := make(map[string]any)

	for _, port := range node.Ports {

		if !port.IsInputPort {
			continue
		}

		fields[port.Identifier] = port.DefaultMessage.Value
	}

	return fields
}
