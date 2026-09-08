import axios, { AxiosError } from "axios";
import type {
  ApiResponse,
  AuthPayload,
  CreateTaskInput,
  Task,
  UpdateTaskInput,
} from "@/types";

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/api";

const client = axios.create({
  baseURL: API_URL,
  headers: { "Content-Type": "application/json" },
});

// Attach the stored JWT (if any) to every outgoing request.
client.interceptors.request.use((config) => {
  if (typeof window !== "undefined") {
    const token = localStorage.getItem("token");
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
  }
  return config;
});

// Normalize errors so callers can always read a friendly message.
function extractErrorMessage(error: unknown): string {
  if (error instanceof AxiosError) {
    return (
      (error.response?.data as ApiResponse<null>)?.message ||
      error.message ||
      "Something went wrong"
    );
  }
  return "Something went wrong";
}

export const api = {
  // --- Auth ---
  async register(name: string, email: string, password: string) {
    try {
      const res = await client.post<ApiResponse<AuthPayload>>("/auth/register", {
        name,
        email,
        password,
      });
      return res.data.data;
    } catch (err) {
      throw new Error(extractErrorMessage(err));
    }
  },

  async login(email: string, password: string) {
    try {
      const res = await client.post<ApiResponse<AuthPayload>>("/auth/login", {
        email,
        password,
      });
      return res.data.data;
    } catch (err) {
      throw new Error(extractErrorMessage(err));
    }
  },

  // --- Tasks ---
  async getTasks() {
    try {
      const res = await client.get<ApiResponse<Task[]>>("/tasks");
      return res.data.data;
    } catch (err) {
      throw new Error(extractErrorMessage(err));
    }
  },

  async createTask(input: CreateTaskInput) {
    try {
      const res = await client.post<ApiResponse<Task>>("/tasks", input);
      return res.data.data;
    } catch (err) {
      throw new Error(extractErrorMessage(err));
    }
  },

  async updateTask(id: number, input: UpdateTaskInput) {
    try {
      const res = await client.put<ApiResponse<Task>>(`/tasks/${id}`, input);
      return res.data.data;
    } catch (err) {
      throw new Error(extractErrorMessage(err));
    }
  },

  async deleteTask(id: number) {
    try {
      await client.delete<ApiResponse<null>>(`/tasks/${id}`);
    } catch (err) {
      throw new Error(extractErrorMessage(err));
    }
  },
};
