import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    redirect: () => {
      return {path: '/workflow-editor'} 
    }
  },
  {
    path: '/workflow-editor',
    name: 'workflow-editor',
    component: () => import("@/views/WorkflowEditorView")
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
