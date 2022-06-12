<template>
  <!-- Load -->
  <div>
    <IconButton :disabled="!canLoad" @click="$emit('clickLoad')">
      <UploadIcon class="h-5 w-5" aria-hidden="true" />
    </IconButton>
  </div>
  
  <!-- Start -->
  <div>
    <IconButton :disabled="!canStart" @click="$emit('clickStart')">
      <PlayIcon class="h-5 w-5" aria-hidden="true" />
    </IconButton>
  </div>
  
  <!-- Stop -->
  <div>
    <IconButton :disabled="!canStop" @click="$emit('clickStop')">
      <PauseIcon class="h-5 w-5" aria-hidden="true" />
    </IconButton>
  </div>

  <!-- Shutdown -->
  <div>
    <IconButton :disabled="!canShutdown" @click="$emit('clickShutdown')">
      <ArchiveIcon class="h-5 w-5" aria-hidden="true" />
    </IconButton>
  </div>
</template>

<script lang="ts">
import { PropType, computed } from 'vue'
import IconButton from "../IconButton.vue"
import { UploadIcon, PlayIcon, PauseIcon, ArchiveIcon } from "@heroicons/vue/solid"
import { WorkflowStatus } from '@/api/dtos/WorkflowInfo';

export default {
  emits: ['clickLoad', 'clickStart', 'clickStop', 'clickShutdown'],
  
  components: {
    IconButton,
    UploadIcon,
    PlayIcon,
    PauseIcon,
    ArchiveIcon,
  },
  
  props: {
    status: {
      type: String as PropType<WorkflowStatus>,
      required: true,
    }
  },

  setup(props: any) {

    const canLoad = computed(() => {
      return props.status === WorkflowStatus.ShutDown;
    });

    const canStart = computed(() => {
      return props.status === WorkflowStatus.Loaded || props.status === WorkflowStatus.Stopped;
    });

    const canStop = computed(() => {
      return props.status === WorkflowStatus.Running;
    });

    const canShutdown = computed(() => {
      return props.status === WorkflowStatus.Loaded
        || props.status === WorkflowStatus.Running
        || props.status === WorkflowStatus.Stopped;
    });

    return {
      canLoad,
      canStart,
      canStop,
      canShutdown
    }
  }
}
</script>
