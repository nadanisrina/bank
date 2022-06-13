CREATE TYPE "TypeBalance" AS ENUM (
  'debit',
  'kredit'
);

CREATE TABLE "user" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "passwordHash" varchar NOT NULL,
  "createdAt" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_balance" (
  "id" bigserial PRIMARY KEY,
  "userId" int,
  "balance" bigint NOT NULL,
  "balanceAchieve" bigint,
  "createdAt" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_balance_history" (
  "id" bigserial PRIMARY KEY,
  "userBalanceId" bigint NOT NULL,
  "balanceBefore" int NOT NULL,
  "balanceAfter" int DEFAULT 0,
  "activity" int,
  "type" "TypeBalance" NOT NULL,
  "ip" varchar,
  "location" int,
  "userAgent" varchar,
  "author" varchar,
  "createdAt" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "user" ("username");

COMMENT ON COLUMN "user_balance"."balance" IS 'can be negative or positive';

ALTER TABLE "user_balance" ADD FOREIGN KEY ("userId") REFERENCES "user" ("id");

ALTER TABLE "user_balance_history" ADD FOREIGN KEY ("userBalanceId") REFERENCES "user_balance" ("id");

