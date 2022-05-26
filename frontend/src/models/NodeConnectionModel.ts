import BezierCurve from "@/components/util/BezierCurve";
import Vector2 from "@/components/util/Vector";
import NodePortModel from "./NodePortModel";

class NodeConnectionModel {
  portOut: NodePortModel;
  portIn: NodePortModel | undefined;
  mousePos: Vector2 | undefined;

  constructor(portOut: NodePortModel, portIn?: NodePortModel, mousePos?: Vector2) {
    this.portOut = portOut;
    this.portIn = portIn;
    this.mousePos = mousePos;
  }

  setMousePos(pos: Vector2) {
    this.mousePos = Vector2.subtract(pos, this.portOut.node.grid.posAbs);
  }

  getCurve(): BezierCurve {
    const pos1 = this.portOut?.gridPos;
    const pos4 = this.portIn ? this.portIn.pos : this.mousePos;
    let pos2 = pos1;
    let pos3 = pos4;
    if (pos4) {
      const centerPoint = pos1.x + (pos4.x - pos1.x) / 2;
      pos2 = new Vector2(centerPoint, pos1.y);
      pos3 = new Vector2(centerPoint, pos4.y);
    }
    return new BezierCurve(pos1, pos2, pos3, pos4);
  }
}

export default NodeConnectionModel;
