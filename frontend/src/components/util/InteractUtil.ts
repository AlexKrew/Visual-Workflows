import Vector2 from "./Vector";
import { InteractEvent } from "@interactjs/types";

class InteractUtil {
  static posToGridPos(pos: Vector2, cellSize: number) {
    const x = Math.floor(pos.x / cellSize) * cellSize;
    const y = Math.floor(pos.y / cellSize) * cellSize;
    return new Vector2(x, y);
  }

  static translateElem(pos: Vector2, event: InteractEvent) {
    event.target.style.transform = "translate(" + pos.x + "px, " + pos.y + "px)";
  }
}

export default InteractUtil;
