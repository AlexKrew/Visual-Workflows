import { Vector2 } from "@/components/util/Vector";
import NodeModel from "./NodeModel";

class GridModel {
  pos: Vector2;
  cellSize: number;
  nodes: NodeModel[] = [];

  constructor(cellSize = 10, pos = new Vector2(0, 0), nodes: NodeModel[] = []) {
    this.pos = pos;
    this.cellSize = cellSize;
    this.nodes = nodes;
  }

  addNode(node: NodeModel) {
    this.nodes.push(node);
  }

  addNodes(nodes: NodeModel[]) {
    nodes.forEach((node) => this.addNode(node));
  }

  addPos(pos: Vector2) {
    this.pos.addVector(pos);
  }
}

export default GridModel;
