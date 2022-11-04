CREATE TABLE listings (
    id SERIAL, 
    name VARCHAR, 
    host_id INTEGER,
    neighbourhood VARCHAR,
    apart_type_id INTEGER,
    price FLOAT,
    minimum_nights INTEGER,

    PRIMARY KEY (id),
    FOREIGN KEY (host_id) REFERENCES hosts (id),
    FOREIGN KEY (apart_type_id) REFERENCES apart_types (id)
);

-- --------------------------------------------------------------------------------------

CREATE TABLE calendar (
    id SERIAL,
    listing_id INTEGER, 
    date DATE, 
    available BOOLEAN,

    PRIMARY KEY (id),
    FOREIGN KEY (listing_id) REFERENCES listings (id)
);

-- --------------------------------------------------------------------------------------

CREATE TABLE apart_types (
    id INTEGER,
    apart_type VARCHAR,

    PRIMARY KEY (id)
);

-- --------------------------------------------------------------------------------------

CREATE TABLE hosts (
    id SERIAL,
    name VARCHAR,

    PRIMARY KEY (id)
);

-- --------------------------------------------------------------------------------------

CREATE TABLE users (
    id SERIAL,
    name VARCHAR,
    login VARCHAR,
    password VARCHAR,
    host_id INTEGER,

    PRIMARY KEY (id),
    FOREIGN KEY (host_id) REFERENCES hosts (id)
)
