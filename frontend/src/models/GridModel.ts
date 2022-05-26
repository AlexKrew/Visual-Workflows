import Vector2 from "@/components/util/Vector";
import NodeModel from "./NodeModel";

class GridModel {
  posAbs: Vector2;
  cellSize: number;
  nodes: NodeModel[] = [];

  constructor(cellSize = 10, posAbs = new Vector2(0, 0)) {
    this.posAbs = posAbs;
    this.cellSize = cellSize;
  }

  addNodes(...nodes: NodeModel[]) {
    nodes.forEach((node) => this.nodes.push(node));
  }

  addPos(posAbs: Vector2) {
    this.posAbs.addVector(posAbs);
    this.nodes.forEach((node) => node.updatePos());
  }
}

export default GridModel;
