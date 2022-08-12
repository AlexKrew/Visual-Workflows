import { UICanvas, UIElement, UIList, UIText, uiTypes } from "./UITypes";

class DashboardElement {
  data: UIElement;
  children: DashboardElement[] = [];

  constructor(data: UIElement) {
    this.data = data;
    if (this.data.children) {
      this.data.children.forEach((child) => this.children.push(new DashboardElement(child)));
    }
  }

  // createDElement(data: UIList | UIText): DList | DText | void {
  //   if (data.type == uiTypes.list) {
  //     return new DList(data as UIList);
  //   } else if (data.type == uiTypes.text) {
  //     return new DText(data as UIText);
  //   }
  // }
}

export default DashboardElement;
