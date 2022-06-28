import Vector2 from "@/components/util/Vector";
import GridModel from "../GridModel";
import NodePortModel from "./NodePortModel";
import InteractUtil from "@/components/util/InteractUtil";
import EditorComponent from "../EditorComponent";

class NodeModel extends EditorComponent {
  category: string;

  constructor(id: string, title = "New Node", category: string, ...ports: NodePortModel[]) {
    super(id, title);
    this.category = category;
    ports.forEach((port) => this.addChildren(port));
  }

  updatePos() {
    if (this.parent) {
      this.posAbs = Vector2.add(this.posRel, this.parent.posAbs);
      this.posGrid = InteractUtil.posToGridPos(this.posRel, (this.parent as GridModel).cellSize);
      this.children.forEach((child) => child.updatePos());
    }
  }
}

export default NodeModel;