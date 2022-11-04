-- --------------------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION get_profile_info_by_login(
    p_login VARCHAR
)
RETURNS TABLE (
    id INTEGER,
    name VARCHAR,
    surname VARCHAR,
    login VARCHAR,
    is_user_renter BOOLEAN,
    is_user_host BOOLEAN
)
AS $$
    BEGIN      
        RETURN QUERY 
            SELECT 
                users.id AS id,
                users.name AS name,
                users.surname AS surname,
                users.login AS login,
                users.is_user_renter AS isUserRenter,
                users.is_user_host AS isUserHost
            FROM 
                users
            WHERE 
                users.login = $1
            ;
    END;
$$ LANGUAGE PLPGSQL;

-- SELECT * from get_profile_info_by_login(p_login => 'temasarkisov');

-- --------------------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION get_detailed_listing_info_by_id(
    p_listing_id INTEGER
)
RETURNS TABLE (
    listing_description TEXT,
    host_name VARCHAR
)
AS $$
    BEGIN      
        RETURN QUERY 
            SELECT 
                listings_detailed.description AS listingDescription,
                host.name AS hostName
            FROM 
                listings
                JOIN listings_detailed ON listings_detailed.id = listings.id
                JOIN host ON host.id = listings.host_id 
            WHERE 
                listings.id = $1
            ;
    END;
$$ LANGUAGE PLPGSQL;

-- SELECT * from get_detailed_listing_info_by_id(p_listing_id => 6369);

-- --------------------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION get_listings(
    p_listing_name VARCHAR DEFAULT NULL,
    p_neighbourhood VARCHAR DEFAULT NULL,
    p_price_from FLOAT DEFAULT NULL,
    p_price_to FLOAT DEFAULT NULL,
    p_date_from TEXT DEFAULT NULL,
    p_date_to TEXT DEFAULT NULL
)
    RETURNS TABLE (
        id INTEGER,
        name VARCHAR, 
        room_type VARCHAR, 
        neighbourhood VARCHAR, 
        price FLOAT, 
        picture_url VARCHAR, 
        date_from DATE,
        date_to DATE,
        dates_range DATE[],
        available_dates DATE[],
        available_dates_number INTEGER
    )
AS $$
    DECLARE
        date_from_format DATE NOT NULL DEFAULT CURRENT_DATE - '1 year'::INTERVAL;
        date_to_format DATE NOT NULL DEFAULT date_from_format + '5 day'::INTERVAL;
        dates_range DATE[] NOT NULL DEFAULT (
            SELECT array_agg(d::date)
            FROM generate_series(
                date_from_format,
                date_to_format,
                '1 day'
            ) AS dr(d)
        );

    BEGIN
        IF (p_date_from IS NOT NULL AND p_date_to IS NOT NULL) THEN 
            date_from_format := to_date(p_date_from, 'YYYY-MM-DD');
            date_to_format := to_date(p_date_to, 'YYYY-MM-DD');
        END IF;

        IF (p_date_from IS NOT NULL AND $6 IS NULL) THEN 
            date_from_format := to_date(p_date_from, 'YYYY-MM-DD');
            date_to_format := date_from_format + '5 day'::INTERVAL;
        END IF;

        IF (p_date_from IS NULL AND p_date_to IS NOT NULL) THEN 
            date_to_format := to_date(p_date_to, 'YYYY-MM-DD');
            date_from_format := date_to_format - '5 day'::INTERVAL;
        END IF;

        dates_range := (
            SELECT array_agg(d::date)
            FROM generate_series(
                date_from_format,
                date_to_format,
                '1 day'
            ) AS dr(d)
        );

        RETURN QUERY 
            SELECT 
                listings.id AS listingId,
                listings.name AS listingName, 
                room_type.room_type AS roomType, 
                listings.neighbourhood AS listingNeighbourhood, 
                listings.price AS listingPrice,
                listings_detailed.picture_url AS pictureUrl,
                date_from_format AS dateFrom, 
                date_to_format AS dateTo,
                dates_range AS datesRange, 
                array_agg(calendar.date::DATE) AS availableDatesInRange,
                array_length(array_agg(calendar.date::DATE), 1) AS availableDatesInRangeNum
            FROM 
                listings
                JOIN listings_detailed ON listings_detailed.id = listings.id
                JOIN host ON host.id = listings.host_id 
                JOIN room_type ON room_type.id = listings.room_type_id
                JOIN calendar ON calendar.listing_id = listings.id
            WHERE 
                (p_listing_name IS NULL OR listings.name = p_listing_name) AND
                (p_neighbourhood IS NULL OR listings.neighbourhood = p_neighbourhood) AND
                (
                    (p_price_from IS NULL AND p_price_to IS NULL) OR
                    (p_price_from IS NULL AND listings.price < $4) OR 
                    (p_price_to IS NULL AND listings.price >= p_price_from) OR 
                    (listings.price BETWEEN p_price_from AND p_price_to)
                ) AND
                (calendar.date BETWEEN date_from_format AND date_to_format) AND
                (calendar.available = TRUE) AND
                (listings.minimum_nights <= array_length(dates_range, 1))
                -- (dates_range = get_available_dates_by_id_1(p_listing_id => listings.id, p_date_from => date_from_format, p_date_to => date_to_format))
            GROUP BY 
                listings.id,
                listings_detailed.picture_url,
                room_type.room_type
            ;
    END;
