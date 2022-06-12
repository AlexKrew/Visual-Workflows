import Vector2 from "./Vector";

class BezierCurve {
  posAnchor1: Vector2;
  posAnchor2: Vector2;
  posControl1: Vector2;
  posControl2: Vector2;

  constructor(posAnchor1?: Vector2, posControl1?: Vector2, posControl2?: Vector2, posAnchor2?: Vector2) {
    this.posAnchor1 = posAnchor1 ? posAnchor1 : new Vector2(0, 0);
    this.posAnchor2 = posAnchor2 ? posAnchor2 : new Vector2(0, 0);
    this.posControl1 = posControl1 ? posControl1 : new Vector2(0, 0);
    this.posControl2 = posControl2 ? posControl2 : new Vector2(0, 0);
  }

  toSVGString() {
    // return "M10,10 C20,0 90,80 70,100"
    return `M ${this.posAnchor1.toString(0, ",")} C ${this.posControl1.toString(0, ",")} ${this.posControl2.toString(
      0,
      ","
    )} ${this.posAnchor2.toString(0, ",")}`;
  }
}

export default BezierCurve;
