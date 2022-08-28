import { emitter } from "@/components/util/Emittery";
import DashboardElement from "./DashboardElement";
import { UpdateFieldType } from "./UITypes";

class DashboardModel {
  static canvas: DashboardElement;

  static setCanvas(canvas: DashboardElement) {
    DashboardModel.canvas = canvas;
    emitter.emit("UpdateDashboard");
  }

  static getElementByID(id: string, element: DashboardElement = DashboardModel.canvas): DashboardElement | undefined {
    if (element.data.id == id) return element;

    let result: DashboardElement | undefined;
    element.children.forEach((child) => {
      const element = DashboardModel.getElementByID(id, child);
      if (element) result = element;
    });

    return result;
  }

  static updateField(newField: UpdateFieldType) {
    const element = DashboardModel.getElementByID(newField.id);
    if (element) element.data.fields[newField.field] = newField.value;
  }
}

export default DashboardModel;
