import { BezierCurve } from "@/components/util/BezierCurve";
import { Vector2 } from "@/components/util/Vector";
import GridModel from "./GridModel";
import NodeModel from "./NodeModel";
import NodePortModel from "./NodePortModel";

class TestModels {
  static getGrid(): GridModel {
    const grid = new GridModel(20, undefined);
    grid.addNodes([
      new NodeModel("N100", "Node1", grid),
      new NodeModel("N200", "Node2", grid),
      new NodeModel("N300", "Node3", grid),
    ]);
    grid.nodes[0].addPorts([
      new NodePortModel("P101", "Port 1", undefined, grid.nodes[0], null, true),
      new NodePortModel("P102", "Port 2", undefined, grid.nodes[0], null, false),
    ]);
    grid.nodes[1].addPorts([
      new NodePortModel("P201", "Port 1", undefined, grid.nodes[1], null, true),
      new NodePortModel("P202", "Port 2", undefined, grid.nodes[1], null, true),
    ]);
    grid.nodes[2].addPorts([
      new NodePortModel("P201", "Port 1", undefined, grid.nodes[2], null, false),
      new NodePortModel("P202", "Port 2", undefined, grid.nodes[2], null, false),
    ]);

    return grid;
  }

  static getCurve(): BezierCurve {
    return new BezierCurve(new Vector2(100, 100), new Vector2(130, 130), new Vector2(330, 130), new Vector2(300, 100));
  }
}

export default TestModels;
