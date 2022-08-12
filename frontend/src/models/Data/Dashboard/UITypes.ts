// UI Structure Elements
type UIElement = {
  id: string;
  type: string;
  children: UIElement[];
};

// UI Elements
type UICanvas = UIElement & {
  websocket: string;
};

type UIList = UIElement & {
  is_vertical: boolean;
};

type UIText = UIElement & {
  font_size: number;
  label: string;
  value: string;
};

enum uiTypes {
  "canvas" = "canvas",
  "list" = "list",
  "text" = "text",
}

export { UICanvas, UIList, UIText, uiTypes, UIElement };
