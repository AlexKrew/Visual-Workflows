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
    newData.id = uuid.v4();

    const port = new PortModel(newData);
    if (this.parent) port.setParent(this.parent);

    return port;
  }

  setDefaultValue(text: string) {
    this.data.defaultValue = text;
  }

  setGroupID(id: string) {
    this.data.added = true;
    this.data.groupID = id;
  }
}

export default PortModel;
