// UI Structure Elements
type UIElement = {
  id: string;
};

type UISingleSE = UIElement & {
  child: UIElement | UISingleSE | UIMultiSE;
};

type UIMultiSE = UIElement & {
  children: (UIElement | UISingleSE | UIMultiSE)[];
};

// UI Elements
type UICanvas = UISingleSE & {
  websocket: string;
};

type UIList = UIMultiSE & {
  isVertical: boolean;
};
type UIText = UIElement & {
  is_label: boolean;
  font_size: number;
  value: string;
};

export { UICanvas, UIList, UIText };
