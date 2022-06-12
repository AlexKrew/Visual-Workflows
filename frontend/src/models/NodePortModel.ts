import Vector2 from "@/components/util/Vector";
import NodeConnectionModel from "./NodeConnectionModel";
import NodeModel from "./NodeModel";

class NodePortModel {
  id: string;
  title: string;
  isInput: boolean;
  node: NodeModel; // Parent Node
  pos: Vector2 = new Vector2(0, 0); // Pos relative to Parent
  posAbs: Vector2 = new Vector2(0, 0); // Absolute Pos
  gridPos: Vector2 = new Vector2(0, 0); // Pos relative to Grid with GridSnap
  connections: NodeConnectionModel[] = [];

  constructor(id: string, title = "New Port", node: NodeModel, isInput = false) {
    this.id = id;
    this.title = title;
    this.node = node;
    this.isInput = isInput;
  }

  setPos(pos: Vector2) {
    this.pos = pos;
    this.updatePos();
  }

  updatePos() {
    this.gridPos = Vector2.add(this.pos, this.node.gridPos);
    this.posAbs = Vector2.add(this.gridPos, this.node.grid.posAbs);
  }
}

export default NodePortModel;
