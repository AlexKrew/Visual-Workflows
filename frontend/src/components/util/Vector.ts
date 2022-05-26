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

  static add(vector1: Vector2, vector2: Vector2): Vector2 {
    return new Vector2(vector1.x + vector2.x, vector1.y + vector2.y);
  }

  static subtract(vector1: Vector2, vector2: Vector2): Vector2 {
    return new Vector2(vector1.x - vector2.x, vector1.y - vector2.y);
  }

  static lerp(vector1: Vector2, vector2: Vector2, percent: number): Vector2 {
    const x = vector1.x * (1 - percent) + vector2.x * percent;
    const y = vector1.y * (1 - percent) + vector2.y * percent;
    return new Vector2(x, y);
  }

  toString(decimal = 2, separator = ", ") {
    return this.x.toFixed(decimal) + separator + this.y.toFixed(decimal);
  }
}

export default Vector2;
