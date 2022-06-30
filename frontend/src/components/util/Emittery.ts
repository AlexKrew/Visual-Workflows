import NodeModel from "@/models/Node/NodeModel";
import Emittery from "emittery";
const emitter = new Emittery<{ PortsUpdatePos: NodeModel }>();

export { emitter, Emittery };
