import Vector2 from "@/components/util/Vector";
import NodeConnectionModel from "./NodeConnectionModel";
import NodeModel from "./NodeModel";

class NodePortModel {
  id: string;
  title: string;
  isInput: boolean;
  node: NodeModel;
  pos: Vector2 = new Vector2(0, 0);
  posAbs: Vector2 = new Vector2(0, 0);
  connectedTo: NodePortModel[] = [];
  connections: NodeConnectionModel[] = [];
  tmpConnection: NodeConnectionModel | null = null;

  constructor(id: string, title = "New Port", node: NodeModel, isInput = false) {
    this.id = id;
    this.title = title;
    this.node = node;
    this.isInput = isInput;
  }

  setTmpConnection(connection: NodeConnectionModel) {
    this.tmpConnection = connection;
  }

  setPosAbs(posAbs: Vector2) {
    this.posAbs = posAbs;
    this.updatePos();
  }

  updatePos() {
    this.pos = Vector2.subtract(this.posAbs, this.node.gridPos, this.node.grid.posAbs);
  }

  updatePosAbs() {
    this.posAbs = Vector2.add(this.pos, this.node.posAbs);
  }

  getConnectionPos(): Vector2 {
    return Vector2.add(this.pos, this.node.gridPos);
  }

  connectTo(port: NodePortModel, connection?: NodeConnectionModel, removeTmpConnecion?: boolean) {
    this.connectedTo.push(port);
    if (!this.isInput && connection) {
      this.connections.push(connection);
    }
    if (removeTmpConnecion) {
      this.tmpConnection = null;
    }
  }
}

export default NodePortModel;
