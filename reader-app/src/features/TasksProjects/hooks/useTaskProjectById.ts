import { TransformedProject } from "@/shared/services/api";
import { useTasksProjects } from "./useTasksProjects";

export const useTaskProjectById = (id: TransformedProject["id"]) => {
  const { data, isLoading, error } = useTasksProjects();
  return { data: data?.find((task) => task.id === id), isLoading, error };
};
