import { Vector2 } from "@/components/util/Vector";
import NodePortModel from "./NodePortModel";

class NodeModel {
  id: string;
  title: string;
  pos: Vector2;
  ports: NodePortModel[] = [];

  constructor(id: string, title = "New Node", pos = new Vector2(0, 0)) {
    this.id = id;
    this.title = title;
    this.pos = pos;
  }

  addPort(port: NodePortModel) {
    this.ports.push(port);
  }

  addPorts(ports: NodePortModel[]) {
    ports.forEach((port) => this.addPort(port));
  }

  toString() {
    return `id: ${this.id}, title: ${this.title}, pos: ${
      this.pos
    }, ports: ${this.ports.toString()}`;
  }
}

export default NodeModel;
