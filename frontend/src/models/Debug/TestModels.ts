import Vector2 from "@/components/util/Vector";
import GridModel from "../GridModel";
import NodeModel from "../Node/NodeModel";
import NodePortModel from "../Node/NodePortModel";

class TestModels {
  static nodeCategorys = ["Web Services", "Debug", "Control"];

  // Nodes
  static nodes = [
    new NodeModel("N1", "HTTP Request", this.nodeCategorys[0], [
      new NodePortModel("N1-P1", "Method", true),
      new NodePortModel("N1-P2", "URL", true, true, "someURL.com"),
      new NodePortModel("N1-P3", "Payload", true, true, "Header:{...}"),
      new NodePortModel("N1-P4", "Response", false),
      new NodePortModel("N1-P5", "Response Code", false),
    ]),
    new NodeModel("N2", "Debug", this.nodeCategorys[1], [new NodePortModel("N2-P1", "Input", true, true, "Print")]),
    new NodeModel("N3", "Mail", this.nodeCategorys[0], [
      new NodePortModel("N3-P1", "E-Mail", true, true, "some@email.com"),
      new NodePortModel("P3-P2", "Message", true, true, "Dear Someone, ..."),
    ]),
    new NodeModel(
      "N4",
      "Switch",
      this.nodeCategorys[2],
      [new NodePortModel("N4-P1", "Check", true)],
      [
        new NodePortModel("N4-AP1", "Operator", true, true, "=="),
        new NodePortModel("N4-AP2", "Value", true, true),
        new NodePortModel("N4-AP3", "In", true, true, "In => Out"),
        new NodePortModel("N4-AP4", "Out", false),
      ]
    ),
  ];

  // Grid
  static readonly grid = new GridModel(new Vector2(200, 20), [
    TestModels.nodes[0],
    // TestModels.nodes[1],
    // TestModels.nodes[2],
    TestModels.nodes[3],
  ]);
}

export default TestModels;
