import { TransformedTaskWithProject } from "@/shared/services/api";
import { useTasksProjects } from "../hooks/useTasksProjects";
import { FC } from "react";
import classNames from "classnames";

type Props = {
  onSelect: (id: TransformedTaskWithProject["id"]) => void;
  selected: TransformedTaskWithProject["id"];
};
export const TasksProjects: FC<Props> = ({ onSelect, selected }) => {
  const { data } = useTasksProjects();

  return (
    <div className="border border-base-300 rounded-lg shadow-md bg-base-100 p-4 transition-all duration-300 hover:shadow-lg">
      {data?.map((task) => (
        <div
          key={task.id}
          className="cursor-pointer hover:bg-base-200 p-3 rounded-lg transition duration-200"
          onClick={() => onSelect(task.id)}
        >
          <p
            className={classNames("text-lg text-gray-800", {
              "font-bold": selected === task.id,
            })}
          >
            {task.title}
          </p>
          <p className="text-xs">Status: {task.status}</p>
        </div>
      ))}
    </div>
  );
};
