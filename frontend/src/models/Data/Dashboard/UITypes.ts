// UI Structure Elements
type UIElementType = {
  id: string;
  type: string;
  fields: UIListType | UITextType | UIGaugeType;
  children: UIElementType[];
};

// UI Elements
type UIListType = {
  is_vertical: boolean;
};

type UITextType = {
  font_size: number;
  label: string;
  value: string;
};

type UIGaugeType = {
  label: string;
  min_value: number;
  max_value: number;
  value: number;
};

// API
type UpdateFieldType = {
  id: string;
  field: string;
  value: string | number | boolean;
};

export { UIListType, UITextType, UIElementType, UIGaugeType, UpdateFieldType };
