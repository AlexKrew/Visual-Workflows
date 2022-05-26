import Vector2 from "@/components/util/Vector";
import GridModel from "./GridModel";
import NodePortModel from "./NodePortModel";
import InteractUtil from "@/components/util/InteractUtil";

class NodeModel {
  id: string;
  title: string;
  grid: GridModel;
  pos: Vector2;
  gridPos: Vector2;
  ports: NodePortModel[] = [];

  constructor(id: string, title = "New Node", grid: GridModel, pos = new Vector2(0, 0)) {
    this.id = id;
    this.title = title;
    this.grid = grid;
    this.pos = pos;
    this.gridPos = pos;
  }

  addPort(port: NodePortModel) {
    this.ports.push(port);
  }

  addPorts(ports: NodePortModel[]) {
    ports.forEach((port) => this.addPort(port));
  }

  changePos(pos: Vector2) {
    this.pos = pos;
    this.updateGridPos();
  }

  addPos(pos: Vector2) {
    this.changePos(Vector2.addVector(this.pos, pos));
  }

  updateGridPos() {
    this.gridPos = InteractUtil.posToGridPos(this.pos, this.grid.cellSize);
  }

  toString() {
    return `id: ${this.id}, title: ${this.title}, pos: ${this.pos}, ports: ${this.ports.toString()}`;
  }
}

export default NodeModel;
