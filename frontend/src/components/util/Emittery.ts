import NodeModel from "@/models/Node/NodeModel";
import Emittery from "emittery";
const emitter = new Emittery<{
  PortsUpdatePos: NodeModel;
  UpdateWorkflowEditor: undefined;
  UpdateDashboard: undefined;
  UpdateNavBar: [number, string];
}>();

export { emitter, Emittery };
