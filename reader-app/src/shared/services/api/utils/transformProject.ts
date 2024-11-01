import { CONFIG } from "@/shared/config";
import { OcrData, ProjectResponse } from "../types/api.types";
import { TransformedProject } from "../types/transformers.types";

export const transformProject = (
  data: ProjectResponse["data"][0],
): TransformedProject => {
  const projectData =
    data.images?.map((image) => {
      // eslint-disable-next-line @typescript-eslint/no-unused-vars
      const [_, fileName] = image.split(`/projects/${data.name}/images/`);
      const [order] = fileName.split("_");
      const fileNameWithoutExt = fileName?.split(".")?.[0];
      const ocrData = (data.ocrData as { data: OcrData; name: string }[]).find(
        (d) => d.name.includes(fileNameWithoutExt),
      )?.data as OcrData;

      const formattedData: TransformedProject["data"][0] = {
        image: `${CONFIG.FILES_BASE_URL}/${image}`,
        ocrData,
        order: Number(order),
      };
      return formattedData;
    }) || [];

  return {
    id: data.id || "00-00",
    name: data.name || "-",
    data: projectData,
  };
};
