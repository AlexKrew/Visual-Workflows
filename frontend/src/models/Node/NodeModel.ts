import PortModel from "./PortModel";
import EditorComponent from "../EditorComponent";
import { uuid } from "vue-uuid";
import { emitter } from "@/components/util/Emittery";
import { NodeType, PortType } from "../Data/Types";
import Vector2 from "@/components/util/Vector";
import GridData from "../Data/GridData";
import GridModel from "../GridModel";
import { deepCopy } from "@/components/util/DeepCopy";

class NodeModel extends EditorComponent {
  data: NodeType;

  constructor(data: NodeType) {
    super(data.id, true);
    this.data = data;

    const pos = new Vector2(data.ui.position[0], data.ui.position[1]);
    this.setPos(pos);

    // Add Ports and Get Missing Fields from the global default Nodes
    const defaultNode = GridData.nodes.find((node) => node.data.type == this.data.type)?.data;
    if (defaultNode) {
      // Get Missing Fields for Ports
      this.data.ports.forEach((port, index) => {
        const defaultPort = defaultNode?.ports[index];
        if (defaultPort) {
          port.hasDefaultField = defaultPort.hasDefaultField;
          port.defaultPlaceholder = defaultPort.defaultPlaceholder;
          port.options = defaultPort.options;
        }

        this.addChild(new PortModel(port), false);
      });

      // Get Addable Ports
      this.data.addablePorts = defaultNode?.addablePorts;
    }
  }

  clone(): EditorComponent {
    // Shallow Clone the Node Data
    const newData: NodeType = deepCopy(this.data);
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

  setName(name: string) {
    this.data.name = name;
  }

  addChildToData(child: EditorComponent): void {
    throw new Error("Method not implemented.");
  }

  removeChild(id: string): void {
    (this.parent as GridModel).removeEdges(undefined, undefined, id);

    const index = this.children.findIndex((port) => (port as PortModel).data.id == id);
    this.children.splice(index, 1);
    this.data.ports.splice(index, 1);
  }

  addChildrenOverload(...children: EditorComponent[]): void {
    return;
  }

  // Adds every addable Ports once, sets their group ID and reloads every Port position
  addAddablePorts() {
    const groupID = "Group-" + uuid.v4();
    this.data.addablePorts.forEach((port) => {
      const portClone: PortType = { ...port };
      portClone.id = "Port-" + uuid.v4();
      portClone.group_id = groupID;
      portClone.added = true;

      this.data.ports.push(portClone);
      this.addChild(new PortModel(portClone), false);
    });

    emitter.emit("PortsUpdatePos", this);
  }

  removeAddablePorts(groupID: string) {
    const ports: PortModel[] = this.children.filter(
      (port) => (port as PortModel).data.group_id == groupID
    ) as PortModel[];

    ports.forEach((port) => this.removeChild(port.data.id));

    emitter.emit("PortsUpdatePos", this);
  }

  updatePosOverload(): void {
    this.data.ui.position = [this.posRel.x, this.posRel.y];
  }
}

export default NodeModel;
