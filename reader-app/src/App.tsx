import { Reader } from "@/features";
import { useHealthz } from "@/hooks/useHealthz";
import { FormattedProject, useProjects } from "@/hooks/useProjects";
import { useState } from "react";
import BaseLayout from "./layouts/BaseLayout";

function App() {
  useHealthz();
  const [selectedProject, setSelectedProject] =
    useState<FormattedProject | null>(null);
  const projects = useProjects();

  const selectProject = async (projectId: string) => {
    const project = projects.find((project) => project.id === projectId);
    if (project) {
      setSelectedProject(project);
    }
  };

  return (
    <BaseLayout>
      <div className="projects p-6 max-w-4xl mx-auto">
        <div className="projects-list border border-base-300 rounded-lg shadow-md bg-base-100 p-4 transition-all duration-300 hover:shadow-lg">
          {projects.map((project) => (
            <div
              key={project.id}
              className="cursor-pointer hover:bg-base-200 p-3 rounded-lg transition duration-200"
              onClick={() => selectProject(project.id || "")}
            >
              <p className="font-bold text-lg text-gray-800">{project.name}</p>
            </div>
          ))}
        </div>

        {selectedProject ? (
          <div className="mt-6 bg-base-100 shadow-lg rounded-lg p-6">
            <h2 className="text-2xl font-bold text-gray-900 mb-4">
              Project Details
            </h2>
            <p className="mb-2 text-gray-700">
              <span className="font-semibold">Project Name:</span>{" "}
              {selectedProject.name}
            </p>
            <p className="mb-4 text-gray-700">
              <span className="font-semibold">Project Id:</span>{" "}
              {selectedProject.id}
            </p>

            <div className="reader space-y-2">
              {selectedProject.data.length > 0 ? (
                selectedProject.data.map((readerData) => (
                  <Reader {...readerData} key={readerData.order} />
                ))
              ) : (
                <p className="text-gray-500">
                  No data available for this project.
                </p>
              )}
            </div>
          </div>
        ) : (
          <h2 className="mt-6 text-gray-500 text-lg">Project not selected</h2>
        )}
      </div>
    </BaseLayout>
  );
}

export default App;
