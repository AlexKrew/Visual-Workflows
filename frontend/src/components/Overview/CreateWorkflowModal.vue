<template>
  <div class="fixed z-10 inset-0 overflow-y-auto">
    <div class="flex items-end sm:items-center justify-center min-h-full p-4 text-center sm:p-0">
      <TransitionChild as="template" enter="ease-out duration-300" enter-from="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95" enter-to="opacity-100 translate-y-0 sm:scale-100" leave="ease-in duration-200" leave-from="opacity-100 translate-y-0 sm:scale-100" leave-to="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95">
        <DialogPanel class="relative bg-gray-200 rounded-lg px-4 pt-5 pb-4 text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:max-w-lg sm:w-full sm:p-6">
          <div>
            <div class="mx-auto flex items-center justify-center h-12 w-12 rounded-full bg-white">
              <PlusCircleIcon class="h-6 w-6" aria-hidden="true" />
            </div>
            <div class="mt-5 text-center sm:mt-5">
              <DialogTitle as="h3" class="text-lg leading-6 font-medium text-gray-900"> Create a new Workflow </DialogTitle>
              <div class="mt-2">
                
                <div class="mt-6 grid grid-cols-1 gap-y-6 gap-x-4 sm:grid-cols-6">
                  
                  <div class="sm:col-span-3">
                    <label for="name" class="block text-sm font-medium text-gray-700">Name*</label>
                    <div class="mt-1">
                      <input
                        v-model="name"
                        type="text"
                        name="name"
                        id="name"
                        autocomplete="workflow-name"
                        class="px-1 py-1 shadow-sm focus:ring-blue-500 focus:border-blue-500 block w-full sm:text-sm border-gray-300 rounded-md"
                      />
                    </div>
                  </div>

                </div>

              </div>
            </div>
          </div>
          <div class="mt-10 sm:mt-6 sm:grid sm:grid-cols-2 sm:gap-3 sm:grid-flow-row-dense">
            <button
              :disabled="!canAction"
              type="button"
              class="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-blue-600 text-base font-medium text-white hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 sm:col-start-2 sm:text-sm"
              @click="onAction"
            >
              Create Workflow
            </button>

            <button
              type="button"
              class="mt-3 w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 sm:mt-0 sm:col-start-1 sm:text-sm"
              @click="onCancel"
              ref="cancelButtonRef"
            >
              Cancel
            </button>
          </div>
        </DialogPanel>
      </TransitionChild>
    </div>
  </div>
</template>

<script lang="ts">
import { ref, computed } from "vue";
import { DialogPanel, DialogTitle, TransitionChild } from '@headlessui/vue'
import { PlusCircleIcon } from '@heroicons/vue/outline'
import { CreateNewWorkflowProps } from './types';

export default {
  components: {
    DialogPanel,
    DialogTitle,
    TransitionChild,
    PlusCircleIcon
  },

  emits: ['action', 'cancel'],

  setup(_, ctx) {
    const name = ref('');

    const canAction = computed(() => {
      return name.value !== '';
    })

    const onAction = () => {

      if(!canAction.value) {
        return onCancel()
      }

      const payload: CreateNewWorkflowProps = {
        name: name.value,
      }
      ctx.emit('action', payload)

      name.value = ''
    }

    const onCancel = () => {
      ctx.emit('cancel');
    }

    return {
      name,
      canAction,
      onAction,
      onCancel,
    }
  }
}
</script>
