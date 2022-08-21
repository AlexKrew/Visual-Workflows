import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import "./index.css";
import UIElement from "./components/workflowEditor/Dashboard/UIElement.vue";
import UIList from "./components/workflowEditor/Dashboard/UIList.vue";

const app = createApp(App);
app.use(router).mount("#app");
app.component("UIElement", UIElement).component("UIList", UIList);
