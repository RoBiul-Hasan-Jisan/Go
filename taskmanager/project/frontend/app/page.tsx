import Link from "next/link";

export default function LandingPage() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-center px-6">
      <div className="max-w-md text-center">
        <h1 className="text-3xl font-semibold tracking-tight text-ink">
          Tasks
        </h1>
        <p className="mt-3 text-[15px] leading-relaxed text-gray-500">
          A small, focused place to keep track of what needs doing. Create an
          account and start listing your tasks in under a minute.
        </p>
        <div className="mt-8 flex justify-center gap-3">
          <Link href="/register" className="btn-primary">
            Get started
          </Link>
          <Link href="/login" className="btn-secondary">
            Log in
          </Link>
        </div>
      </div>
    </main>
  );
}
