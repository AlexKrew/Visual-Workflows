import EditorComponent from "../EditorComponent";
import { uuid } from "vue-uuid";
import { PortType } from "../Data/Types";
import { emitter } from "@/components/util/Emittery";
import NodeModel from "./NodeModel";
import { Datatype } from "../Data/DataTypes";

class PortModel extends EditorComponent {
  data: PortType;
  portSize = 15;

  constructor(data: PortType) {
    super(data.id, false);
    this.data = data;
  }

  clone(): EditorComponent {
    const newData: PortType = { ...this.data };
    newData.id = "Port-" + uuid.v4();

    const port = new PortModel(newData);
    if (this.parent) port.setParent(this.parent);

    return port;
  }

  addChildToData(child: EditorComponent): void {
    throw new Error("Method not implemented.");
  }

  removeChild(id: string): void {
    throw new Error("Method not implemented.");
  }

  setDefaultValue(value: string | boolean | number, datatype: Datatype) {
    this.data.default_value.datatype = datatype;

    if (datatype == Datatype.STRING) this.data.default_value.value = value.toString();
    else if (datatype == Datatype.NUMBER) this.data.default_value.value = +value;
    else if (datatype == Datatype.BOOLEAN) this.data.default_value.value = value as boolean;
  }

  setGroupID(id: string) {
    this.data.added = true;
    this.data.group_id = id;
  }

  setDefaultFieldHidden(hidden: boolean) {
    this.data.defaultFieldHidden = hidden;
    emitter.emit("PortsUpdatePos", this.parent as NodeModel);
  }

  updatePosOverload(): void {
    return;
  }

  addChildrenOverload(...children: EditorComponent[]): void {
    return;
  }
}

export default PortModel;
