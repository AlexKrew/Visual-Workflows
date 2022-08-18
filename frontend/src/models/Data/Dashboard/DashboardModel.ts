import DashboardElement from "./DashboardElement";
import { UpdateFieldType } from "./UITypes";

class DashboardModel {
  static canvas: DashboardElement;

  static getElementByID(id: string, element: DashboardElement = DashboardModel.canvas): DashboardElement | undefined {
    if (element.data.id == id) return element;

    let result: DashboardElement | undefined;
    element.children.forEach((child) => {
      const element = DashboardModel.getElementByID(id, child);
      if (element) result = element;
    });

    return result;
  }

  static updateFields(...newFields: UpdateFieldType[]) {
    newFields.forEach((newField) => {
      const element = DashboardModel.getElementByID(newField.id);
      if (element) element.data.fields[newField.field] = newField.value;
    });
  }
}

export default DashboardModel;
