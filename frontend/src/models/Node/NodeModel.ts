import PortModel from "./PortModel";
import EditorComponent from "../EditorComponent";
import { uuid } from "vue-uuid";
import { emitter } from "@/components/util/Emittery";
import { NodeType } from "../Data/Types";

class NodeModel extends EditorComponent {
  data: NodeType;
  addedPorts: string[] = []; // Can be removed

  constructor(data: NodeType) {
    super(data.id, true);
    this.data = data;

    this.data.ports.forEach((port) => {
      this.addChildren(new PortModel(port));
    });
  }

  clone(): EditorComponent {
    const newData: NodeType = { ...this.data };
    newData.id = uuid.v4();

    const node = new NodeModel(newData);
    node.addedPorts = this.addedPorts;
    if (this.parent) node.setParent(this.parent);

    return node;
  }

  // Adds every addable Ports once, sets their group ID and reloads every Port position
  addAddablePorts() {
    const groupID = uuid.v4();
    this.data.addablePorts.forEach((port) => {
      const portClone = new PortModel({ ...port });
      portClone.setParent(this);
      portClone.setGroupID(groupID);

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
