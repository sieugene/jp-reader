import { useHealthz } from "@/hooks/useHealthz";
import { useEffect, useState } from "react";
import {
  GetProjectsResponse,
  OcrData,
  useMokuroApi,
} from "./shared/services/mokuro";
import { Reader } from "./features";

function App() {
  const [selectedProject, setSelectedProject] = useState<
    | {
        imageSrc: string;
        data: OcrData;
      }[]
    | null
  >(null);
  const [projects, setProjects] = useState<GetProjectsResponse["projects"]>([]);
  useHealthz();
  const api = useMokuroApi();

  useEffect(() => {
    (async () => {
      const response = await api.getProjects();
      setProjects(response.projects);
    })();
  }, []);

  const fetchProject = async (project: string) => {
    const response = await api.getProject(project);

    setSelectedProject(response);
  };

  return (
    <div
      style={{
        maxWidth: "1280px",
        margin: "0 auto",
        display: "flex",
        justifyContent: "center",
      }}
    >
      {projects.map((project, index) => (
        <div key={index}>
          <p
            style={{ fontWeight: "bold" }}
            onClick={() => {
              fetchProject(project.name);
            }}
          >
            {project.name}
          </p>
        </div>
      ))}
      {selectedProject?.map((project, index) => {
        return (
          <Reader data={project.data} imageSrc={project.imageSrc} key={index} />
        );
      })}
    </div>
  );
}

export default App;
