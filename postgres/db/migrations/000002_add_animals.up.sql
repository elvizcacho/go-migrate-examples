create extension if not exists "uuid-ossp";

CREATE TABLE IF NOT EXISTS animals(
    id uuid PRIMARY KEY default uuid_generate_v4(),
    name VARCHAR (50) NOT NULL
);
