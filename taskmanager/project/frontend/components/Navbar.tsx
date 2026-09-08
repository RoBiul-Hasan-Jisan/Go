"use client";

import { useRouter } from "next/navigation";
import { auth } from "@/lib/auth";

export default function Navbar({ userName }: { userName?: string }) {
  const router = useRouter();

  function handleLogout() {
    auth.clearSession();
    router.push("/login");
  }

  return (
    <header className="border-b border-gray-200 bg-white">
      <div className="mx-auto flex max-w-3xl items-center justify-between px-6 py-4">
        <span className="text-[15px] font-semibold tracking-tight text-ink">
          Tasks
        </span>
        <div className="flex items-center gap-4">
          {userName && (
            <span className="text-sm text-gray-500">{userName}</span>
          )}
          <button onClick={handleLogout} className="btn-secondary">
            Log out
          </button>
        </div>
      </div>
    </header>
  );
}
