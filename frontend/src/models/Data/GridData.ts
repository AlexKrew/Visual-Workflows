import GridModel from "../GridModel";
import NodeModel from "../Node/NodeModel";
import NodesJSON from "./JSON/Nodes.json";
import { LogType, NodeType, PortType, WorkflowType } from "./Types";

class GridData {
  static nodeCategorys: string[] = [];
  static nodeTypes: string[] = [];
  static logs: LogType[] = [];
  static cellSize = 20;

  static workflow: WorkflowType;
  static grid: GridModel;
  static nodes: NodeModel[] = [];

  static loadDefaultData() {
    GridData.nodes = [];

    // Load default Nodes
    const nodeTypes: NodeType[] = JSON.parse(JSON.stringify(NodesJSON));
    for (let i = 0; i < nodeTypes.length; i++) {
      GridData.nodes.push(new NodeModel(nodeTypes[i]));
    }

    // Load Categorys and Types
    const categorys: string[] = [];
    const types: string[] = [];
    nodeTypes.forEach((nodeType) => {
      categorys.push(nodeType.category);
      types.push(nodeType.type);
    });
    GridData.nodeCategorys = [...new Set(categorys)];
    GridData.nodeTypes = [...new Set(types)];
  }

  static loadWorkflow(json: JSON) {
    this.workflow = JSON.parse(JSON.stringify(json));
    GridData.grid = new GridModel(this.workflow);
    GridData.grid.loadEdges();
  }

  static getWorkflowName(): string {
    if (GridData.grid && GridData.grid.data) {
      return GridData.grid.data.name;
    }
    return "";
  }
}

export default GridData;
