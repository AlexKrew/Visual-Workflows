import { Vector2 } from "@/components/util/Vector";
import NodeModel from "./NodeModel";

class NodePortModel {
  id: string;
  title: string;
  pos: Vector2;
  node: NodeModel;
  connectedTo: NodePortModel | null;
  isInput: boolean;

  constructor(
    id: string,
    title = "New Port",
    pos = new Vector2(0, 0),
    node: NodeModel,
    connectedTo = null,
    isInput = false
  ) {
    this.id = id;
    this.title = title;
    this.pos = pos;
    this.node = node;
    this.connectedTo = connectedTo;
    this.isInput = isInput;
  }

  connectTo(port: NodePortModel) {
    this.connectedTo = port;
  }

  toString() {
    return `id: ${this.id}, title: ${this.title}, node: ${this.node.title}, connectedTo: ${this.connectedTo?.title}`;
  }
}

export default NodePortModel;
