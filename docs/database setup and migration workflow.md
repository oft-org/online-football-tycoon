Database Setup and Migration Workflow
This project uses PostgreSQL for data storage and golang-migrate for database migrations.

Setup steps
1. Drop the existing database (optional):

task db-drop


This will delete the existing database to start fresh.

-------------------------------------------------------------------



2. Create a new database:

task db-create


This command creates the database with the configured name.

-------------------------------------------------------------------



3. Run migrations:

Using the CLI:

go run cmd/main.go migrations


Or using the Taskfile:

task migration-up


This applies all pending migrations to the database.

-------------------------------------------------------------------



----  Why this workflow?
Ensures a clean and consistent database state.

Enables repeatable and automated environment setup.

Avoids issues from leftover schema or corrupted data.

Suitable for development, testing, and CI pipelines.