class Vector2 {
  x: number;
  y: number;

  constructor(x: number, y: number) {
    this.x = x;
    this.y = y;
  }

  add(x: number, y: number) {
    this.x += x;
    this.y += y;
  }

  addVector(vector: Vector2) {
    this.x += vector.x;
    this.y += vector.y;
  }

  static addVector(vector1: Vector2, vector2: Vector2): Vector2 {
    return new Vector2(vector1.x + vector2.x, vector1.y + vector2.y);
  }

  toString(decimal = 2, separator = ":") {
    return this.x.toFixed(decimal) + separator + this.y.toFixed(decimal);
  }
}

export default Vector2;
