import { HandlersProject, HandlersTaskWithProject } from "@/api/Api";

export type ProjectResponse = {
  data: (Omit<HandlersProject, "ocrData"> & {
    ocrData: { data: OcrData; name: string }[];
  })[];
};

export type TaskWithProjectResponse = {
  data: Omit<HandlersTaskWithProject, "project"> & {
    project: ProjectResponse["data"][0];
  };
};

export interface OcrData {
  version: string;
  img_width: number;
  img_height: number;
  blocks: Block[];
}

export interface Block {
  box: number[];
  vertical: boolean;
  font_size: number;
  lines_coords: number[][][];
  lines: string[];
}
