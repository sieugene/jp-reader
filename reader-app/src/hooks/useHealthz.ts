import { useEffect } from "react";
import { useApi } from "./useApi";

export const useHealthz = () => {
  const api = useApi();
  useEffect(() => {
    (async () => {
      await api.healthz.healthzList();
    })();
  }, []);
};
