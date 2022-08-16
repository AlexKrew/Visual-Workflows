import { DashboardInstanceServiceImpl } from "./DashboardInstanceService";
import { WorkflowInstancesServiceImpl } from "./WorkflowInstancesService";

const workflowInstancesService = new WorkflowInstancesServiceImpl();
const dashboardInstanceService = new DashboardInstanceServiceImpl();

export { workflowInstancesService, dashboardInstanceService };
