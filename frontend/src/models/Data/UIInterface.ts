// UI Structure Elements
interface UIElement {
  id: string;
}

interface UISingleSE extends UIElement {
  child: UIElement | UISingleSE | UIMultiSE;
}

interface UIMultiSE extends UIElement {
  children: (UIElement | UISingleSE | UIMultiSE)[];
}

// UI Elements
interface Text extends UIElement {
  is_label: boolean;
  font_size: number;
  value: string;
}
