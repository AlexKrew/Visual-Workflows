import Vector2 from "@/components/util/Vector";
import NodeModel from "./NodeModel";

class NodePortModel {
  id: string;
  title: string;
  isInput: boolean;

  node: NodeModel | undefined; // Parent Node

  pos: Vector2 = new Vector2(0, 0); // Pos relative to Parent
  posAbs: Vector2 = new Vector2(0, 0); // Absolute Pos
  gridPos: Vector2 = new Vector2(0, 0); // Pos relative to Grid with GridSnap

  constructor(id: string, title = "New Port", isInput = false) {
    this.id = id;
    this.title = title;
    this.isInput = isInput;
  }

  //#region Position
  setPos(pos: Vector2) {
    this.pos = pos;
    this.updatePos();
  }

  updatePos() {
    if (this.node?.grid) {
      this.gridPos = Vector2.add(this.pos, this.node.gridPos);
      this.posAbs = Vector2.add(this.gridPos, this.node.grid.posAbs);
    }
  }
  //#endregion
}

export default NodePortModel;
