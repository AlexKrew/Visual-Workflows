import Vector2 from "@/components/util/Vector";
import GridModel from "./GridModel";
import NodePortModel from "./NodePortModel";
import InteractUtil from "@/components/util/InteractUtil";

class NodeModel {
  id: string;
  title: string;
  grid: GridModel;
  pos: Vector2 = new Vector2(0, 0);
  posAbs: Vector2 = new Vector2(0, 0);
  gridPos: Vector2 = new Vector2(0, 0);
  ports: NodePortModel[] = [];

  constructor(id: string, title = "New Node", grid: GridModel, pos = new Vector2(0, 0)) {
    this.id = id;
    this.title = title;
    this.grid = grid;
    this.changePos(pos);
  }

  addPort(port: NodePortModel) {
    this.ports.push(port);
  }

  addPorts(ports: NodePortModel[]) {
    ports.forEach((port) => this.addPort(port));
  }

  changePos(pos: Vector2) {
    this.pos = pos;
    this.posAbs = Vector2.add(this.pos, this.grid.posAbs);
    this.updatePos();
  }

  addPos(pos: Vector2) {
    this.changePos(Vector2.add(this.pos, pos));
  }

  updatePos() {
    this.posAbs = Vector2.add(this.pos, this.grid.posAbs);
    this.gridPos = InteractUtil.posToGridPos(this.pos, this.grid.cellSize);
    this.ports.forEach((port) => port.updatePosAbs());
  }

  toString() {
    return `id: ${this.id}, title: ${this.title}, pos: ${this.pos}, ports: ${this.ports.toString()}`;
  }
}

export default NodeModel;
