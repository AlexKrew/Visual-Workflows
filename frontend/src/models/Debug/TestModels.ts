import Vector2 from "@/components/util/Vector";
import GridModel from "../GridModel";
import NodeModel from "../Node/NodeModel";
import NodePortModel from "../Node/NodePortModel";

class TestModels {
  static nodeCategorys = ["Web Services", "Debug"];

  // Nodes
  static nodes = [
    new NodeModel(
      "N100",
      "HTTP Request",
      this.nodeCategorys[0],
      new NodePortModel("P100", "Method", true),
      new NodePortModel("P101", "URL", true, true),
      new NodePortModel("P102", "Payload", true, true),
      new NodePortModel("P103", "Response", false),
      new NodePortModel("P104", "Response Code", false)
    ),
    new NodeModel("N200", "Debug", this.nodeCategorys[1], new NodePortModel("P200", "Input", true, true)),
    new NodeModel(
      "N300",
      "Mail",
      this.nodeCategorys[0],
      new NodePortModel("P300", "E-Mail", true, true),
      new NodePortModel("P301", "Message", true, true)
    ),
  ];

  // Grid
  static readonly grid = new GridModel(
    new Vector2(200, 20),
    TestModels.nodes[0],
    TestModels.nodes[1],
    TestModels.nodes[2]
  );
}

export default TestModels;
