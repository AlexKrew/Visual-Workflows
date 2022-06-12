<template>
  <TransitionRoot as="template" :show="modelValue">
    <Dialog as="div" class="relative z-10" @close="closeModal">
      <TransitionChild as="template" enter="ease-out duration-300" enter-from="opacity-0" enter-to="opacity-100" leave="ease-in duration-200" leave-from="opacity-100" leave-to="opacity-0">
        <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" />
      </TransitionChild>
      
      <slot />

    </Dialog>
  </TransitionRoot>
</template>

<script>
import { Dialog, TransitionChild, TransitionRoot } from '@headlessui/vue'

export default {
  components: {
    Dialog,
    TransitionChild,
    TransitionRoot,
  },

  props: {
    modelValue: {
      type: Boolean,
      required: true,
    }
  },

  setup(_, ctx) {
    
    const closeModal = () => {
      ctx.emit('update:modelValue', false)
    }

    return {
      closeModal,
    }
  }
}
</script>
