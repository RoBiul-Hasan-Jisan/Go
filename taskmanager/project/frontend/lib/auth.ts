import type { User } from "@/types";

const TOKEN_KEY = "token";
const USER_KEY = "user";

// Thin wrapper around localStorage so the storage keys and JSON
// parsing/serialization live in exactly one place.
export const auth = {
  saveSession(token: string, user: User) {
    localStorage.setItem(TOKEN_KEY, token);
    localStorage.setItem(USER_KEY, JSON.stringify(user));
  },

  getToken(): string | null {
    if (typeof window === "undefined") return null;
    return localStorage.getItem(TOKEN_KEY);
  },

  getUser(): User | null {
    if (typeof window === "undefined") return null;
    const raw = localStorage.getItem(USER_KEY);
    return raw ? (JSON.parse(raw) as User) : null;
  },

  isAuthenticated(): boolean {
    return !!this.getToken();
  },

  clearSession() {
    localStorage.removeItem(TOKEN_KEY);
    localStorage.removeItem(USER_KEY);
  },
};
