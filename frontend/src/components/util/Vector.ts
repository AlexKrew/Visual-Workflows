class Vector2 {
  x: number;
  y: number;

  constructor(x, y) {
    this.x = x;
    this.y = y;
  }

  add(x, y) {
    this.x += x;
    this.y += y;
  }

  toString(decimal) {
    return this.x.toFixed(decimal) + ":" + this.y.toFixed(decimal);
  }
}

export { Vector2 }