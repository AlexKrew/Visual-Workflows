import { emitter } from "@/components/util/Emittery";
import Vector2 from "@/components/util/Vector";
import GridData from "./Data/GridData";
import { EdgeType, WorkflowType } from "./Data/Types";
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
    this.posRel = new Vector2(0, 64);

    data.nodes.forEach((node) => {
      this.addChild(new NodeModel(node), false);
    });
  }

  loadEdges() {
    this.edges = [];
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

  removeChild(id: string): void {
    // Remove Edges
    this.removeEdges(undefined, id, undefined);

    // Remove Child From Data
    const childData = this.data.nodes.find((node) => node.id == id);
    if (childData) this.data.nodes.splice(this.data.nodes.indexOf(childData), 1);

    //Remove Child
    const child = this.children.find((child) => child.id == id);
    if (child) {
      this.children.splice(this.children.indexOf(child), 1);
    }
  }

  removeEdges(edgeID?: string, nodeID?: string, portID?: string) {
    let edges: EdgeType[] = [];

    if (edgeID) edges = this.data.edges.filter((edge) => edge.id == edgeID);
    else if (nodeID)
      edges = this.data.edges.filter((edge) => edge.origin.node_id == nodeID || edge.target.node_id == nodeID);
    else if (portID)
      edges = this.data.edges.filter((edge) => edge.origin.port_id == portID || edge.target.port_id == portID);
    else return;

    edges.forEach((edge) => this.deleteEdge(edge.id));
  }

  updatePos(): void {
    this.children.forEach((child) => child.updatePos);
  }

  // Can be integrated in EditorComponent recursive
  getPortByID(id: string): PortModel | undefined {
    let result: PortModel | undefined = undefined;

    this.children.forEach((node) => {
      const port = (node as NodeModel).getChildById(id);
      if (port) result = port as PortModel;
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
    this.data.edges.splice(
      this.data.edges.findIndex((edge) => edge.id == id),
      1
    );

    this.loadEdges();
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
