"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { api } from "@/lib/api";
import { auth } from "@/lib/auth";
import type { Task, CreateTaskInput, UpdateTaskInput, User } from "@/types";
import Navbar from "@/components/Navbar";
import TaskForm from "@/components/TaskForm";
import TaskList from "@/components/TaskList";

export default function DashboardPage() {
  const router = useRouter();
  const [user, setUser] = useState<User | null>(null);
  const [tasks, setTasks] = useState<Task[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    if (!auth.isAuthenticated()) {
      router.push("/login");
      return;
    }
    setUser(auth.getUser());
    loadTasks();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  async function loadTasks() {
    setLoading(true);
    setError("");
    try {
      const data = await api.getTasks();
      setTasks(data);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to load tasks");
    } finally {
      setLoading(false);
    }
  }

  async function handleCreate(input: CreateTaskInput) {
    const task = await api.createTask(input);
    setTasks((prev) => [task, ...prev]);
  }

  async function handleUpdate(id: number, input: UpdateTaskInput) {
    const updated = await api.updateTask(id, input);
    setTasks((prev) => prev.map((t) => (t.id === id ? updated : t)));
  }

  async function handleDelete(id: number) {
    await api.deleteTask(id);
    setTasks((prev) => prev.filter((t) => t.id !== id));
  }

  const remaining = tasks.filter((t) => !t.completed).length;

  return (
    <div className="min-h-screen">
      <Navbar userName={user?.name} />
      <main className="mx-auto max-w-3xl px-6 py-10">
        <div className="mb-6">
          <h1 className="text-xl font-semibold tracking-tight text-ink">
            Your tasks
          </h1>
          {!loading && (
            <p className="mt-1 text-sm text-gray-500">
              {remaining === 0
                ? "Everything is done."
                : `${remaining} task${remaining === 1 ? "" : "s"} remaining`}
            </p>
          )}
        </div>

        <div className="mb-6">
          <TaskForm onCreate={handleCreate} />
        </div>

        {error && <p className="mb-4 text-sm text-red-600">{error}</p>}

        {loading ? (
          <p className="text-sm text-gray-500">Loading tasks...</p>
        ) : (
          <TaskList tasks={tasks} onUpdate={handleUpdate} onDelete={handleDelete} />
        )}
      </main>
    </div>
  );
}
