import Vector2 from "@/components/util/Vector";
import NodeConnectionModel from "./NodeConnectionModel";
import NodeModel from "./NodeModel";

class NodePortModel {
  id: string;
  title: string;
  isInput: boolean;
  node: NodeModel; // Parent Node
  pos: Vector2 = new Vector2(0, 0); // Pos relative to Parent
  posAbs: Vector2 = new Vector2(0, 0); // Absolute Pos
  gridPos: Vector2 = new Vector2(0, 0); // Pos relative to Grid with GridSnap
  connectedTo: NodePortModel[] = [];
  connections: NodeConnectionModel[] = [];
  tmpConnection: NodeConnectionModel | null = null; // Connection as long as it is moved by the mouse and not connected to a port

  constructor(id: string, title = "New Port", node: NodeModel, isInput = false) {
    this.id = id;
    this.title = title;
    this.node = node;
    this.isInput = isInput;
  }

  setPos(pos: Vector2) {
    this.pos = pos;
    this.updatePos();
  }

  updatePos() {
    this.gridPos = Vector2.add(this.pos, this.node.gridPos);
    this.posAbs = Vector2.add(this.gridPos, this.node.grid.posAbs);
  }

  setTmpConnection(connection: NodeConnectionModel | null) {
    this.tmpConnection = connection;
  }

  saveTmpConnection(portIn: NodePortModel) {
    if (!this.tmpConnection) return;

    this.tmpConnection?.setPortIn(portIn);
    this.connections.push(this.tmpConnection);
    portIn.connections.push(this.tmpConnection);
    this.tmpConnection = null;

    console.log(this.connections[0]);
  }

  getConnectionFromID(id: string): NodeConnectionModel | null {
    for (const connection of this.connections) {
      if (connection.id == id) {
        return connection;
      }
    }
    if (this.tmpConnection?.id == id) {
      return this.tmpConnection;
    }

    return null;
  }

  moveConnectionToTmp(id: string) {
    const connection: NodeConnectionModel | null = this.getConnectionFromID(id);
    if (!connection) return;

    this.tmpConnection = connection;
    const index = this.connections.indexOf(connection);
    if (index >= 0) {
      this.connections.splice(index, 1);
    }
  }

  clearConnections() {
    this.connections = [];
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
