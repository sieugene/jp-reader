import { useHealthz } from "@/hooks/useHealthz";
import { FormattedProject, useProjects } from "@/hooks/useProjects";
import { useState } from "react";
import { Reader } from "./features";

function App() {
  useHealthz();
  const [selectedProject, setSelectedProject] = useState<
    FormattedProject[] | null
  >(null);
  const projects = useProjects();

  const selectProject = async (projectId: string) => {
    const project = projects.filter((project) => project.id === projectId);
    if (project) {
      setSelectedProject(project);
    }
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
              selectProject(project.id || "");
            }}
          >
            {project.id}
          </p>
        </div>
      ))}
      {selectedProject ? (
        <>
          {selectedProject?.map((project, index) => {
            return (
              <Reader
                data={project.data}
                imageSrc={project.imageSrc}
                id={project.id}
                key={index}
              />
            );
          })}
        </>
      ) : (
        <h2> project not selected</h2>
      )}
    </div>
  );
}

export default App;
