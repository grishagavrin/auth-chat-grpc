-- +goose Up
create table if not exists note (
    id serial primary key,
    title text not null,
    content text,
    created_at timestamp not null default now(),
    updated_at timestamp
);

-- +goose Down
drop table note;

