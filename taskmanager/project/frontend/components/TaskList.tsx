"use client";

import type { Task, UpdateTaskInput } from "@/types";
import TaskItem from "./TaskItem";

interface TaskListProps {
  tasks: Task[];
  onUpdate: (id: number, input: UpdateTaskInput) => Promise<void>;
  onDelete: (id: number) => Promise<void>;
}

export default function TaskList({ tasks, onUpdate, onDelete }: TaskListProps) {
  if (tasks.length === 0) {
    return (
      <div className="card text-center text-sm text-gray-500">
        No tasks yet. Add your first one above.
      </div>
    );
  }

  return (
    <ul className="space-y-3">
      {tasks.map((task) => (
        <TaskItem key={task.id} task={task} onUpdate={onUpdate} onDelete={onDelete} />
      ))}
    </ul>
  );
}
