import NodePortModel from "./NodePortModel";
import EditorComponent from "../EditorComponent";
import { uuid } from "vue-uuid";
import { emitter } from "@/components/util/Emittery";

class NodeModel extends EditorComponent {
  category: string;
  addablePorts: NodePortModel[];
  addedPorts: string[] = [];

  constructor(
    id: string,
    title = "New Node",
    category: string,
    ports: NodePortModel[],
    addablePorts: NodePortModel[] = []
  ) {
    super(id, title, true, ports);
    this.category = category;
    this.addablePorts = addablePorts;
  }

  clone(): EditorComponent {
    const node = new NodeModel(
      uuid.v4(),
      this.title,
      this.category,
      this.children as NodePortModel[],
      this.addablePorts
    );
    node.addedPorts = this.addedPorts;
    if (this.parent) node.setParent(this.parent);
    return node;
  }

  addAddablePorts() {
    //TODO
    this.addablePorts.forEach((port) => {
      const portClone = port.clone() as NodePortModel;
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
