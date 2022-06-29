import NodePortModel from "./NodePortModel";
import EditorComponent from "../EditorComponent";
import { uuid } from "vue-uuid";

class NodeModel extends EditorComponent {
  category: string;
  addablePorts: NodePortModel[];
  addedPorts: NodePortModel[] = [];

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
    return new NodeModel(uuid.v4(), this.title, this.category, this.children as NodePortModel[], this.addablePorts);
  }

  addAddablePorts() {
    //TODO
    this.addablePorts.forEach((port) => {
      this.addedPorts.push(port.clone() as NodePortModel);
    });
  }

  removeAddablePorts() {
    //TODO
  }
}

export default NodeModel;
