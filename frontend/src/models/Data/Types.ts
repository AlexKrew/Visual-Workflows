type WorkflowType = {
  id: string;
  name: string;
  nodes: Node[];
};

type NodeType = {
  id: string;
  name: string;
  category: string;
  type: string;
  ports: PortType[];
  addablePorts: PortType[];
};

type PortType = {
  id: string;
  groupID: string;
  name: string;
  added: boolean;
  is_input: boolean;
  hasDefaultField: boolean;
  defaultValue: string;
  defaultPlaceholder: string;
};

export { WorkflowType, NodeType, PortType };
