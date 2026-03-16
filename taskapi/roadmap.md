### Task API Roadmap

This roadmap describes the phases to build a small Task API service in Go using **TDD**, **Clean Architecture**, and **SQLite**.

---

### Phase 1 – Domain modeling (no DB, no HTTP yet)

**Goal**: Define core domain and use cases with pure Go + TDD.

- **1.1 Define entities**
  - Create `internal/domain/task.go`:
    - `type Task struct { ID int64; Title string; Description string; CreatedAt time.Time; Done bool }` (or similar).
  - Keep it framework-agnostic (no HTTP, no SQL tags here).

- **1.2 Define repository interfaces**
  - In `internal/domain` or `internal/domain/ports`:
    - `type TaskRepository interface { Create(ctx, task) (Task, error); List(ctx) ([]Task, error); Delete(ctx, id int64) error }`.

- **1.3 Define use cases (application layer)**
  - In `internal/usecase` or `internal/application`:
    - `CreateTaskUseCase` with method `Execute(ctx, input) (Task, error)`.
    - `ListTasksUseCase` with method `Execute(ctx) ([]Task, error)`.
    - `DeleteTaskUseCase` with method `Execute(ctx, id int64) error`.

- **1.4 Write TDD tests for use cases with mocks/fakes**
  - Use `testing` (and optionally `testify`) with fake implementations of `TaskRepository`.
  - Tests should cover:
    - Create: fails if title is empty, returns task with ID > 0, calls repo `Create`.
    - List: returns tasks from repo.
    - Delete: calls repo `Delete`, handles a “not found” error type.

**Deliverable**: Green tests for domain and use cases; no infrastructure.

---

### Phase 2 – Infrastructure: SQLite repository

**Goal**: Implement `TaskRepository` for SQLite under `internal/infra/persistence`, still driven by tests.

- **2.1 Choose DB library**
  - Use standard `database/sql` + `_ "github.com/mattn/go-sqlite3"` driver.

- **2.2 Design SQL schema**
  - Table `tasks`:
    - `id INTEGER PRIMARY KEY AUTOINCREMENT`
    - `title TEXT NOT NULL`
    - `description TEXT`
    - `created_at DATETIME NOT NULL`
    - `done BOOLEAN NOT NULL DEFAULT 0`

- **2.3 Migration/init**
  - On startup, run `CREATE TABLE IF NOT EXISTS tasks (...)`.
  - In interviews, mention that in production you would use real migrations.

- **2.4 Implement `SQLiteTaskRepository`**
  - In `internal/infra/persistence/sqlite_task_repository.go`:
    - `Create`: `INSERT INTO tasks (...)` and get `last_insert_rowid()`.
    - `List`: `SELECT ... FROM tasks ORDER BY id`.
    - `Delete`: `DELETE FROM tasks WHERE id = ?`; if `RowsAffected == 0`, return `ErrNotFound`.

- **2.5 Repository-level tests**
  - Use in-memory DB: `sql.Open("sqlite3", ":memory:")`.
  - Run schema creation in `TestMain` or per-test setup.
  - Tests:
    - `Create` inserts a row; ID > 0.
    - `List` returns what’s in DB.
    - `Delete` removes the row; a second delete on the same ID returns `ErrNotFound`.

**Deliverable**: Working SQLite repository covered by tests, still without HTTP.

---

### Phase 3 – HTTP layer (Clean Architecture adapters)

**Goal**: Implement HTTP handlers as adapters around your use cases.

- **3.1 Choose minimal HTTP stack**
  - Either:
    - Standard `net/http` only, or
    - `net/http` + a small router like `chi` or `gorilla/mux`.
  - Keep this layer separate, e.g. `internal/infra/http` or `internal/interfaces/http`.

