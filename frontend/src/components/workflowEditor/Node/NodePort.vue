<template>
  <div class="fill-width flex">
    <div ref="portRef" :id="portModel.id" class="circle fill-width bg-gray-400" :class="circlePosClass">
      <!-- <div
        v-if="isDragging"
        class="circle fill-width absolute drop-number bg-blue-700"
        :style="{ left: mousePosRel.x + 'px', top: mousePosRel.y + 'px' }"
      ></div> -->
    </div>
    <span class="justify-center flex-auto">{{ portModel.title }}</span>
  </div>
</template>

<script lang="ts">
import Vector2 from "@/components/util/Vector";
import GridModel from "@/models/GridModel";
import NodeConnectionModel from "@/models/Node/NodeConnectionModel";
import NodeModel from "@/models/Node/NodeModel";
import NodePortModel from "@/models/Node/NodePortModel";
import { InteractEvent } from "@interactjs/types";
import interact from "interactjs";
import { defineComponent, onMounted, ref } from "vue";

export default defineComponent({
  components: {},
  props: {
    portModel: {
      type: NodePortModel,
      required: true,
    },
  },
  setup(props) {
    let circlePosClass = props.portModel.isInput ? "" : "absolute right-0";
    const portRef = ref<HTMLInputElement>();
    const node = props.portModel.parent as NodeModel;
    const grid = props.portModel.parent?.parent as GridModel;

    onMounted(() => {
      if (!props.portModel.parent?.parent) return;
      setPortPos();
      interact(`#${props.portModel.id}`)
        .draggable({})
        .on("dragstart", onDragStart)
        .on("dragmove", onDragMove)
        .on("dragend", onDragEnd)
        .dropzone({})
        .on("drop", onDrop);
      // .on("dragenter", onDragEnter);
    });

    function setPortPos() {
      if (!grid || !node || !portRef.value) return;
      const rect: DOMRect = portRef.value.getBoundingClientRect();

      const posAbs = new Vector2(rect.x + rect.width / 2, rect.y + rect.height / 2);
      const pos = Vector2.subtract(posAbs, node.posGridCell, grid.posRel);

      props.portModel.setPos(pos);
    }

    function onDragStart(event: InteractEvent) {
      if (!grid) return;
      if (props.portModel.isInput) {
        const connection: NodeConnectionModel | undefined = grid.getConnection(undefined, props.portModel.id);
        if (connection) {
          connection.setPortIn(undefined);
          grid.setTmp(connection.id);
        }
      } else {
        (grid as GridModel).addConnection(
          new NodeConnectionModel(props.portModel, undefined, new Vector2(event.clientX, event.clientY)),
          true
        );
      }
    }

    function onDragMove(event: InteractEvent) {
      if (!grid) return;
      if (grid.tmpConnectionIndex >= 0) {
        grid.connections[grid.tmpConnectionIndex].setMousePos(new Vector2(event.clientX, event.clientY));
      }
    }

    function onDragEnd() {
      if (!props.portModel.parent?.parent) return;
      grid.resetTmp(true);
    }

    function onDrop(event: InteractEvent) {
      if (!grid) return;
      // event.target         = the Element on which it gets dropped
      // event.relatedTarget  = the Element which dropped
      if (grid.getPortByID(event.target.id)?.isInput) {
        const connection = grid.getTmpConnection();
        connection.setPortIn(grid.getPortByID(event.target.id));
        grid.resetTmp();
      } else {
        grid.resetTmp(true);
      }
    }

    return {
      circlePosClass,
      portRef,
    };
  },
});
</script>

<style>
.circle {
  width: 15px;
  height: 15px;
  -webkit-border-radius: 25px;
  -moz-border-radius: 25px;
  border-radius: 25px;
}
</style>
