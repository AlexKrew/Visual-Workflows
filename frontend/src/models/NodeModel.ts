import Vector2 from "@/components/util/Vector";
import GridModel from "./GridModel";
import NodePortModel from "./NodePortModel";
import InteractUtil from "@/components/util/InteractUtil";

class NodeModel {
  id: string;
  title: string;
  grid: GridModel; // Parent Grid
  pos: Vector2 = new Vector2(0, 0); // Pos relative to Parent
  posAbs: Vector2 = new Vector2(0, 0); // Absolute Pos
  gridPos: Vector2 = new Vector2(0, 0); // Pos relative to Grid with GridSnap
  ports: NodePortModel[] = [];

  constructor(id: string, title = "New Node", grid: GridModel, pos = new Vector2(0, 0)) {
    this.id = id;
    this.title = title;
    this.grid = grid;
    this.changePos(pos);
  }

  changePos(pos: Vector2) {
    this.pos = pos;
    this.updatePos();
  }

  addPos(pos: Vector2) {
    this.changePos(Vector2.add(this.pos, pos));
  }

  updatePos() {
    this.posAbs = Vector2.add(this.pos, this.grid.posAbs);
    this.gridPos = InteractUtil.posToGridPos(this.pos, this.grid.cellSize);
    this.ports.forEach((port) => port.updatePos());
  }

  addPorts(...ports: NodePortModel[]) {
    ports.forEach((port) => this.ports.push(port));
  }

  getPortByID(id: string): NodePortModel | null {
    for (const port of this.ports) {
      if (port.id == id) {
        return port;
      }
    }
    return null;
  }
}

export default NodeModel;
