"use client";

import { useState } from "react";
import type { Task, UpdateTaskInput } from "@/types";

interface TaskItemProps {
  task: Task;
  onUpdate: (id: number, input: UpdateTaskInput) => Promise<void>;
  onDelete: (id: number) => Promise<void>;
}

export default function TaskItem({ task, onUpdate, onDelete }: TaskItemProps) {
  const [isEditing, setIsEditing] = useState(false);
  const [title, setTitle] = useState(task.title);
  const [description, setDescription] = useState(task.description);
  const [busy, setBusy] = useState(false);

  async function toggleCompleted() {
    setBusy(true);
    try {
      await onUpdate(task.id, { completed: !task.completed });
    } finally {
      setBusy(false);
    }
  }

  async function saveEdit() {
    if (!title.trim()) return;
    setBusy(true);
    try {
      await onUpdate(task.id, { title: title.trim(), description: description.trim() });
      setIsEditing(false);
    } finally {
      setBusy(false);
    }
  }

  function cancelEdit() {
    setTitle(task.title);
    setDescription(task.description);
    setIsEditing(false);
  }

  async function handleDelete() {
    if (!confirm("Delete this task?")) return;
    setBusy(true);
    try {
      await onDelete(task.id);
    } finally {
      setBusy(false);
    }
  }

  if (isEditing) {
    return (
      <li className="card space-y-3">
        <input
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          className="input-field"
          placeholder="Title"
        />
        <textarea
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          className="input-field resize-none"
          rows={2}
          placeholder="Description"
        />
        <div className="flex gap-2">
          <button onClick={saveEdit} disabled={busy} className="btn-primary">
            Save
          </button>
          <button onClick={cancelEdit} className="btn-secondary">
            Cancel
          </button>
        </div>
      </li>
    );
  }

  return (
    <li className="card flex items-start gap-3">
      <input
        type="checkbox"
        checked={task.completed}
        onChange={toggleCompleted}
        disabled={busy}
        className="mt-1 h-4 w-4 shrink-0 cursor-pointer accent-accent"
        aria-label={task.completed ? "Mark as not completed" : "Mark as completed"}
      />
      <div className="min-w-0 flex-1">
        <p
          className={`text-sm font-medium ${
            task.completed ? "text-gray-400 line-through" : "text-ink"
          }`}
        >
          {task.title}
        </p>
        {task.description && (
          <p
            className={`mt-1 text-sm ${
              task.completed ? "text-gray-300" : "text-gray-500"
            }`}
          >
            {task.description}
          </p>
        )}
      </div>
      <div className="flex shrink-0 gap-2">
        <button onClick={() => setIsEditing(true)} className="btn-secondary">
          Edit
        </button>
        <button onClick={handleDelete} disabled={busy} className="btn-danger">
          Delete
        </button>
      </div>
    </li>
  );
}
