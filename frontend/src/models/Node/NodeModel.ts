import PortModel from "./PortModel";
import EditorComponent from "../EditorComponent";
import { uuid } from "vue-uuid";
import { emitter } from "@/components/util/Emittery";
import { NodeType, PortType } from "../Data/Types";
import Vector2 from "@/components/util/Vector";
import GridData from "../Data/GridData";

class NodeModel extends EditorComponent {
  data: NodeType;
  addedPorts: string[] = []; // Can be removed

  constructor(data: NodeType) {
    super(data.id, true);
    this.data = data;

    const pos = new Vector2(data.ui.position[0], data.ui.position[1]);
    this.setPos(pos);

    // Add Ports and Get Missing Fields from the global default Nodes
    const defaultNode = GridData.nodes.find((node) => node.data.type == this.data.type)?.data;
    this.data.ports.forEach((port, index) => {
      const defaultPort = defaultNode?.ports[index];
      if (defaultPort) {
        port.hasDefaultField = defaultPort.hasDefaultField;
        port.defaultPlaceholder = defaultPort.defaultPlaceholder;
      }

      this.addChild(new PortModel(port), false);
    });

    this.data.addablePorts = [];
  }

  clone(): EditorComponent {
    // Shallow Clone the Node Data
    const newData: NodeType = { ...this.data };
    newData.id = "Node-" + uuid.v4();

    // Shallow Clone ports
    const newPorts: PortType[] = [];
    newData.ports.forEach((port) => {
      const newPort: PortType = { ...port };
      newPort.id = "Port-" + uuid.v4();
      newPorts.push(newPort);
    });
    newData.ports = newPorts;

    const node = new NodeModel(newData);
    if (this.parent) node.setParent(this.parent);

    return node;
  }

  addChildToData(child: EditorComponent): void {
    throw new Error("Method not implemented.");
  }

  removeChild(id: string): void {
    throw new Error("Method not implemented.");
  }

  addChildrenOverload(...children: EditorComponent[]): void {
    return;
  }

  // Adds every addable Ports once, sets their group ID and reloads every Port position
  addAddablePorts() {
    const groupID = uuid.v4();
    this.data.addablePorts.forEach((port) => {
      const portClone = new PortModel({ ...port });
      portClone.setParent(this);
      portClone.setGroupID(groupID);

      this.addedPorts.push(portClone.id);
      this.addChild(portClone, false);
      emitter.emit("PortsUpdatePos", this);
    });
  }

  removeAddablePorts() {
    //TODO
  }

  updatePosOverload(): void {
    this.data.ui.position = [this.posRel.x, this.posRel.y];
  }
}

export default NodeModel;
