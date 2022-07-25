import BezierCurve from "@/components/util/BezierCurve";
import Vector2 from "@/components/util/Vector";
import NodePortModel from "./PortModel";
import { uuid } from "vue-uuid";
import ISerializable from "../ISerializable";

class EdgeModel implements ISerializable {
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
    const pos1 = this.portOut?.posGridCell;
    let pos2 = new Vector2(0, 0);
    let pos3 = new Vector2(0, 0);
    const pos4 = this.portIn ? this.portIn.posGridCell : this.mousePos;

    if (pos4) {
      const dist = Vector2.dist(pos1, pos4);
      const multi = 0.5 * dist;
      pos2 = Vector2.add(pos1, new Vector2(multi, 0));
      pos3 = Vector2.add(pos4, new Vector2(-multi, 0));
    }
    return new BezierCurve(pos1, pos2, pos3, pos4);
  }

  //#region Serialization
  fromJSON(json: JSON): ISerializable {
    throw new Error("Method not implemented.");
  }
  toJSON(): JSON {
    const json = JSON.parse(JSON.stringify({}));

    json["id"] = this.id;
    json["target"]["node"] = this.portIn?.parent;
    json["target"]["port"] = this.portIn;
    json["origin"]["node"] = this.portOut?.parent;
    json["origin"]["port"] = this.portOut;

    return json;
  }
  //#endregion
}

export default EdgeModel;
