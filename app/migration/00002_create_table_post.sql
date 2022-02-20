-- +goose Up
-- +goose StatementBegin
CREATE TABLE "post" (
  "id" SERIAL PRIMARY KEY,
  "title" varchar,
  "content" varchar,
  "author_id" int,
  "created_at" timestamp,
  "updated_at" timestamp
);

ALTER TABLE "post" ADD FOREIGN KEY ("author_id") REFERENCES "author" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "post";
-- +goose StatementEnd
