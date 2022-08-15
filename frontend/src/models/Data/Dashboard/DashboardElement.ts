import { UIElement } from "./UITypes";

class DashboardElement {
  data: UIElement;
  children: DashboardElement[] = [];

  constructor(data: UIElement) {
    this.data = data;
    if (this.data.children) {
      this.data.children.forEach((child) => this.children.push(new DashboardElement(child)));
    }
  }
}

export default DashboardElement;
