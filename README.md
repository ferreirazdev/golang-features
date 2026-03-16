# golang-features

> A hands-on study guide for building production-style Go services with **TDD**, **Clean Architecture**, and **SQL**.

---

## What is this?

This repository is a **structured learning project** for Go backend development. Instead of scattered snippets, you get a single, end-to-end example: a small **Task API** built phase by phase, with practices that matter in startup and senior interviews.

Think of it as a **study guide you can run**: each phase adds one layer (domain → persistence → HTTP → wiring), so you see how Clean Architecture and TDD fit together in real code.

---

## What you'll practice

| Area | What you'll do |
|------|----------------|
| **TDD** | Write tests first for use cases, repository, and HTTP handlers; keep a fast, reliable test suite. |
| **Clean Architecture** | Keep domain and use cases independent of frameworks; plug in SQLite and HTTP as adapters. |
| **SQL** | Design a simple schema, use `database/sql` with SQLite, and test against an in-memory DB. |
| **APIs** | Implement `POST/GET/DELETE` for tasks and map cleanly between HTTP and your domain. |

---

## How to use this guide

1. **Read the roadmap**  
   Open [`taskapi/roadmap.md`](taskapi/roadmap.md) and follow the phases in order. Each phase has a clear goal and deliverable.

2. **Code along**  
   Implement each phase yourself before peeking at any reference code. The roadmap tells you *what* to build; you decide *how*.

3. **Run and test**  
   By the end you’ll have a runnable service and a test suite you can walk through in an interview.

4. **Extend (optional)**  
   Use the “Optional Phase 6” ideas (validation, middleware, pagination) to go deeper.

---

## Repository layout

```
golang-features/
├── README.md           ← You are here
└── taskapi/
    ├── roadmap.md      ← Your study roadmap (phases 1–6)
    └── main.go         ← Entrypoint; you’ll grow this as you wire layers
```

As you progress, the roadmap suggests an `internal/` layout (domain, usecase, infra) so the final structure stays clean and interview-ready.

---

## Quick start (after you’ve built it)

```bash
cd taskapi
go run .
# Service runs at http://localhost:8080
```

Then try:

- `POST /tasks` — create a task (JSON body with `title`, optional `description`)
- `GET /tasks` — list all tasks
- `DELETE /tasks/:id` — delete a task by ID

---

*Happy studying — and good luck with the interviews.*
