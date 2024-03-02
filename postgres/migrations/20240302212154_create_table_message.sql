-- +goose Up
-- +goose StatementBegin
create table message (
  id serial primary key,
  from_user varchar,
  content varchar,
  created_at timestamp not null default now(),
  updated_at timestamp
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table message;
-- +goose StatementEnd
