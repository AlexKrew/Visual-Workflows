import { Vector2 } from "./Vector";

class InteractUtil {
  static updateGridPos(pos: Vector2, cellSize: number) {
    var x = Math.floor(pos.x / cellSize) * cellSize;
    var y = Math.floor(pos.y / cellSize) * cellSize;
    return new Vector2(x, y);
  }

  static translateElem(pos: Vector2, event: any) {
    event.target.style.transform = 'translate(' + pos.x + 'px, ' + pos.y + 'px)';
  }
}


export { InteractUtil }