import { GetProjectResponse, GetProjectsResponse, OcrData } from "./types";

export class MokuroService {
  BASE_URL = "http://127.0.0.1:5001";

  public async getProjects() {
    const response = await fetch(`${this.BASE_URL}/projects`);
    const json: GetProjectsResponse = await response.json();
    return json;
  }

  public async getProject(name: string) {
    const response = await fetch(`${this.BASE_URL}/projects/${name}`);
    const json: GetProjectResponse = await response.json();

    return json.images.map((imageUrl) => {
      const fileName = imageUrl
        .split("/projects/new/images/")?.[1]
        .split(".")?.[0];
      const ocrData = json.ocrData.find((d) =>
        d.name.includes(fileName),
      )?.data as unknown as OcrData;

      return {
        imageSrc: `${this.BASE_URL}${imageUrl}`,
        data: ocrData,
      };
    });
  }
}
