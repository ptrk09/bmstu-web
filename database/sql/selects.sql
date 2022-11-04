SELECT (
    listings.id, 
    pg_typeof(listings.id),
    listings.name, 
    room_type.room_type, 
    listings.neighbourhood
)
FROM (
    listings
    JOIN host ON host.id = listings.host_id 
    JOIN room_type ON room_type.id = listings.room_type_id
);