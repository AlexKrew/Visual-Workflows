import Vector2 from "@/components/util/Vector";
import EditorComponent from "../EditorComponent";

class NodePortModel extends EditorComponent {
  isInput: boolean;

  constructor(id: string, title = "New Port", isInput = false) {
    super(id, title, false);
    this.isInput = isInput;
  }
}

export default NodePortModel;
