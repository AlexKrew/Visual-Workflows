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
  gridPos: Vector2 = new Vector2(0, 0);
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

  setPos(pos: Vector2) {
    this.pos = pos;
    this.updatePos();
  }

  updatePos() {
    this.gridPos = Vector2.add(this.pos, this.node.gridPos);
    this.posAbs = Vector2.add(this.gridPos, this.node.grid.posAbs);
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
