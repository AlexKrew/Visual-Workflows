// UI Structure Elements
type UIElement = {
  id: string;
  type: string;
  fields: UIList | UIText | UIGauge;
  children: UIElement[];
};

// UI Elements
type UIList = {
  is_vertical: boolean;
};

type UIText = {
  font_size: number;
  label: string;
  value: string;
};

type UIGauge = {
  label: string;
  min_value: number;
  max_value: number;
  value: number;
};

export { UIList, UIText, UIElement, UIGauge };
