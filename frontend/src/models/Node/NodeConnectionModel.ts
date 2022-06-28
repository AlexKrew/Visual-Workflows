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
    if (!this.portOut.parent?.parent) return;
    this.mousePos = Vector2.subtract(pos, this.portOut.parent.parent.posRel);
  }

  setPortIn(port: NodePortModel | undefined) {
    this.portIn = port;
  }

  getCurve(): BezierCurve {
    const pos1 = this.portOut?.posGrid;
    let pos2 = new Vector2(0, 0);
    let pos3 = new Vector2(0, 0);
    const pos4 = this.portIn ? this.portIn.posGrid : this.mousePos;

    if (pos4) {
      const dist = Vector2.dist(pos1, pos4);
      const multi = 0.5 * dist;
      pos2 = Vector2.add(pos1, new Vector2(multi, 0));
      pos3 = Vector2.add(pos4, new Vector2(-multi, 0));
    }
    return new BezierCurve(pos1, pos2, pos3, pos4);
  }
}

export default NodeConnectionModel;