$$ LANGUAGE PLPGSQL;

-- SELECT * from get_listings_info(p_listing_name => 'Rooftop terrace room ,  ensuite bathroom', p_date_from => '2021-04-15', p_date_to => '2021-04-18');
-- SELECT get_listings_info(p_date_from => '2021-04-10', p_date_to => '2021-04-20');
-- SELECT * FROM get_listings_info(p_price_from => 100.0, p_date_from => '2021-04-10', p_date_to => '2021-04-20');
-- SELECT * FROM get_listings_info(p_price_from => 50.0, p_price_to => 100.0, p_date_from => '2021-04-10', p_date_to => '2021-04-20');
-- SELECT * FROM get_listings_info(p_price_from => 1000.0);

-- --------------------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION get_available_dates_by_id(
    p_listing_id INTEGER DEFAULT NULL,
    p_date_from TEXT DEFAULT NULL,
    p_date_to TEXT DEFAULT NULL
)
RETURNS TABLE (
    listing_id INTEGER,
    available_dates DATE[]
)
AS $$
    BEGIN
        RETURN QUERY
            SELECT 
                listings.id,
                array_agg(calendar.date::DATE)            
            FROM 
                listings
                JOIN calendar ON calendar.listing_id = listings.id
            WHERE
                ($1 IS NULL OR listings.id = $1) AND
                (
                    ($2 IS NULL AND $3 IS NULL) OR 
                    EXISTS (
                        SELECT 1
                        FROM generate_series(to_date($2, 'YYYY-MM-DD'), to_date($3, 'YYYY-MM-DD'), '1 day'::INTERVAL) AS date_range(idx)
                        WHERE date_range.idx = calendar.date 
                    )
                ) AND
                (calendar.available = TRUE)
            GROUP BY
                listings.id
            ;
    END;
$$ LANGUAGE PLPGSQL;

-- SELECT get_available_dates_by_id(p_listing_id => 167183, p_date_from => '2021-04-10', p_date_to => '2021-04-20');

-- --------------------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION is_login_unique(p_login VARCHAR)
RETURNS BOOLEAN
AS $$
    BEGIN
        IF EXISTS (SELECT FROM users WHERE users.login = p_login) THEN
            RETURN FALSE;
        ELSE
            RETURN TRUE;
        END IF;
    END;
$$ LANGUAGE PLPGSQL;

-- SELECT is_login_unique(login => 'temasarkisov');

-- --------------------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION is_signin_success(
    p_login VARCHAR,
    p_password VARCHAR
)
RETURNS BOOLEAN
AS $$
    BEGIN
        IF EXISTS (SELECT FROM users WHERE users.login = p_login AND users.password = p_password) THEN
            RETURN TRUE;
        ELSE
            RETURN FALSE;
        END IF;
    END;
$$ LANGUAGE PLPGSQL;

-- SELECT is_signin_success(p_login => 'temasarkisov', p_password => 'qwerty');

-- --------------------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION is_user_renter(
    p_id INTEGER
)
RETURNS BOOLEAN
AS $$
    BEGIN
        RETURN (SELECT is_user_renter FROM users WHERE users.id = p_id);
    END;
$$ LANGUAGE PLPGSQL;

-- SELECT is_user_renter(p_id => 1);

-- --------------------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION is_user_host(
    p_id INTEGER
)
RETURNS BOOLEAN
AS $$
    BEGIN
        RETURN (SELECT is_user_host FROM users WHERE users.id = p_id);
    END;
$$ LANGUAGE PLPGSQL;

-- SELECT is_user_host(p_id => 1);