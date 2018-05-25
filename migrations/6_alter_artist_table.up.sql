ALTER TABLE artist
  ADD COLUMN img_src VARCHAR,
  ADD COLUMN external_url VARCHAR,
  ADD COLUMN followers INTEGER,
  ADD COLUMN popularity INTEGER;

CREATE TABLE genre(
  id UUID UNIQUE,
  name VARCHAR NOT NULL UNIQUE
);

CREATE TABLE artist_genre(
  id UUID UNIQUE,
  artist_id UUID REFERENCES artist(id),
  genre_id UUID REFERENCES genre(id)
);

CREATE TABLE related_artist(
  id UUID UNIQUE,
  artist_id UUID REFERENCES artist(id),
  related_id UUID REFERENCES artist(id)
);

CREATE TABLE top_track(
  id UUID UNIQUE,
  artist_id UUID REFERENCES artist(id),
  name varchar NOT NULL,
  img_src varchar,
  external_url varchar
);

CREATE TABLE album(
  id UUID UNIQUE NOT NULL,
  artist_id UUID REFERENCES artist(id),
  name VARCHAR,
  img_src VARCHAR,
  externalURL VARCHAR
);
