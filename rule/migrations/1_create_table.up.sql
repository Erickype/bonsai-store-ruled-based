create table rule(
    id serial primary key,
    name text,
    "desc" text,
    salience bigint,
    "when" text,
    "then" text[]
);