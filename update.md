# Project Update

## Built So Far
- Go HTTP API server with Postgres connection setup.
- User registration and login endpoints.
- JWT token generation and cookie-based auth middleware.
- Auth-protected task create, update, and delete endpoints.
- Task retrieval endpoint.
- Repository layer for users and tasks.
- Database schema for users and tasks.
- Seed SQL files for schema setup.

## Current Flow
- User registers with name, email, and password.
- User logs in and receives a JWT token in an HTTP-only cookie.
- Auth middleware validates the cookie before task write operations.
- Tasks are created, updated, and deleted using repository calls against Postgres.

## What you can build over this

- Task listing by user, including filters for status, favorites, and search.
- Pagination and sorting for large task sets.
- Task detail endpoint by ID.
- Mark task as completed, favorite, or archived with dedicated endpoints.
- User profile endpoints: get current user, update profile, change password, delete account.
- Logout and token refresh flow.
- Soft delete for tasks and users instead of hard delete.
- Input validation and consistent JSON response wrappers.
- Request logging, error middleware, and better auth middleware typing.
- Swagger/OpenAPI docs and a proper README.
- Tests for handlers, repositories, and auth flow.
- Docker setup and database migrations.

## Important gaps I noticed

- TaskRepositary.go has an ownership/update flow issue: UpdateTask runs a second update by task ID only, so the modified timestamp update is not tied to the same authorization check.
- tasks.go does not return immediately after some error responses, which can lead to extra processing.
- users.go sets cookie max age using 24 * time.Now().Hour(), which is not the right way to represent one day.
- UserRepositary.go references task_user, but that table is not defined in the SQL files, so favorites are not fully implemented yet.
- There is no pagination, filtering, or task ownership-based list endpoint yet.