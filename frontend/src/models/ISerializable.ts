interface ISerializable {
  fromJSON(json: JSON): ISerializable;
  toJSON(): JSON;
}

export default ISerializable;
