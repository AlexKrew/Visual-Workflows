import GridModel from "./GridModel";
import NodeModel from "./NodeModel";
import NodePortModel from "./NodePortModel";

class TestModels {
  static getGrid(): GridModel {
    const res = new GridModel(undefined, [
      new NodeModel("100", "Node1"),
      new NodeModel("200", "Node2"),
      new NodeModel("300", "Node3"),
    ]);
    res.nodes[0].addPorts([
      new NodePortModel("101", "Port 1", undefined, res.nodes[0], null, true),
      new NodePortModel("102", "Port 2", undefined, res.nodes[0], null, false),
    ]);
    res.nodes[1].addPorts([
      new NodePortModel("201", "Port 1", undefined, res.nodes[1], null, true),
      new NodePortModel("202", "Port 2", undefined, res.nodes[1], null, true),
    ]);
    res.nodes[2].addPorts([
      new NodePortModel("201", "Port 1", undefined, res.nodes[2], null, false),
      new NodePortModel("202", "Port 2", undefined, res.nodes[2], null, false),
    ]);

    return res;
  }
}

export default TestModels;
