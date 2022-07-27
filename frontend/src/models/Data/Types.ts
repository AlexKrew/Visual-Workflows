type WorkflowType = {
  id: string;
  name: string;
  nodes: NodeType[];
  edges: EdgeType[];
};

type NodeType = {
  id: string;
  name: string;
  category: string;
  type: string;
  ports: PortType[];
  addablePorts: PortType[];
  ui: {
    position: number[];
  };
};

type PortType = {
  id: string;
  group_id: string;
  label: string;
  added: boolean;
  is_input: boolean;
  default_value: {
    datatype: "STRING";
    value: string;
  };
  hasDefaultField: boolean;
  defaultPlaceholder: string;
};

type EdgeType = {
  id: string;
  origin: {
    node_id: string;
    port_id: string;
  };
  target: {
    node_id: string;
    port_id: string;
  };
};

export { WorkflowType, NodeType, PortType, EdgeType };
