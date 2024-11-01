import { Reader } from "@/features";
import { useTaskProjectById } from "@/features/TasksProjects/hooks/useTaskProjectById";
import { TransformedProject } from "@/shared/services/api";
import { FC } from "react";

type Props = {
  id: TransformedProject["id"];
};
export const TaskProject: FC<Props> = ({ id }) => {
  const { data } = useTaskProjectById(id);
  if (!data) {
    return <h2>not found or not selected</h2>;
  }
  if (data.status !== "completed") {
    return <h2>project is {data.status}, only completed can be show</h2>;
  }
  return (
    <>
      <div className="mt-6 bg-base-100 shadow-lg rounded-lg p-6">
        <h2 className="text-2xl font-bold text-gray-900 mb-4">
          Project Details
        </h2>
        <p className="mb-2 text-gray-700">
          <span className="font-semibold">Project Name:</span> {data.title}
        </p>
        <p className="mb-4 text-gray-700">
          <span className="font-semibold">Project Id:</span> {data.id}
        </p>

        <div className="reader space-y-2">
          {data.project.data.map((readerData) => (
            <Reader {...readerData} key={readerData.order} />
          ))}
        </div>
      </div>
    </>
  );
};
