CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	user_id integer NOT NULL
);


create table segments (
	id SERIAL PRIMARY KEY,
	slug varchar NOT NULL
);

create table user_segment(
	segment_id integer REFERENCES segments(id) ON UPDATE CASCADE ON DELETE CASCADE,
	user_id integer REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT user_segment_pk PRIMARY KEY(segment_id, user_id),
	expires timestamptz 
);

SET timezone = 'Europe/Moscow';