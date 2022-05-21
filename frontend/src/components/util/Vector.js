class Vector2 {
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