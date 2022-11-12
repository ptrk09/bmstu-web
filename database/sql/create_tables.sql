CREATE TABLE IF NOT EXISTS apart_types (
    id SERIAL,
    apart_type VARCHAR,

    PRIMARY KEY (id)
);

-- --------------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS roles (
    id SERIAL,
    name VARCHAR,

    PRIMARY KEY (id)
);

-- --------------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS users (
    id SERIAL,
    name VARCHAR,
    login VARCHAR,
    password VARCHAR,
    role_id INTEGER,

    PRIMARY KEY (id),
    FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- --------------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS listings (
    id SERIAL,
    name VARCHAR,
    user_id INTEGER,

    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- --------------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS listings_detailed (
    id SERIAL, 
    description TEXT,
    neighbourhood VARCHAR,
    apart_type_id INTEGER,
    price FLOAT,
    minimum_nights INTEGER,

    PRIMARY KEY (id),
    FOREIGN KEY (apart_type_id) REFERENCES apart_types (id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- --------------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS listings_images (
    id SERIAL, 
    listing_id INTEGER,
    image_path VARCHAR,

    PRIMARY KEY (id),
    FOREIGN KEY (listing_id) REFERENCES listings (id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- --------------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS calendar (
    id SERIAL,
    listing_id INTEGER, 
    date DATE, 
    available BOOLEAN,

    PRIMARY KEY (id),
    FOREIGN KEY (listing_id) REFERENCES listings (id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- --------------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS bookings (
    id SERIAL,
    user_id INTEGER,
    listing_id INTEGER, 
    date DATE, 

    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (listing_id) REFERENCES listings (id) ON DELETE CASCADE ON UPDATE CASCADE
);