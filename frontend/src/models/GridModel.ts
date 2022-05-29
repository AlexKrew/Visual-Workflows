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

  getPortByID(id: string): NodePortModel | null {
    for (const node of this.nodes) {
      const port = node.getPortByID(id);
      if (port) {
        return port;
      }
    }

    return null;
  }

  //#region +++++ Connections +++++
  addConnection(connection: NodeConnectionModel, isTmp = false) {
    this.connections.push(connection);
    if (isTmp) {
      this.tmpConnectionIndex = this.connections.length - 1;
    }
  }

  getConnectionFromPortInID(id: string): NodeConnectionModel | undefined {
    return this.connections.find((connection) => connection.portIn?.id == id);
  }

  resetTmpConnection(deleteConnection = false) {
    if (this.tmpConnectionIndex < 0) return;
    if (deleteConnection) {
      this.connections.splice(this.tmpConnectionIndex, 1);
    }
    this.tmpConnectionIndex = -1;
  }

  saveTmpConnection(portIn?: NodePortModel, portInID?: string) {
    if (portIn) {
      this.connections[this.tmpConnectionIndex].setPortIn(portIn);
    } else if (portInID) {
      const port = this.getPortByID(portInID);
      if (port) {
        this.connections[this.tmpConnectionIndex].setPortIn(port);
      }
    }
    this.resetTmpConnection(false);
  }

  connectionToTmp(connection: NodeConnectionModel) {
    const index = this.connections.indexOf(connection);
    if (index >= 0) {
      this.tmpConnectionIndex = index;
      connection.portIn = undefined;
    }
  }
  //#endregion
}

export default GridModel;
