import EditorComponent from "../EditorComponent";
import { uuid } from "vue-uuid";

class NodePortModel extends EditorComponent {
  placeholder: string;
  isInput: boolean;
  portSize = 15;
  hasDefaultField: boolean;

  constructor(id: string, title = "New Port", isInput = false, hasDefaultField = false, placeholder = "") {
    super(id, title, false);
    this.isInput = isInput;
    this.hasDefaultField = hasDefaultField;
    this.placeholder = placeholder;
  }

  clone(): EditorComponent {
    const port = new NodePortModel(
      "port-" + uuid.v4(),
      this.title,
      this.isInput,
      this.hasDefaultField,
      this.placeholder
    );
    if (this.parent) port.setParent(this.parent);
    return port;
  }
}

export default NodePortModel;
