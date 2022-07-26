import Vector2 from "@/components/util/Vector";
import GridData from "./Data/GridData";
import GridModel from "./GridModel";
import NodeModel from "./Node/NodeModel";

abstract class EditorComponent {
  readonly id: string;

  parent: EditorComponent | undefined;
  children: EditorComponent[] = [];

  posRel: Vector2 = new Vector2(0, 0); // Pos relative to Parent
  posGrid: Vector2 = new Vector2(0, 0); // Absolute Pos
  posGridCell: Vector2 = new Vector2(0, 0); // Pos relative to Grid with GridSnap
  snapToGrid = false;

  constructor(id: string, snapToGrid: boolean, children: EditorComponent[] = []) {
    this.id = id;
    this.snapToGrid = snapToGrid;
    children.forEach((child) => this.addChildren(child));
  }

  abstract clone(): EditorComponent;

  //#region Parent & Child
  setParent(parent: EditorComponent) {
    this.parent = parent;
    this.updatePos();
  }

  addChildren(...children: EditorComponent[]) {
    children.forEach((child) => {
      this.children.push(child);
      child.setParent(this);
    });

    this.addChildrenOverload(...children);
  }

  abstract addChildrenOverload(...children: EditorComponent[]): void;

  getChildIndex(id: string): number {
    let index = -1;
    this.children.forEach((child, i) => {
      if (child.id == id) {
        index = i;
      }
    });
    return index;
  }

  getChildById(id: string): EditorComponent | undefined {
    const index = this.getChildIndex(id);
    if (index >= 0) return this.children[index];
    else return undefined;
  }

  //#endregion

  //#region Position
  setPos(pos: Vector2) {
    this.posRel = pos;
    this.updatePos();
  }

  addPos(pos: Vector2) {
    this.setPos(Vector2.add(this.posRel, pos));
  }

  updatePos(): void {
    if (this.parent) this.posGrid = Vector2.add(this.posRel, this.parent?.posGrid);

    if (this.snapToGrid) this.posGridCell = EditorComponent.posToGridPos(this.posGrid, GridData.cellSize);
    else if (this.parent) this.posGridCell = Vector2.add(this.posRel, this.parent?.posGridCell);

    this.children.forEach((child) => child.updatePos());

    this.updatePosOverload();
  }

  abstract updatePosOverload(): void;

  addPosToChildren(pos: Vector2, startIndex = 0) {
    for (let i = startIndex; i < this.children.length; i++) {
      this.children[i].addPos(pos);
    }
  }

  static posToGridPos(pos: Vector2, cellSize: number) {
    const x = Math.floor(pos.x / cellSize) * cellSize;
    const y = Math.floor(pos.y / cellSize) * cellSize;
    return new Vector2(x, y);
  }
  //#endregion
}

export default EditorComponent;
