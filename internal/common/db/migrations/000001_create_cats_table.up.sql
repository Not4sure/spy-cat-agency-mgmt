
CREATE TABLE cats (
  id uuid PRIMARY KEY,
  created_at timestamp NOT NULL,
  name text NOT NULL,
  years_of_experience smallint NOT NULL,
  breed text NOT NULL,
  salary smallint NOT NULL
);
