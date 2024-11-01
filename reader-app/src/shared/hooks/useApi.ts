import { useRef } from "react";
import { Api } from "../../api/Api";
import { CONFIG } from '../config';


const JP_READER_API = new Api({
  baseUrl: CONFIG.READER_API,
});

export const useApi = () => {
  const api = useRef(JP_READER_API);

  return api.current;
};
