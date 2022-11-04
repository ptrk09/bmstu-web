CREATE TYPE listing_info 
AS (
    listing_id INTEGER,
    listing_name VARCHAR, 
    room_type VARCHAR, 
    neighbourhood VARCHAR, 
    date_from DATE,
    date_to DATE
);