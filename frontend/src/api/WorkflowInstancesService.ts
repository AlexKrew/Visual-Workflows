import { WorkflowInfo } from "./dtos/WorkflowInfo";
import { HTTP } from "./http/http_common";

export interface WorkflowInstancesService {
  getWorkflows(): Promise<any>;
  loadWorkflow(id: string): Promise<any>;
  updateWorkflow(id: string, json: JSON): Promise<any>;
  startWorkflow(id: string): Promise<any>;
  stopWorkflow(id: string): Promise<any>;
  shutdownWorkflow(id: string): Promise<any>;
}

export class WorkflowInstancesServiceImpl implements WorkflowInstancesService {
  // Fetch all Workflows to show in Index View
  public async getWorkflows(): Promise<WorkflowInfo[]> {
    try {
      const response = await HTTP.get("/workflows");
      console.log("GET /workflows response:", response);

      const workflowInfos: WorkflowInfo[] = [];

      const data = response.data as unknown[];
      for (const entry of data) {
        const info = WorkflowInfo.fromJSON(entry);
        workflowInfos.push(info);
      }

      return workflowInfos;
    } catch (err) {
      console.log("Failed to GET /workflows", err);
      return [];
    }
  }

  // Create new empty Workflow
  public async createWorkflow(name: string): Promise<any> {
    try {
      const response = await HTTP.post("/workflows/new", { name });
      console.log("Response", response);
    } catch (err) {
      return err;
    }
  }

  // Fetch existing Workflow Data
  public async loadWorkflow(id: string): Promise<any> {
    try {
      const response = await HTTP.get("/workflows/" + id);
      console.log("Response", response);
      return response.data["workflow"]
    } catch (e) {
      console.log("Failed to GET /workflows/" + id + "\n", e);
    }
  }

  // Update existing Workflow
  public async updateWorkflow(id: string, json: JSON): Promise<any> {
    try {
      const response = await HTTP.patch("/workflows/" + id + "/start", json);
      console.log("Response", response);
    } catch (e) {
      console.log("Failed to GET /workflows/" + id + "\n", e);
    }
  }

  public async startWorkflow(id: string): Promise<any> {
    return null;
  }

  public async stopWorkflow(id: string): Promise<any> {
    return null;
  }

  public async shutdownWorkflow(id: string): Promise<any> {
    return null;
  }
}