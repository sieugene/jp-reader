import { OcrData, ProjectResponse } from "@/shared/types/api";
import { useEffect, useState } from "react";
import { useApi } from "./useApi";

export type FormattedProject = {
  imageSrc: string;
  data: OcrData;
  id: string;
};

const FILES_BASE_URL = "http://127.0.0.1:5001";

export const useProjects = () => {
  const [projects, setProjects] = useState<FormattedProject[]>([]);
  const api = useApi();
  useEffect(() => {
    (async () => {
      const response =
        (await api.projects.projectsList()) as unknown as ProjectResponse;
      const readerData = response.data.reduce((project, current) => {
        const mappedData = current.Images?.map((imageUrl) => {
          const fileName = imageUrl
            .split("/projects/new/images/")?.[1]
            .split(".")?.[0];

          const ocrData = (
            current.OcrData as { data: OcrData; name: string }[]
          ).find((d) => d.name.includes(fileName))?.data as OcrData;

          const formatted: FormattedProject = {
            data: ocrData,
            imageSrc: `${FILES_BASE_URL}${imageUrl}`,
            id: current.ID || "",
          };

          return formatted;
        });

        if (mappedData?.length) {
          project = [...project, ...mappedData];
        }
        return project;
      }, [] as FormattedProject[]);

      setProjects(readerData);
    })();
  }, []);
  return projects;
};
