import Vector2 from "./Vector";

class InteractUtil {
  static posToGridPos(pos: Vector2, cellSize: number) {
    const x = Math.floor(pos.x / cellSize) * cellSize;
    const y = Math.floor(pos.y / cellSize) * cellSize;
    return new Vector2(x, y);
  }
}

export default InteractUtil;
