import { Vector2 } from "@/components/util/Vector";
import NodeModel from "./NodeModel";

class GridModel {
  pos: Vector2;
  nodes: NodeModel[] = [];

  constructor(pos = new Vector2(0, 0), nodes: NodeModel[] = []) {
    this.pos = pos;
    this.nodes = nodes;
  }

  addNode(node: NodeModel) {
    this.nodes.push(node);
  }

  addNodes(nodes: NodeModel[]) {
    nodes.forEach((node) => this.addNode(node));
  }

  addPos(posChange: Vector2) {
    this.pos.addVector(posChange);
  }
}

export default GridModel;
