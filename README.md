# Precize Golang Assignment

This project contains a **CLI** built with Golang. It manages student records with functionality to add, update, delete, and retrieve student data from an SQLite database. The project uses `sqlc` for SQL query generation and `goose` for database migrations.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
  - [CLI Commands](#cli-commands)
- [Database Management](#database-management)
- [Makefile Commands](#makefile-commands)

## Installation

### 1. Clone the Repository

```bash
git clone https://github.com/Raghav1909/sat_app
cd sat_app
```

### 2. Install Dependencies

Ensure you have Golang and the following tools installed:

- **goose** (for migrations):

  ```bash
  go install github.com/pressly/goose/cmd/goose
  ```

- **sqlc** (for query generation):

  ```bash
  go install github.com/kyleconroy/sqlc/cmd/sqlc
  ```

### 3. Set up the Database

Run the following command to create and apply all migrations:

```bash
make reset-db
```

This command will delete the current database (if exists), create a new one, and run the migrations.

### 4. Build the Project

You can build the CLI application using:

```bash
make build
```

This will generate the `sat-cli` binary.

## Usage

### CLI Commands

1. **Add a student**:

   ```bash
   ./sat-cli create
   ```

2. **List all students**:

   ```bash
   ./sat-cli list
   ```

3. **Update a student's SAT score**:

   ```bash
   ./sat-cli update --name="John Doe"
   ```

4. **Delete a student**:

   ```bash
   ./sat-cli delete --name="John Doe"
   ```
## Database Management

The project uses SQLite as the database, and `goose` is used for managing database migrations. Migrations are stored in the `db/migrations` directory.

To create a new migration:

```bash
make create-migration
```

To apply migrations:

```bash
make migrate-up
```

## Makefile Commands

The **Makefile** provides several commands to automate tasks. Here are the available commands:

1. **`make sqlc`**: Generate Go code from SQL queries using `sqlc`.
   
2. **`make migrate-up`**: Apply all pending migrations using `goose`.

3. **`make create-db`**: Create the SQLite database file (`db.sqlite3`).

4. **`make delete-db`**: Delete the SQLite database file.

5. **`make reset-db`**: Reset the database by deleting the current file, creating a new one, and applying all migrations.

6. **`make build`**: Build the CLI application and output the binary (`sat-cli`).
