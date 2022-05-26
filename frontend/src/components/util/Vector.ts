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

  toString(decimal = 2, separator = ":") {
    return this.x.toFixed(decimal) + separator + this.y.toFixed(decimal);
  }
}

export { Vector2 };
