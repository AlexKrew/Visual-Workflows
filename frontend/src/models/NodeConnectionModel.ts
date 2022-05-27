import BezierCurve from "@/components/util/BezierCurve";
import Vector2 from "@/components/util/Vector";
import NodePortModel from "./NodePortModel";
import { uuid } from "vue-uuid";

class NodeConnectionModel {
  id: string;
  portOut: NodePortModel;
  portIn: NodePortModel | undefined;
  mousePos: Vector2 | undefined;

  constructor(portOut: NodePortModel, portIn?: NodePortModel, mousePos?: Vector2) {
    this.id = uuid.v4();
    this.portOut = portOut;
    this.portIn = portIn;
    this.mousePos = mousePos;
  }

  setMousePos(pos: Vector2) {
    this.mousePos = Vector2.subtract(pos, this.portOut.node.grid.posAbs);
  }

  setPortIn(port: NodePortModel | undefined) {
    this.portIn = port;
  }

  getCurve(): BezierCurve {
    const pos1 = this.portOut?.gridPos;
    const pos4 = this.portIn ? this.portIn.gridPos : this.mousePos;
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
