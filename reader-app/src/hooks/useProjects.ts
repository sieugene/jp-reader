import { OcrData, ProjectResponse } from "@/shared/types/api";
import { useEffect, useState } from "react";
import { useApi } from "./useApi";

export type FormattedProject = {
  name: string;
  id: string;
  data: {
    image: string;
    ocrData: OcrData;
    order: number;
  }[];
};

const FILES_BASE_URL = "http://127.0.0.1:5001";

export const useProjects = () => {
  const [projects, setProjects] = useState<FormattedProject[]>([]);
  const api = useApi();
  useEffect(() => {
    (async () => {
      const response =
        (await api.projects.getProjects()) as unknown as ProjectResponse;
      const readerData = response.data.map((data) => {
        const projectData =
          data.images?.map((image) => {
            // eslint-disable-next-line @typescript-eslint/no-unused-vars
            const [_, fileName] = image.split(`/projects/${data.name}/images/`);
            const [order] = fileName.split("_");
            const fileNameWithoutExt = fileName?.split(".")?.[0];
            const ocrData = (
              data.ocrData as { data: OcrData; name: string }[]
            ).find((d) => d.name.includes(fileNameWithoutExt))?.data as OcrData;

            const formattedData: FormattedProject["data"][0] = {
              image: `${FILES_BASE_URL}/${image}`,
              ocrData,
              order: Number(order),
            };
            return formattedData;
          }) || [];
        const project: FormattedProject = {
          id: data.id || "00-00",
          name: data.name || "-",
          data: projectData,
        };
        return project;
      });

      setProjects(readerData);
    })();
  }, []);
  return projects;
};
