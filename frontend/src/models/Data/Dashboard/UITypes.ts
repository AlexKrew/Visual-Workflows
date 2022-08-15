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

type UIGauge = UIElement & {
  label: string;
  min_value: number;
  max_value: number;
  value: number;
};

export { UICanvas, UIList, UIText, UIElement, UIGauge };
