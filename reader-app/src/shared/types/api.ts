import { HandlersProject } from "@/api/Api";

export type ProjectResponse = {
  data: (Omit<HandlersProject, "ocrData"> & {
    ocrData: { data: OcrData; name: string }[];
  })[];
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
