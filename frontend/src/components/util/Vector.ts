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

  negateReturn(): Vector2 {
    return new Vector2(-this.x, -this.y);
  }

  static add(...vectors: Vector2[]): Vector2 {
    let x = 0;
    let y = 0;
    vectors.forEach((vector) => {
      x += vector.x;
      y += vector.y;
    });
    return new Vector2(x, y);
  }

  static subtract(...vectors: Vector2[]): Vector2 {
    let x = vectors[0].x;
    let y = vectors[0].y;
    for (let i = 1; i < vectors.length; i++) {
      x -= vectors[i].x;
      y -= vectors[i].y;
    }
    return new Vector2(x, y);
  }

  static lerp(vector1: Vector2, vector2: Vector2, percent: number): Vector2 {
    const x = vector1.x * (1 - percent) + vector2.x * percent;
    const y = vector1.y * (1 - percent) + vector2.y * percent;
    return new Vector2(x, y);
  }

  static dist(vector1: Vector2, vector2: Vector2): number {
    return Math.sqrt(Math.pow(vector2.x - vector1.x, 2) + Math.pow(vector2.y - vector1.y, 2));
  }

  toString(decimal = 2, separator = ", ") {
    return this.x.toFixed(decimal) + separator + this.y.toFixed(decimal);
  }
}

export default Vector2;
