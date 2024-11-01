import { OcrData, TaskWithProjectResponse } from "./api.types";

export type TransformedProject = {
  name: string;
  id: string;
  data: {
    image: string;
    ocrData: OcrData;
    order: number;
  }[];
};

export type TransformedTaskWithProject = Omit<
  TaskWithProjectResponse["data"],
  "project"
> & { project: TransformedProject };
