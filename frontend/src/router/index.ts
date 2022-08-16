import MainLayout from "@/layouts/MainLayout.vue";
import DashboardView from "@/views/DashboardView.vue";
import IndexView from "@/views/IndexView.vue";
import WorkflowEditorView from "@/views/WorkflowEditorView.vue";
import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    component: MainLayout,
    children: [
      { path: "/", redirect: { name: "home" } },
      { path: "overview", name: "home", component: IndexView },
      { path: "workflow-editor/:workflowId", name: "workflow-editor", component: WorkflowEditorView, props: true },
      { path: "dashboard/:workflowId/", name: "dashboard", component: DashboardView, props: true },
    ],
  },
  // {
  //   path: "/workflow-editor/:workflowId",
  //   name: "workflow-editor",
  //   component: WorkflowEditorView,
  //   props: true,
  // },
  // {
  //   path: "/workflow-editor",
  //   name: "workflow-editor DEBUG",
  //   component: WorkflowEditorView,
  //   props: true,
  // },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