- **3.2 Define DTOs and mapping**
  - `POST /tasks` request DTO:
    - `{ "title": "string", "description": "string" }`.
  - Response DTO:
    - `{ "id": 1, "title": "...", "description": "...", "created_at": "...", "done": false }`.
  - `GET /tasks`: returns an array of the response DTO.
  - `DELETE /tasks/:id`: returns `204 No Content` on success, `404` on not found.

- **3.3 Implement handlers**
  - `CreateTaskHandler`:
    - Parse JSON body.
    - Validate (title required); return `400 Bad Request` if invalid.
    - Call `CreateTaskUseCase.Execute`.
    - On success, return `201 Created` with created task JSON.
  - `ListTasksHandler`:
    - Call `ListTasksUseCase.Execute`.
    - Return `200 OK` with array of tasks.
  - `DeleteTaskHandler`:
    - Parse `id` from URL.
    - Call `DeleteTaskUseCase.Execute`.
    - Return `204 No Content` if OK, `404 Not Found` if `ErrNotFound`.

- **3.4 TDD at handler level**
  - Use `httptest.NewRecorder()` and `http.NewRequest`.
  - Inject fake use cases into handlers so tests don’t hit the DB.
  - Tests should assert:
    - Correct status codes for valid/invalid payloads.
    - Proper JSON mapping between DTOs and domain.

**Deliverable**: HTTP handlers tested with TDD, cleanly separated from domain and infra.

---

### Phase 4 – Wiring in `main.go`

**Goal**: Compose the layers in `main.go` without adding new business logic.

- **4.1 Main composition**
  - Steps:
    - Read configuration (for now, hardcode DB path like `./taskapi.db`).
    - Open SQLite connection.
    - Initialize repository: `NewSQLiteTaskRepository(db)`.
    - Initialize use cases: `NewCreateTaskUseCase(repo)`, `NewListTasksUseCase(repo)`, `NewDeleteTaskUseCase(repo)`.
    - Initialize HTTP handlers/router, injecting use cases.
    - Start HTTP server on port `:8080`.

- **4.2 Manual testing**
  - `POST /tasks` with JSON to create tasks.
  - `GET /tasks` to verify they are persisted.
  - `DELETE /tasks/:id` to delete a task and confirm subsequent `GET` no longer returns it.

**Deliverable**: Running service with clean architecture wiring.

---

### Phase 5 – Interview storytelling: TDD and Clean Architecture

**Goal**: Make the project easy to explain in a startup interview.

- **5.1 Suggested project structure**

```text
taskapi/
  cmd/
    taskapi/
      main.go
  internal/
    domain/
      task.go
      errors.go
      task_repository.go
    usecase/
      create_task.go
      list_tasks.go
      delete_task.go
    infra/
      persistence/
        sqlite_task_repository.go
        schema.go
      http/
        router.go
        task_handlers.go
    test/
      (optional helpers/mocks)
```

- **5.2 Talking points for interviews**
  - **TDD**:
    - Started from domain use cases with fake repositories.
    - Added real SQLite repository with DB-focused tests.
    - Implemented HTTP handlers with handler-level tests.
  - **Clean Architecture**:
    - Domain and use cases do not depend on HTTP or the database.
    - Infrastructure (SQLite, HTTP) depends on the domain, not the other way around.
    - `main.go` acts as a composition root, wiring everything together.
  - **SQLite choice**:
    - Simple, file-based storage ideal for local development and coding interviews.
    - Could switch to Postgres or another DB by reimplementing only `TaskRepository`.

---

### Optional Phase 6 – Extra polish

- **Validation and error types**
  - Introduce `ErrValidation`, `ErrNotFound`, etc., in `internal/domain/errors.go`.

- **Middleware**
  - Basic logging middleware.
  - Panic recovery middleware that returns `500` instead of crashing.

- **Pagination**
  - Add query parameters `limit` and `offset` to `GET /tasks`.

- **Extra endpoint**
  - `PATCH /tasks/:id/done` to mark a task as done.

These extras are not required for the first version but are great to discuss if you have time.

