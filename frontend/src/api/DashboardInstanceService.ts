import { WorkflowInfo } from "./dtos/WorkflowInfo";
import { HTTP } from "./http/http_common";

export interface DashboardInstanceService {
  getDashboard(id: string): Promise<any>;
}

export class DashboardInstanceServiceImpl implements DashboardInstanceService {
  public async getDashboard(id: string): Promise<any> {
    try {
      const response = await HTTP.get("/dashboard/" + id);
      console.log(response.data);
      return response;
    } catch (e) {
      console.log("Failed to GET Workflow", e);
      return [];
    }
  }
}
