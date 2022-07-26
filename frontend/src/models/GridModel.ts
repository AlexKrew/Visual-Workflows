import Vector2 from "@/components/util/Vector";
import { WorkflowType } from "./Data/Types";
import EditorComponent from "./EditorComponent";
import EdgeModel from "./Node/EdgeModel";
import NodeModel from "./Node/NodeModel";
import PortModel from "./Node/PortModel";

class GridModel extends EditorComponent {
  data: WorkflowType;

  edges: EdgeModel[] = [];
  tmpEdgeIndex = -1; // the Connection that is currently dragged

  constructor(data: WorkflowType) {
    super(data.id, false);
    this.data = data;
    this.posRel = new Vector2(200, 0);

    data.nodes.forEach((node) => {
      this.addChild(new NodeModel(node), false);
    });
  }

  loadEdges() {
    this.data.edges.forEach((edge) => {
      this.edges.push(new EdgeModel(edge));
    });
  }

  clone(): EditorComponent {
    throw new Error("Method not implemented.");
  }

  addChildToData(child: EditorComponent): void {
    this.data.nodes.push((child as NodeModel).data);
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

  //#region Edges
  addEdge(edge: EdgeModel, isTmp = false) {
    this.data.edges.push(edge.data);
    this.edges.push(edge);
    if (isTmp) {
      this.tmpEdgeIndex = this.edges.length - 1;
    }
  }

  deleteEdge(id: string) {
    // Delete Edge from Data
    this.data.edges.splice(
      this.data.edges.findIndex((edge) => {
        edge.id == id;
      }),
      1
    );

    // Delete Edge From Edges
    this.edges.splice(
      this.edges.findIndex((edge) => {
        edge.data.id == id;
      }),
      1
    );
  }

  setTmp(id: string) {
    this.tmpEdgeIndex = this.getEdgeIndex(id);
  }

  resetTmp(deleteEdge = false) {
    if (this.tmpEdgeIndex < 0) return;
    if (deleteEdge) {
      this.deleteEdge(this.edges[this.tmpEdgeIndex].data.id);
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
        return edge.data.id == edgeID;
      } else if (portInId) {
        return edge.target?.id == portInId;
      }
      return false;
    });

    return index;
  }
  //#endregion

  updatePosOverload(): void {
    return;
  }
}

export default GridModel;
