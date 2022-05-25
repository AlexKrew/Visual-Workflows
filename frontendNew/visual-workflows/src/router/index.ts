import WorkflowEditorView from "@/views/WorkflowEditorView.vue";
import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "home",
    redirect: () => {
      return { path: "/workflow-editor" };
    },
  },
  {
    path: "/workflow-editor",
    name: "Workflow Editor",
    component: WorkflowEditorView,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
