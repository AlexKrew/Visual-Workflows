import EditorComponent from "../EditorComponent";
import { uuid } from "vue-uuid";

class PortModel extends EditorComponent {
  placeholder: string;
  isInput: boolean;
  portSize = 15;
  hasDefaultField: boolean;
  defaultValue = "";

  added = false;
  groupID = "";

  constructor(
    id: string,
    label = "New Port",
    isInput = false,
    hasDefaultField = false,
    placeholder = "",
    defaultValue = ""
  ) {
    super(id, label, false);
    this.isInput = isInput;
    this.hasDefaultField = hasDefaultField;
    this.placeholder = placeholder;
    this.defaultValue = defaultValue;
  }

  clone(): EditorComponent {
    const port = new PortModel("port-" + uuid.v4(), this.label, this.isInput, this.hasDefaultField, this.placeholder);
    if (this.parent) port.setParent(this.parent);
    return port;
  }

  setDefaultValue(text: string) {
    this.defaultValue = text;
  }

  setGroupID(id: string) {
    this.added = true;
    this.groupID = id;
  }

  //#region Serialization
  static fromJSON(json: JSON): PortModel {
    const port = new PortModel(
      json["id" as keyof JSON] as string,
      json["label" as keyof JSON] as string,
      JSON.parse(json["is-input" as keyof JSON] as string),
      true,
      json["default-value" as keyof JSON] as string,
      json["default-value" as keyof JSON] as string
    );

    // port.groupID = json["group-id" as keyof JSON] as string;
    // port.added = JSON.parse(json["added" as keyof JSON] as string);

    return port;
  }
  toJSON(): JSON {
    const json = JSON.parse(JSON.stringify({}));

    json["id"] = this.id;
    json["group-id"] = this.groupID; // TODO
    json["added"] = this.added; // TODO
    json["label"] = this.label;
    json["is-input"] = this.isInput;
    json["datatype"] = "ANY"; // TODO
    json["default-value"] = this.defaultValue;

    return json;
  }
  //#endregion Serialization
}

export default PortModel;
