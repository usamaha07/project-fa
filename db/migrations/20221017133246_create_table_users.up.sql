CREATE TABLE IF NOT EXISTS users(
	id serial primary key,
    username varchar(30) not null unique,
	email varchar(50) not null unique,
	password varchar(100) not null,
    phone_number varchar(30) not null,
    age int not null,
    created_at timestamp DEFAULT current_timestamp not null,
    updated_at timestamp
);