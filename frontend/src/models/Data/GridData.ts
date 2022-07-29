import GridModel from "../GridModel";
import NodeModel from "../Node/NodeModel";
import NodesJSON from "./JSON/Nodes.json";
import { LogType, NodeType, WorkflowType } from "./Types";

class GridData {
  static nodeCategorys = ["Web Services", "Debug", "Control"]; // TODO Dynamic
  static nodeTypes = ["Http", "Debug", "Mail", "Switch"]; // TODO Dynamic
  static logs: LogType[] = [];
  static cellSize = 20;

  static workflow: WorkflowType;
  static grid: GridModel;
  static nodes: NodeModel[] = [];

  static loadDefaultNodes() {
    const nodeTypes: NodeType[] = JSON.parse(JSON.stringify(NodesJSON));
    nodeTypes.forEach((node) => {
      GridData.nodes.push(new NodeModel(node));
    });
  }

  static loadWorkflow(json: JSON) {
    this.workflow = JSON.parse(JSON.stringify(json));
    GridData.grid = new GridModel(this.workflow);
    GridData.grid.loadEdges();
  }
}

export default GridData;
