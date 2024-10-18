export type GetProjectsResponse = {
  projects: {
    link: string;
    name: string;
  }[];
};

export type GetProjectResponse = {
  images: string[];
  ocrData: {
    data: OcrData;
    name: string;
  }[];
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
