import { useHealthz } from "@/shared/hooks/useHealthz";
import { SWRConfig } from "swr";
import { TasksProjects } from "./features/TasksProjects";
import BaseLayout from "./layouts/BaseLayout";
import { useState } from "react";
import { TaskProject } from "./entities/TaskProject/ui";

function App() {
  useHealthz();
  const [selected, setSelected] = useState("");

  return (
    <SWRConfig value={{ provider: () => new Map() }}>
      <BaseLayout>
        <div className="projects p-6 max-w-4xl mx-auto">
          <TasksProjects
            selected={selected}
            onSelect={(id) => setSelected(id || "")}
          />
          <TaskProject id={selected} />
        </div>
      </BaseLayout>
    </SWRConfig>
  );
}

export default App;
