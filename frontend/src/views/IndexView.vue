<template>
  <header>
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="md:flex md:items-center md:justify-between">
        <div class="flex-1 min-w-0">
          <h2 class="text-2xl font-bold leading-7 text-gray-900 sm:text-3xl sm:truncate">Overview</h2>
        </div>
        <div class="mt-4 flex md:mt-0 md:ml-4">
          <button
            type="button"
            class="inline-flex items-center px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            Create Workflow from template
          </button>
          <button
            type="button"
            class="ml-3 inline-flex items-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            Create a new Workflow
          </button>
        </div>
      </div>
    </div>
  </header>
  <main>
    <div class="max-w-7xl mx-auto sm:px-6 lg:px-8">
      
      <!-- Search workflow -->
      <div class="pt-2">
        <label for="first-name" class="block text-sm font-medium text-gray-700 sm:mt-px sm:pt-2"> Search for workflow </label>
        <input v-model="searchWorkflow" type="text" name="search-wf" id="search-wf" placeholder="..." class="py-1 px-2 mt-2 max-w-lg block w-full shadow-sm focus:ring-blue-500 focus:border-blue-500 sm:max-w-xs sm:text-sm border-gray-300 rounded-md" />
      </div>

      <!-- Workflow list -->
      <div class="mt-6">
        <div class="my-3">
          <h5>Workflows</h5>
        </div>

        <div class="bg-white shadow overflow-hidden sm:rounded-md">
          <ul role="list" class="divide-y divide-gray-200">
            <template v-for="wf in workflows" :key="wf.id">
              <WorkflowListItem
                v-show="filterWorkflows(wf)"
                :workflow="wf"
                @clickEditor="openEditor(wf.id)"
              ></WorkflowListItem>
            </template>
          </ul>
        </div>
      </div>
    </div>
  </main>
</template>

<script lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import WorkflowListItem from "@/components/Overview/WorkflowListItem.vue"
import { workflowInstancesService } from '@/api';
import { WorkflowInfo } from '@/api/dtos/WorkflowInfo';

// const _workflows: WorkflowInfo[] = [
//   { id: 'wf1', name: 'The first workflow', status: WorkflowStatus.Loaded },
//   { id: 'wf2', name: 'Some other workflow', status: WorkflowStatus.Running }
// ]

export default {
  components: {
    WorkflowListItem
  },
  setup() {
    const router = useRouter()

    const workflows = ref<WorkflowInfo[]>()
    const searchWorkflow = ref<string>('');

    workflowInstancesService.getWorkflows()

    const fetchWorkflows = async () => {
      const wfInfos = await workflowInstancesService.getWorkflows()
      workflows.value = wfInfos;
    }

    const filterWorkflows = (workflow: WorkflowInfo) => {
      return searchWorkflow.value === '' || workflow.name.toLowerCase().includes(searchWorkflow.value.toLowerCase())
    }

    const openEditor = (workflowId: string) => {
      router.push({
        name: 'workflow-editor',
        params: { workflowId }
      });
    }

    fetchWorkflows()

    return {
      workflows,

      searchWorkflow,
      filterWorkflows,

      openEditor,
    }
  }
}
</script>
