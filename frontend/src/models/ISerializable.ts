interface ISerializable {
  fromJSON(json: JSON): void;
  toJSON(): JSON;
}

export default ISerializable;
