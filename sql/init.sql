-- user: id, login, password, plevel
-- team: id, name, owner, date
-- player: id, fname, lname, country, dob
-- teamplayer: team_id, player_id

\c football_teams ptrk

CREATE TABLE IF NOT EXISTS users (
    id serial primary key,
    login varchar,
    password varchar,
    plevel int
);

CREATE TABLE IF NOT EXISTS teams (
    id serial primary key,
    name varchar,
    owner_id integer,
    FOREIGN KEY (owner_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS players (
    id serial primary key,
    fname varchar,
    lname varchar,
    country varchar,
    dob date
);

CREATE TABLE IF NOT EXISTS teamplayer (
    team_id integer,
    player_id integer,
    FOREIGN KEY (team_id) REFERENCES teams (id) ON DELETE CASCADE,
    FOREIGN KEY (player_id) REFERENCES players (id) ON DELETE CASCADE
);
