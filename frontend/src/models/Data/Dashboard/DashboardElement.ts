import { UIElementType, UIListType } from "./UITypes";

class DashboardElement {
  data: UIElementType;
  children: DashboardElement[] = [];

  constructor(data: UIElementType) {
    this.data = data;
    if (this.data.children) {
      this.data.children.forEach((child) => this.children.push(new DashboardElement(child)));
    }
  }
}

export default DashboardElement;
