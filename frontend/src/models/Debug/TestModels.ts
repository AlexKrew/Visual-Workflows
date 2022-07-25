import Vector2 from "@/components/util/Vector";
import GridModel from "../GridModel";
import NodeModel from "../Node/NodeModel";
import PortModel from "../Node/PortModel";

class TestModels {
  static nodeCategorys = ["Web Services", "Debug", "Control"];

  // Nodes
  static nodes = [
    new NodeModel("N1", "HTTP Request", this.nodeCategorys[0], [
      new PortModel("N1-P1", "Method", true),
      new PortModel("N1-P2", "URL", true, true, "someURL.com"),
      new PortModel("N1-P3", "Payload", true, true, "Header:{...}"),
      new PortModel("N1-P4", "Response", false),
      new PortModel("N1-P5", "Response Code", false),
    ]),
    new NodeModel("N2", "Debug", this.nodeCategorys[1], [
      new PortModel("N2-P1", "Input", true),
      new PortModel("N2-P2", "Output", false),
    ]),
    new NodeModel("N3", "Mail", this.nodeCategorys[0], [
      new PortModel("N3-P1", "E-Mail", true, true, "some@email.com"),
      new PortModel("P3-P2", "Message", true, true, "Dear Someone, ..."),
    ]),
    new NodeModel(
      "N4",
      "Switch",
      this.nodeCategorys[2],
      [new PortModel("N4-P1", "Check", true)],
      [
        new PortModel("N4-AP1", "Operator", true, true, "=="),
        new PortModel("N4-AP2", "Value", true, true),
        new PortModel("N4-AP3", "In", true, true, "In => Out"),
        new PortModel("N4-AP4", "Out", false),
      ]
    ),
  ];

  // Grid
  static readonly grid = new GridModel(new Vector2(220, 0), [
    TestModels.nodes[0],
    TestModels.nodes[1],
    TestModels.nodes[2],
    // TestModels.nodes[3],
  ]);
}

export default TestModels;
