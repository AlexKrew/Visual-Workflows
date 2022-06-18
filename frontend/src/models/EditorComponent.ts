import Vector2 from "@/components/util/Vector";

abstract class EditorComponent {
  readonly id: string;
  title: string;

  parent: EditorComponent | undefined;
  children: EditorComponent[] = [];

  posRel: Vector2 = new Vector2(0, 0); // Pos relative to Parent
  posAbs: Vector2 = new Vector2(0, 0); // Absolute Pos
  posGrid: Vector2 = new Vector2(0, 0); // Pos relative to Grid with GridSnap

  constructor(id: string, title = "") {
    this.id = id;
    this.title = title;
  }

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
  }

  getChildById(id: string): EditorComponent | undefined {
    for (const child of this.children) {
      if (child.id == id) {
        return child;
      }
    }
    return undefined;
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

  abstract updatePos(): void;
  //#endregion
}

export default EditorComponent;
