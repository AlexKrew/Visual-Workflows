import Vector2 from "@/components/util/Vector";
import EditorComponent from "../EditorComponent";
import { uuid } from "vue-uuid";

class NodePortModel extends EditorComponent {
  placeholder: string;
  isInput: boolean;
  portSize = 15;
  hasDefaultField: boolean;
  textAreaScrollHeight = 24;

  constructor(id: string, title = "New Port", isInput = false, hasDefaultField = false, placeholder = "") {
    super(id, title, false);
    this.isInput = isInput;
    this.hasDefaultField = hasDefaultField;
    this.placeholder = placeholder;
  }

  clone(): EditorComponent {
    return new NodePortModel(uuid.v4(), this.title, this.isInput, this.hasDefaultField, this.placeholder);
  }

  changeTextAreaHeight(newHeight: number, resizeChilds = false) {
    if (this.textAreaScrollHeight == newHeight) return;

    const dif = newHeight - this.textAreaScrollHeight;
    this.textAreaScrollHeight = newHeight;

    if (resizeChilds) this.parent?.addPosToChildren(new Vector2(0, dif), this.parent.getChildIndex(this.id) + 1);
  }
}

export default NodePortModel;
