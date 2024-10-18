import { MokuroService } from "../model";

const API = new MokuroService();

export const useMokuroApi = () => {
  return API;
};
