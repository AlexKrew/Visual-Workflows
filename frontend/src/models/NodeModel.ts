import Vector2 from "@/components/util/Vector";
import GridModel from "./GridModel";
import NodePortModel from "./NodePortModel";
import InteractUtil from "@/components/util/InteractUtil";

class NodeModel {
  id: string;
  title: string;

  grid: GridModel | undefined; // Parent Grid
  ports: NodePortModel[] = [];

  pos: Vector2 = new Vector2(0, 0); // Pos relative to Parent
  posAbs: Vector2 = new Vector2(0, 0); // Absolute Pos
  gridPos: Vector2 = new Vector2(0, 0); // Pos relative to Grid with GridSnap

  constructor(id: string, title = "New Node", pos = new Vector2(0, 0), ...ports: NodePortModel[]) {
    this.id = id;
    this.title = title;
    this.setPos(pos);
    ports.forEach((port) => this.addPort(port));
  }

  setGrid(grid: GridModel) {
    this.grid = grid;
    this.updatePos();
  }

  //#region Position
  setPos(pos: Vector2) {
    this.pos = pos;
    this.updatePos();
  }

  addPos(pos: Vector2) {
    this.setPos(Vector2.add(this.pos, pos));
  }

  updatePos() {
    if (this.grid) {
      this.posAbs = Vector2.add(this.pos, this.grid.posAbs);
      this.gridPos = InteractUtil.posToGridPos(this.pos, this.grid.cellSize);
      this.ports.forEach((port) => port.updatePos());
    }
  }
  //#endregion

  //#region Ports
  addPort(port: NodePortModel) {
    this.ports.push(port);
    port.node = this;
  }

  getPortByID(id: string): NodePortModel | null {
    for (const port of this.ports) {
      if (port.id == id) {
        return port;
      }
    }
    return null;
  }
  //#endregion
}

export default NodeModel;
