import Vector2 from "@/components/util/Vector";
import EditorComponent from "./EditorComponent";
import NodeConnectionModel from "./Node/NodeConnectionModel";
import NodeModel from "./Node/NodeModel";
import NodePortModel from "./Node/NodePortModel";

class GridModel extends EditorComponent {
  static cellSize = 20;

  connections: NodeConnectionModel[] = [];
  tmpConnectionIndex = -1; // the Connection that is currently dragged

  constructor(posRel = new Vector2(0, 0), ...nodes: NodeModel[]) {
    super("GridID", "Grid", false);
    this.posRel = posRel;
    nodes.forEach((node) => this.addChildren(node));
  }

  updatePos(): void {
    this.children.forEach((child) => child.updatePos);
  }

  // Can be integrated in EditorComponent recursive
  getPortByID(id: string): NodePortModel | undefined {
    for (const node of this.children) {
      const port = node.getChildById(id);
      if (port) {
        return port as NodePortModel;
      }
    }

    return undefined;
  }

  //#region Connections
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
