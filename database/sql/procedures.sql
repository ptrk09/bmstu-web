CREATE OR REPLACE PROCEDURE create_listing(
    p_user_id INTEGER,
    p_name VARCHAR,
    p_neighbourhood VARCHAR,
    p_room_type_id INTEGER,
    p_price FLOAT,
    p_minimum_nights INTEGER DEFAULT NULL,
    p_date_from TEXT DEFAULT NULL,
    p_date_to TEXT DEFAULT NULL, 
    p_description TEXT DEFAULT NULL
)
AS $$
    DECLARE
        date_from_format DATE DEFAULT NULL;
        date_to_format DATE DEFAULT NULL;
        new_listing_id INTEGER DEFAULT NULL;
        host_id_by_user INTEGER DEFAULT NULL;
    BEGIN
        IF (p_date_from IS NOT NULL) THEN 
            date_from_format := to_date(p_date_from, 'YYYY-MM-DD');
        END IF;

        IF (p_date_to IS NOT NULL) THEN 
            date_to_format := to_date(p_date_to, 'YYYY-MM-DD');
        END IF;

        SELECT listings.id + 1
        INTO new_listing_id
        FROM listings
        ORDER BY listings.id DESC 
        LIMIT 1;

        SELECT users.host_id
        INTO host_id_by_user
        FROM users
        WHERE users.id = p_user_id;

        INSERT INTO listings (id, name, host_id, neighbourhood, room_type_id, price, minimum_nights)
        VALUES (new_listing_id, p_name, host_id_by_user, p_neighbourhood, p_room_type_id, p_price, p_minimum_nights);

        INSERT INTO listings_detailed (id, description)
        VALUES (new_listing_id, p_description);

        IF (date_from_format IS NOT NULL AND date_to_format IS NOT NULL) THEN 
            INSERT INTO calendar(listing_id, date, available)
            SELECT new_listing_id, d, TRUE         
            FROM generate_series(date_from_format, date_to_format, '1day') g(d);       
        END IF;
    END;
$$ LANGUAGE PLPGSQL;

-- CALL create_listing(p_user_id => 1, p_name => 'trap house', p_neighbourhood => 'Barcelona', p_room_type_id => 1, p_price => 100, );

-- --------------------------------------------------------------------------------------

CREATE OR REPLACE PROCEDURE add_host(
    p_user_id INTEGER
)
AS $$
    DECLARE
        new_host_id INTEGER DEFAULT NULL;
        name_by_user_id VARCHAR DEFAULT NULL;
    BEGIN
        SELECT host.id + 1
        INTO new_host_id
        FROM host
        ORDER BY host.id DESC 
        LIMIT 1;

        SELECT users.name
        INTO name_by_user_id
        FROM users
        WHERE users.id = p_user_id;

        RAISE NOTICE 'Host will be created with id = %', new_host_id;
        RAISE NOTICE 'Host will be created with name = %', name_by_user_id;

        INSERT INTO host 
        VALUES (new_host_id, name_by_user_id, 0);

        UPDATE users
        SET host_id = new_host_id
        WHERE users.id = p_user_id;
    END;
$$ LANGUAGE PLPGSQL;

-- CALL add_host(1);

-- --------------------------------------------------------------------------------------