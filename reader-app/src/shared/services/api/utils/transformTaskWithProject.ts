import { TaskWithProjectResponse } from "../types/api.types";
import { TransformedTaskWithProject } from "../types/transformers.types";
import { transformProject } from "./transformProject";

export const transformTaskWithProject = (
  response: TaskWithProjectResponse["data"],
): TransformedTaskWithProject => {
  const { project, ...task } = response;
  return {
    ...task,
    project: transformProject(project),
  };
};
