import Vector2 from "@/components/util/Vector";
import GridModel from "./GridModel";
import NodeModel from "./NodeModel";
import NodePortModel from "./NodePortModel";

class TestModels {
  // Nodes
  static nodes = [
    new NodeModel(
      "N100",
      "HTTP Request",
      new NodePortModel("P100", "Method", true),
      new NodePortModel("P101", "URL", true),
      new NodePortModel("P102", "Payload", true),
      new NodePortModel("P103", "Response", false),
      new NodePortModel("P104", "Response Code", false)
    ),
    new NodeModel("N200", "Debug", new NodePortModel("P200", "Input", true)),
    new NodeModel(
      "N300",
      "Mail",
      new NodePortModel("P300", "E-Mail", true),
      new NodePortModel("P301", "Message", true)
    ),
  ];

  // Grid
  static readonly grid = new GridModel(
    20,
    new Vector2(200, 20),
    TestModels.nodes[0],
    TestModels.nodes[1],
    TestModels.nodes[2]
  );
}

export default TestModels;
