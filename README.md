# Simple Go RSS Agregattor

## Environment file setup

```sh
PORT=8080
HOSTNAME=localhost
DBUSERNAME=postgres
DBPASSWORD=
DBPORT=5432
DBNAME=rssagg
DBURL=postgres://${DBUSERNAME}:${DBPASSWORD}@${HOSTNAME}:${DBPORT}/${DBNAME}
```

## The users DataBase Schematics

### Migration tool

Goose is a command line database migration tool written in Go. It runs migrations from the same SQL files that SQLC uses, making the pair of tools a perfect fit.
I recommend [installing](https://github.com/pressly/goose#install) it using `go install`:
Of course if you do not have `go` I recommend to go (did you get what I did there), to the [official go Webpage](https://go.dev/) and install it.

```sh
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Run `goose -version` to make sure it's installed correctly. if is not then you should:

```sh
# For zsh
echo >> "$HOME/.zshrc" "export PATH=\"$PATH:$HOME/go/bin\""
# For bash
echo >> "$HOME/.bashrc" "export PATH=\"$PATH:$HOME/go/bin\""
```

I recommend creating an `sql` directory in the root of your project, and in there creating a `schema` directory.

```sh
mkdir -p sql/schema
```

A "migration" is a SQL file that describes a change to your database schema. For now, we need our first migration to create a `users` table. The simplest format for these files is:

```sh
touch sql/schema/number_name.sql
```

For example this is the file `001_users.sql` in the project:

```sql
-- +goose Up
CREATE TABLE users (...);

-- +goose Down
DROP TABLE users;
```

The `-- +goose Up` and `-- +goose Down` comments are required. They tell Goose how to run the migration. An "up" migration moves your database from its old state to a new state. A "down" migration moves your database from its new state back to its old state.

Run the migration

Move/`cd` into the `sql/schema` directory and run:

```sh
goose postgres $DBURL up
```

Where `DBURL` is an environment variable exported used as the connection string for your database. The format is:

```
protocol://username:password@host:port/database
```

Run your migration! Make sure it works, I for example used [PGAdmin](https://github.com/pgadmin-org/pgadmin4) to find my newly created `users` table. Of Course you should also have I running instance of a database, in the test case I used `postgres`.
