create table category(
    id serial primary key,
    name text,
    description text
);

create table product(
    id serial primary key,
    category integer,
    name text
);