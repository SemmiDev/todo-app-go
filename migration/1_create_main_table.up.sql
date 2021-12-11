CREATE TABLE IF NOT EXISTS todo(
   id serial PRIMARY KEY,
   title VARCHAR NOT NULL,
   status boolean NOT NULL default false
);
