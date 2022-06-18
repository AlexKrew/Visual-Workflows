type WorkflowInfoProps = {
  id: string;
  name: string;
  status: WorkflowStatus;
}

export enum WorkflowStatus {
  ShutDown = "shutdown",
  Loading = "loading",
  Loaded = "loaded",
  Running = "running",
  Stopping = "stopping",
  Stopped = "stopped",
  ShuttingDown = "shuttingdown",
}

export class WorkflowInfo {

  public readonly id: string;
  public readonly name: string;
  public readonly status: WorkflowStatus;

  private constructor(props: WorkflowInfoProps) {
    this.id = props.id;
    this.name = props.name;
    this.status = WorkflowStatus.Loaded;
  }

  public static fromJSON(json: any): WorkflowInfoProps {
    if(json === null || json === undefined) {
      throw new Error('Missing json')
    }

    if(typeof json["id"] !== "string") {
      throw new Error("Missing id")
    }

    if(typeof json["name"] !== "string") {
      throw new Error("Missing name")
    }

    return new WorkflowInfo(json as WorkflowInfoProps)
  }

}