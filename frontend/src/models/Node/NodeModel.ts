import PortModel from "./PortModel";
import EditorComponent from "../EditorComponent";
import { uuid } from "vue-uuid";
import { emitter } from "@/components/util/Emittery";
import ISerializable from "../ISerializable";

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

  //#region Serialization
  fromJSON(json: JSON): ISerializable {
    throw new Error("Method not implemented.");
  }
  toJSON(): JSON {
    const json = JSON.parse(JSON.stringify({}));

    json["id"] = this.id;
    json["name"] = this.title;
    json["type"] = "Debug"; // TODO
    // json["ui"]["position"] = [this.posGrid.x, this.posGrid.y];
    json["ports"] = [];

    this.children.forEach((child) => {
      json["ports"].push(child.toJSON());
    });

    return json;
  }
  //#endregion
}

export default NodeModel;
