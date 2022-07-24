import PortModel from "./PortModel";
import EditorComponent from "../EditorComponent";
import { uuid } from "vue-uuid";
import { emitter } from "@/components/util/Emittery";

class NodeModel extends EditorComponent {
  category: string;
  addablePorts: PortModel[];
  addedPorts: string[] = [];

  constructor(id: string, title = "New Node", category: string, ports: PortModel[], addablePorts: PortModel[] = []) {
    super(id, title, true, ports);
    this.category = category;
    this.addablePorts = addablePorts;
  }

  clone(): EditorComponent {
    const node = new NodeModel("node-" + uuid.v4(), this.title, this.category, [], this.addablePorts);

    this.children.forEach((child) => {
      node.addChildren(child.clone());
    });

    node.addedPorts = this.addedPorts;
    if (this.parent) node.setParent(this.parent);
    return node;
  }

  addAddablePorts() {
    //TODO
    this.addablePorts.forEach((port) => {
      const portClone = port.clone() as PortModel;
      portClone.setParent(this);
      this.addedPorts.push(portClone.id);
      this.addChildren(portClone);
      emitter.emit("PortsUpdatePos", this);
    });
  }

  removeAddablePorts() {
    //TODO
  }
}

export default NodeModel;
