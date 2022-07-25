import Vector2 from "@/components/util/Vector";
import EditorComponent from "./EditorComponent";
import ISerializable from "./ISerializable";
import EdgeModel from "./Node/EdgeModel";
import NodeModel from "./Node/NodeModel";
import PortModel from "./Node/PortModel";

class GridModel extends EditorComponent {
  static cellSize = 20;

  edges: EdgeModel[] = [];
  tmpEdgeIndex = -1; // the Connection that is currently dragged

  constructor(posRel = new Vector2(0, 0), nodes: NodeModel[]) {
    super("GridID", "Grid", false, nodes);
    this.posRel = posRel;
  }

  clone(): EditorComponent {
    throw new Error("Method not implemented.");
  }

  updatePos(): void {
    this.children.forEach((child) => child.updatePos);
  }

  // Can be integrated in EditorComponent recursive
  getPortByID(id: string): PortModel | undefined {
    let result = undefined;

    this.children.forEach((node) => {
      const port = (node as NodeModel).getChildById(id);
      if (port) result = port;
    });

    return result;
  }

  //#region Connections
  addEdge(edge: EdgeModel, isTmp = false) {
    this.edges.push(edge);
    if (isTmp) {
      this.tmpEdgeIndex = this.edges.length - 1;
    }
  }

  deleteEdge(id: string) {
    const index = this.getEdgeIndex(id);
    if (index < 0) {
      throw new Error(`Can not delete Connection: ${id}`);
    }
    if (index == this.tmpEdgeIndex) {
      this.tmpEdgeIndex = -1;
    }
    this.edges.splice(index, 1);
  }

  setTmp(id: string) {
    this.tmpEdgeIndex = this.getEdgeIndex(id);
  }

  resetTmp(deleteEdge = false) {
    if (this.tmpEdgeIndex < 0) return;
    if (deleteEdge) {
      this.deleteEdge(this.edges[this.tmpEdgeIndex].id);
    }
    this.tmpEdgeIndex = -1;
  }

  getEdge(edgeID?: string, portInId?: string): EdgeModel | undefined {
    const index = this.getEdgeIndex(edgeID, portInId);

    if (index < 0) return undefined;

    const connection = this.edges[index];
    return connection;
  }

  getTmpEdge(): EdgeModel {
    return this.edges[this.tmpEdgeIndex];
  }

  private getEdgeIndex(edgeID?: string, portInId?: string) {
    const index = this.edges.findIndex((edge) => {
      if (edgeID) {
        return edge.id == edgeID;
      } else if (portInId) {
        return edge.portIn?.id == portInId;
      }
      return false;
    });

    return index;
  }

  //#endregion

  //#region Serialization
  fromJSON(json: JSON): ISerializable {
    throw new Error("Method not implemented.");
  }

  toJSON(): JSON {
    const json = JSON.parse(JSON.stringify({}));

    json["id"] = this.id;
    json["name"] = "Workflow Name"; //TODO
    json["nodes"] = [];
    json["edges"] = [];

    this.children.forEach((child) => {
      json["nodes"].push(child.toJSON());
    });
    this.edges.forEach((edge) => {
      json["edges"].push(edge.toJSON());
    });

    return json;
  }
  //#endregion Serialization
}

export default GridModel;
