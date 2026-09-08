"use client";

import { useState, FormEvent } from "react";
import type { CreateTaskInput } from "@/types";

interface TaskFormProps {
  onCreate: (input: CreateTaskInput) => Promise<void>;
}

export default function TaskForm({ onCreate }: TaskFormProps) {
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [submitting, setSubmitting] = useState(false);
  const [error, setError] = useState("");

  async function handleSubmit(e: FormEvent) {
    e.preventDefault();
    if (!title.trim()) {
      setError("Give the task a title");
      return;
    }
    setError("");
    setSubmitting(true);
    try {
      await onCreate({ title: title.trim(), description: description.trim() });
      setTitle("");
      setDescription("");
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to create task");
    } finally {
      setSubmitting(false);
    }
  }

  return (
    <form onSubmit={handleSubmit} className="card space-y-3">
      <div>
        <label htmlFor="title" className="mb-1 block text-sm font-medium text-ink">
          Title
        </label>
        <input
          id="title"
          type="text"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          placeholder="e.g. Renew passport"
          className="input-field"
        />
      </div>
      <div>
        <label htmlFor="description" className="mb-1 block text-sm font-medium text-ink">
          Description <span className="text-gray-400">(optional)</span>
        </label>
        <textarea
          id="description"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          placeholder="Add any details..."
          rows={2}
          className="input-field resize-none"
        />
      </div>
      {error && <p className="text-sm text-red-600">{error}</p>}
      <button type="submit" disabled={submitting} className="btn-primary">
        {submitting ? "Adding..." : "Add task"}
      </button>
    </form>
  );
}
