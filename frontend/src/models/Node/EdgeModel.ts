import BezierCurve from "@/components/util/BezierCurve";
import Vector2 from "@/components/util/Vector";
import NodePortModel from "./PortModel";
import { EdgeType } from "../Data/Types";
import GridData from "../Data/GridData";
import NodeModel from "./NodeModel";
import { uuid } from "vue-uuid";

class EdgeModel {
  data: EdgeType;
  origin: NodePortModel | undefined;
  target: NodePortModel | undefined;
  mousePos: Vector2 | undefined;

  constructor(data: EdgeType) {
    this.data = data;
    this.setPortsFromData();
  }

  static NewEdgeFromPort(originPort: NodePortModel): EdgeModel {
    const newEdge: EdgeType = {
      id: uuid.v4(),
      origin: {
        node_id: (originPort.parent as NodeModel).data.id,
        port_id: originPort.data.id,
      },
      target: {
        node_id: "",
        port_id: "",
      },
    };

    return new EdgeModel(newEdge);
  }

  setPortsFromData() {
    this.target = GridData.grid.getPortByID(this.data.target.port_id);
    const originPort = GridData.grid.getPortByID(this.data.origin.port_id);
    if (originPort) this.origin = originPort;
    else throw new Error("No Port found with ID: " + this.data.origin.port_id);
  }

  setPortIn(port: NodePortModel | undefined) {
    this.data.target.node_id = port ? (port.parent as NodeModel).data.id : "";
    this.data.target.port_id = port ? port.data.id : "";
    this.setPortsFromData();
  }

  setMousePos(pos: Vector2) {
    if (!this.origin?.parent?.parent) return;
    this.mousePos = Vector2.subtract(pos, this.origin.parent.parent.posRel);
  }

  getCurve(): BezierCurve {
    const pos1 = this.origin ? this.origin.posGridCell : new Vector2(0, 0);
    let pos2 = new Vector2(0, 0);
    let pos3 = new Vector2(0, 0);
    const pos4 = this.target ? this.target.posGridCell : this.mousePos;

    if (pos4) {
      const dist = Vector2.dist(pos1, pos4);
      const multi = 0.5 * dist;
      pos2 = Vector2.add(pos1, new Vector2(multi, 0));
      pos3 = Vector2.add(pos4, new Vector2(-multi, 0));
    }
    return new BezierCurve(pos1, pos2, pos3, pos4);
  }
}

export default EdgeModel;
