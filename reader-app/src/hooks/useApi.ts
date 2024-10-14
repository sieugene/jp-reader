import { useRef } from "react";
import { Api } from "../api/Api";

const JP_READER_API = new Api({
  baseUrl: "http://localhost:3000/v1",
});

export const useApi = () => {
  const api = useRef(JP_READER_API);

  return api.current;
};
