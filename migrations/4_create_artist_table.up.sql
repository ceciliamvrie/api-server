CREATE TABLE artist(
  id UUID UNIQUE,
  name VARCHAR NOT NULL UNIQUE
);

CREATE TABLE festival_artist(
  id UUID UNIQUE,
  festival_id UUID REFERENCES festival(id),
  artist_id UUID REFERENCES artist(id)
);

