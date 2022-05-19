-- create "users" table
CREATE TABLE "users" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "first_name" character varying NOT NULL, "last_name" character varying NOT NULL, "updated_at" timestamptz NOT NULL, "created_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- create "todos" table
CREATE TABLE "todos" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "title" character varying NOT NULL, "description" text NOT NULL, "status" character varying NOT NULL, "updated_at" timestamptz NOT NULL, "created_at" timestamptz NOT NULL, "user_todos" bigint NULL, PRIMARY KEY ("id"), CONSTRAINT "todos_users_todos" FOREIGN KEY ("user_todos") REFERENCES "users" ("id") ON DELETE SET NULL);
