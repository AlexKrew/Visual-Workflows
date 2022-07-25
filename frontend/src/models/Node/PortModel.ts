import EditorComponent from "../EditorComponent";
import { uuid } from "vue-uuid";
import ISerializable from "../ISerializable";

class PortModel extends EditorComponent {
  placeholder: string;
  isInput: boolean;
  portSize = 15;
  hasDefaultField: boolean;

  constructor(id: string, title = "New Port", isInput = false, hasDefaultField = false, placeholder = "") {
    super(id, title, false);
    this.isInput = isInput;
    this.hasDefaultField = hasDefaultField;
    this.placeholder = placeholder;
  }

  clone(): EditorComponent {
    const port = new PortModel("port-" + uuid.v4(), this.title, this.isInput, this.hasDefaultField, this.placeholder);
    if (this.parent) port.setParent(this.parent);
    return port;
  }

  //#region Serialization
  fromJSON(json: JSON): ISerializable {
    throw new Error("Method not implemented.");
  }
  toJSON(): JSON {
    const json = JSON.parse(JSON.stringify({}));

    json["id"] = this.id;
    json["group-id"] = ""; // TODO
    json["added"] = false; // TODO
    json["label"] = this.title;
    json["is-input"] = this.isInput;
    json["datatype"] = "ANY"; // TODO
    json["default-value"] = ""; // TODO

    return json;
  }
  //#endregion Serialization
}

export default PortModel;
