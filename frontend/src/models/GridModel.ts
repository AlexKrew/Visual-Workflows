import Vector2 from "@/components/util/Vector";
import NodeModel from "./NodeModel";
import NodePortModel from "./NodePortModel";

class GridModel {
  posAbs: Vector2;
  cellSize: number;
  nodes: NodeModel[] = [];

  constructor(cellSize = 10, posAbs = new Vector2(0, 0)) {
    this.posAbs = posAbs;
    this.cellSize = cellSize;
  }

  addPos(posAbs: Vector2) {
    this.posAbs.addVector(posAbs);
    this.nodes.forEach((node) => node.updatePos());
  }

  addNodes(...nodes: NodeModel[]) {
    nodes.forEach((node) => this.nodes.push(node));
  }

  getPortByID(id: string): NodePortModel | null {
    for (const node of this.nodes) {
      const port = node.getPortByID(id);
      if (port) {
        return port;
      }
    }

    return null;
  }
}

export default GridModel;
