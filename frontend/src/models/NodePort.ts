class NodePort{
  name: String;
  node: Node;
  connectedTo: NodePort[];
  portType: PortType;
  portPos: Vector2;
}

enum PortType{
  input,
  output,
}

export {NodePort, PortType}