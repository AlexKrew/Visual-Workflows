import Vector2 from "@/components/util/Vector";
import NodeConnectionModel from "./NodeConnectionModel";
import NodeModel from "./NodeModel";
import NodePortModel from "./NodePortModel";

class GridModel {
  posAbs: Vector2;
  cellSize: number;
  nodes: NodeModel[] = [];
  connections: NodeConnectionModel[] = [];
  tmpConnectionIndex = -1; // the Connection that is currently dragged

  constructor(cellSize = 10, posAbs = new Vector2(0, 0)) {
    this.posAbs = posAbs;
    this.cellSize = cellSize;
  }

  addPos(posAbs: Vector2) {
    this.posAbs.addVector(posAbs);
    this.nodes.forEach((node) => node.updatePos());
  }

  addNodes(...nodes: NodeModel[]) {
    nodes.forEach((node) => this.nodes.push(node));
  }

  getPortByID(id: string): NodePortModel | undefined {
    for (const node of this.nodes) {
      const port = node.getPortByID(id);
      if (port) {
        return port;
      }
    }

    return undefined;
  }

  //#region +++++ Connections +++++
  addConnection(connection: NodeConnectionModel, isTmp = false) {
    this.connections.push(connection);
    if (isTmp) {
      this.tmpConnectionIndex = this.connections.length - 1;
    }
  }

  deleteConnection(id: string) {
    const index = this.getConnectionIndex(id);
    if (index < 0) {
      throw new Error(`Can not delete Connection: ${id}`);
    }
    if (index == this.tmpConnectionIndex) {
      this.tmpConnectionIndex = -1;
    }
    this.connections.splice(index, 1);
  }

  setTmp(id: string) {
    this.tmpConnectionIndex = this.getConnectionIndex(id);
  }

  resetTmp(deleteConnection = false) {
    if (this.tmpConnectionIndex < 0) return;
    if (deleteConnection) {
      this.deleteConnection(this.connections[this.tmpConnectionIndex].id);
    }
    this.tmpConnectionIndex = -1;
  }

  getConnection(connectionId?: string, portInId?: string): NodeConnectionModel {
    const connection = this.connections[this.getConnectionIndex(connectionId, portInId)];
    return connection;
  }

  getTmpConnection(): NodeConnectionModel {
    return this.connections[this.tmpConnectionIndex];
  }

  private getConnectionIndex(connectionId?: string, portInId?: string) {
    const index = this.connections.findIndex((connection) => {
      if (connectionId) {
        return connection.id == connectionId;
      } else if (portInId) {
        return connection.portIn?.id == portInId;
      }
      return false;
    });

    if (index < 0) {
      throw new Error(`No Connection Index found with id: ${connectionId}, or portInID: ${portInId}`);
    }
    return index;
  }

  //#endregion
}

export default GridModel;
