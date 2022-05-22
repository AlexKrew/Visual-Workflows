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

  toString(decimal: number) {
    return this.x.toFixed(decimal) + ":" + this.y.toFixed(decimal);
  }
}

export { Vector2 }