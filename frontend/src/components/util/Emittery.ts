import NodeModel from "@/models/Node/NodeModel";
import Emittery from "emittery";
const emitter = new Emittery<{
  PortsUpdatePos: NodeModel;
  UpdatePort: string;
  UpdateWorkflowEditor: undefined;
  UpdateDashboard: undefined;
  UpdateNavBar: [number, string];
  OpenImportExportModal: boolean;
}>();

export { emitter, Emittery };
