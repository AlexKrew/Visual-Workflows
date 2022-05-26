import BezierCurve from "@/components/util/BezierCurve";
import Vector2 from "@/components/util/Vector";
import GridModel from "./GridModel";
import NodeModel from "./NodeModel";
import NodePortModel from "./NodePortModel";

class TestModels {
  static getGrid(): GridModel {
    const grid = new GridModel(20, new Vector2(250, 50));
    grid.addNodes([
      new NodeModel("N100", "Node1", grid),
      new NodeModel("N200", "Node2", grid),
      new NodeModel("N300", "Node3", grid),
    ]);
    grid.nodes[0].addPorts([
      new NodePortModel("P101", "Port 1", grid.nodes[0], [], true),
      new NodePortModel("P102", "Port 2", grid.nodes[0], [], false),
    ]);
    grid.nodes[1].addPorts([
      new NodePortModel("P201", "Port 1", grid.nodes[1], [], true),
      new NodePortModel("P202", "Port 2", grid.nodes[1], [], true),
    ]);
    grid.nodes[2].addPorts([
      new NodePortModel("P301", "Port 1", grid.nodes[2], [], false),
      new NodePortModel("P302", "Port 2", grid.nodes[2], [], false),
    ]);

    return grid;
  }

  static getCurve(): BezierCurve {
    return new BezierCurve(new Vector2(100, 100), new Vector2(130, 130), new Vector2(330, 130), new Vector2(300, 100));
  }
}

export default TestModels;
