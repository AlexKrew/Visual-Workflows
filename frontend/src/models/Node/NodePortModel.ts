import Vector2 from "@/components/util/Vector";
import EditorComponent from "../EditorComponent";

class NodePortModel extends EditorComponent {
  isInput: boolean;

  constructor(id: string, title = "New Port", isInput = false) {
    super(id, title);
    this.isInput = isInput;
  }

  updatePos() {
    if (this.parent?.parent) {
      this.posGrid = Vector2.add(this.posRel, this.parent.posGrid);
      this.posAbs = Vector2.add(this.posGrid, this.parent.parent.posAbs);
    }
  }
}

export default NodePortModel;
