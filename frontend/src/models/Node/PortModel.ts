import EditorComponent from "../EditorComponent";
import { uuid } from "vue-uuid";
import { PortType } from "../Data/Types";

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

  setDefaultValue(text: string) {
    this.data.default_value.value = text;
  }

  setGroupID(id: string) {
    this.data.added = true;
    this.data.group_id = id;
  }

  updatePosOverload(): void {
    return;
  }

  addChildrenOverload(...children: EditorComponent[]): void {
    return;
  }
}

export default PortModel;
