CREATE OR REPLACE FUNCTION add_host_by_new_user()
RETURNS TRIGGER 
AS $$
    BEGIN
        IF NEW.is_user_host = TRUE THEN 
            CALL add_host(NEW.id);
        END IF;
        RETURN NEW;
    END;
$$ LANGUAGE PLPGSQL;

CREATE TRIGGER add_host_by_new_user_trigger
AFTER INSERT
ON users
FOR ROW
EXECUTE PROCEDURE add_host_by_new_user();

-- --------------------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION delete_info_with_user()
RETURNS TRIGGER 
AS $$
    BEGIN
        DELETE 
        FROM listings
        WHERE listings.host_id = OLD.host_id;

        DELETE 
        FROM listings_detailed
        WHERE listings_detailed.id = (
            SELECT listings.id 
            FROM listings
            WHERE listings.host_id = OLD.host_id
        );

        DELETE 
        FROM calendar
        WHERE calendar.listing_id = (
            SELECT listings.id 
            FROM listings
            WHERE listings.host_id = OLD.host_id
        );
        
        DELETE 
        FROM host
        WHERE host.id = OLD.host_id;

        RETURN OLD;
    END;
$$ LANGUAGE PLPGSQL;

CREATE TRIGGER delete_info_with_user_trigger
AFTER DELETE
ON users
FOR ROW
EXECUTE PROCEDURE delete_info_with_user();

-- --------------------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION delete_details_with_listing()
RETURNS TRIGGER 
AS $$
    BEGIN
        DELETE 
        FROM listings_detailed
        WHERE listings_detailed.id = OLD.id;

        DELETE 
        FROM calendar
        WHERE calendar.listing_id = OLD.id;

        RETURN OLD;
    END;
$$ LANGUAGE PLPGSQL;

CREATE TRIGGER delete_details_with_listing_trigger
BEFORE DELETE
ON listings
FOR EACH ROW
EXECUTE PROCEDURE delete_details_with_listing();

-- --------------------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION plus_host_listings_count()
RETURNS TRIGGER 
AS $$
    BEGIN
        UPDATE host  
        SET calculated_host_listings_count = (
            SELECT host.calculated_host_listings_count + 1
            FROM host
            WHERE host.id = NEW.host_id
        )
        WHERE host.id = NEW.host_id;

        RETURN NEW;
    END;
$$ LANGUAGE PLPGSQL;

CREATE TRIGGER plus_host_listings_count_trigger
AFTER INSERT
ON listings
FOR ROW
EXECUTE PROCEDURE plus_host_listings_count();

-- --------------------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION minus_host_listings_count()
RETURNS TRIGGER 
AS $$
    BEGIN
        UPDATE host  
        SET calculated_host_listings_count = (
            CASE WHEN calculated_host_listings_count > 1 
            THEN (calculated_host_listings_count - 1) 
            ELSE 0 END
        )
        WHERE host.id = OLD.host_id;

        RETURN OLD;
    END;
$$ LANGUAGE PLPGSQL;

CREATE TRIGGER minus_host_listings_count_trigger
AFTER DELETE
ON listings
FOR ROW
EXECUTE PROCEDURE minus_host_listings_count();