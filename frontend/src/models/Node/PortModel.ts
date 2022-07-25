import EditorComponent from "../EditorComponent";
import { uuid } from "vue-uuid";

class PortModel extends EditorComponent {
  placeholder: string;
  isInput: boolean;
  portSize = 15;
  hasDefaultField: boolean;

  constructor(id: string, label = "New Port", isInput = false, hasDefaultField = false, placeholder = "") {
    super(id, label, false);
    this.isInput = isInput;
    this.hasDefaultField = hasDefaultField;
    this.placeholder = placeholder;
  }

  clone(): EditorComponent {
    const port = new PortModel("port-" + uuid.v4(), this.label, this.isInput, this.hasDefaultField, this.placeholder);
    if (this.parent) port.setParent(this.parent);
    return port;
  }

  //#region Serialization
  static fromJSON(json: JSON): PortModel {
    const port = new PortModel(
      json["id" as keyof JSON] as string,
      json["label" as keyof JSON] as string,
      JSON.parse(json["is-input" as keyof JSON] as string),
      true,
      json["default-value" as keyof JSON] as string
    );

    return port;
  }
  toJSON(): JSON {
    const json = JSON.parse(JSON.stringify({}));

    json["id"] = this.id;
    json["group-id"] = ""; // TODO
    json["added"] = false; // TODO
    json["label"] = this.label;
    json["is-input"] = this.isInput;
    json["datatype"] = "ANY"; // TODO
    json["default-value"] = ""; // TODO

    return json;
  }
  //#endregion Serialization
}

export default PortModel;
