import PortModel from "./PortModel";
import EditorComponent from "../EditorComponent";
import { uuid } from "vue-uuid";
import { emitter } from "@/components/util/Emittery";

class NodeModel extends EditorComponent {
  category: string;
  type: string;
  addablePorts: PortModel[];
  addedPorts: string[] = [];

  constructor(
    id: string,
    title = "New Node",
    category: string,
    type: string,
    ports: PortModel[],
    addablePorts: PortModel[] = []
  ) {
    super(id, title, true, ports);
    this.category = category;
    this.type = type;
    this.addablePorts = addablePorts;
  }

  clone(): EditorComponent {
    const node = new NodeModel("node-" + uuid.v4(), this.label, this.category, this.type, [], this.addablePorts);

    this.children.forEach((child) => {
      node.addChildren(child.clone());
    });

    node.addedPorts = this.addedPorts;
    if (this.parent) node.setParent(this.parent);
    return node;
  }

  // Adds every addable Ports once, sets their group ID and reloads every Port position
  addAddablePorts() {
    const groupID = uuid.v4();
    this.addablePorts.forEach((port) => {
      const portClone = port.clone() as PortModel;
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

  //#region Serialization
  static fromJSON(json: JSON): NodeModel {
    const portsJson = JSON.parse(JSON.stringify(json["ports" as keyof JSON]));
    const ports: PortModel[] = [];
    portsJson.forEach((port: JSON) => {
      ports.push(PortModel.fromJSON(port));
    });

    return new NodeModel(
      json["id" as keyof JSON] as string,
      json["name" as keyof JSON] as string,
      "Imported",
      "Type",
      ports
    );
  }
  toJSON(): JSON {
    const json = JSON.parse(JSON.stringify({}));

    json["id"] = this.id;
    json["name"] = this.label;
    json["type"] = this.type;
    json["ui"] = {};
    json["ui"]["position"] = [this.posGrid.x, this.posGrid.y];
    json["ports"] = [];

    this.children.forEach((child) => {
      json["ports"].push((child as PortModel).toJSON());
    });

    return json;
  }
  //#endregion
}

export default NodeModel;
