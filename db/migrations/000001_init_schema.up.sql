CREATE TYPE "TypeBalance" AS ENUM (
  'debit',
  'kredit'
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password_hash" varchar NOT NULL,
  "avatar_file_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_balance" (
  "id" bigserial PRIMARY KEY,
  "user_id" int,
  "balance" bigint NOT NULL,
  "balance_achieve" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_balance_history" (
  "id" bigserial PRIMARY KEY,
  "user_balanceId" bigint NOT NULL,
  "balance_before" int NOT NULL,
  "balance_after" int DEFAULT 0,
  "activity" int,
  "type" "TypeBalance" NOT NULL,
  "ip" varchar,
  "location" int,
  "user_agent" varchar,
  "author" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("username");

COMMENT ON COLUMN "user_balance"."balance" IS 'can be negative or positive';

ALTER TABLE "user_balance" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_balance_history" ADD FOREIGN KEY ("user_balanceId") REFERENCES "user_balance" ("id");

