import EditorComponent from "../EditorComponent";

class NodePortModel extends EditorComponent {
  isInput: boolean;
  portSize = 15;

  constructor(id: string, title = "New Port", isInput = false) {
    super(id, title, false);
    this.isInput = isInput;
  }
}

export default NodePortModel;
