import Vector2 from "@/components/util/Vector";
import GridModel from "../GridModel";
import NodePortModel from "./NodePortModel";
import InteractUtil from "@/components/util/InteractUtil";
import EditorComponent from "../EditorComponent";

class NodeModel extends EditorComponent {
  category: string;

  constructor(id: string, title = "New Node", category: string, ...ports: NodePortModel[]) {
    super(id, title, true);
    this.category = category;
    ports.forEach((port) => this.addChildren(port));
  }
}

export default NodeModel;
