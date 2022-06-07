import { HTTP } from "./http/http_common";

export interface WorkflowInstancesService {
  getWorkflows(): Promise<any>;
  loadWorkflow(id: string): Promise<any>;
  startWorkflow(id: string): Promise<any>;
  stopWorkflow(id: string): Promise<any>;
  shutdownWorkflow(id: string): Promise<any>;
}

// TODO: add and setup axios

export class WorkflowInstancesServiceImpl implements WorkflowInstancesService {

  public async getWorkflows(): Promise<any> {
      HTTP.get('/workflows')
      .then(response => {
        console.log("GET /workflows response:", response)
      })
      .catch(err => {
        console.log("Failed to GET /workflows", err)
      })
  }

  public async loadWorkflow(id: string): Promise<any> {
      return null;
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