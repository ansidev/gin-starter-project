-- +goose Up
-- +goose StatementBegin
CREATE TABLE "author" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar,
  "created_at" timestamp,
  "updated_at" timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "author";
-- +goose StatementEnd
