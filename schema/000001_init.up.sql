CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password varchar(255) not null 
);

CREATE TABLE items
(
    id serial not null unique,
    name varchar(255) not null,
    cost int not null,
    stok boolean not null default false
);

CREATE TABLE cart
(
    id serial not null unique,
    user_id int references users (id) on delete cascade not null, 
    item_id int references items (id) on delete cascade not null
);

CREATE TABLE order
(
    id serial not null unique,
    cart_id int references cart (id) on delete cascade not null,mi
    adr varchar(255) not null 
);