import { useApi } from "@/shared/hooks/useApi";
import { TaskWithProjectResponse } from "@/shared/services/api/types/api.types";

import {
  transformTaskWithProject,
  TransformedTaskWithProject,
} from "@/shared/services/api";
import useSWR from "swr";

export const useTasksProjects = () => {
  const api = useApi();
  return useSWR<TransformedTaskWithProject[]>(
    "/api/useTasksProjects",
    async () => {
      const response = await api.tasks.projectsList();
      return response.data.map((rawData) =>
        transformTaskWithProject(rawData as TaskWithProjectResponse["data"]),
      );
    },
    {
      refreshInterval: 5000
    }
  );
};
