create table Subcategory(
    id serial primary key,
    category integer,
    name text,
    description text
);

alter table product
rename column category to subcategory;

alter table Subcategory
add constraint fk_category
foreign key (category)
references category (id);

alter table product
add constraint fk_subcategory
foreign key (subcategory)
references subcategory (id);