CREATE ROLE db_cp_guest WITH LOGIN PASSWORD 'guest_password';

REVOKE ALL ON Calendar FROM db_cp_guest;

GRANT SElECT, INSERT, UPDATE ON Host TO db_cp_guest;
REVOKE DELETE ON Host FROM db_cp_guest;

REVOKE ALL ON Listings FROM db_cp_guest;

REVOKE ALL ON Listings_detailed FROM db_cp_guest;

REVOKE ALL ON Room_type FROM db_cp_guest;

GRANT SElECT, INSERT, UPDATE ON Users TO db_cp_guest;
REVOKE DELETE ON Users FROM db_cp_guest;
GRANT ALL ON SEQUENCE users_id_seq TO db_cp_guest;

-- --------------------------------------------------------------------------------------

CREATE ROLE db_cp_renter WITH LOGIN PASSWORD 'renter_password';

GRANT SELECT ON Calendar TO db_cp_renter;
REVOKE INSERT, UPDATE, DELETE ON Calendar FROM renter_user;

GRANT SELECT ON Host TO db_cp_renter;
REVOKE INSERT, UPDATE, DELETE ON Host FROM renter_user;

GRANT SELECT ON Listings TO db_cp_renter;
REVOKE INSERT, UPDATE, DELETE ON Listings FROM renter_user;

GRANT SELECT ON Listings_detailed TO db_cp_renter;
REVOKE INSERT, UPDATE, DELETE ON Listings_detailed FROM renter_user;

GRANT SELECT ON Room_type TO db_cp_renter;
REVOKE INSERT, UPDATE, DELETE ON Room_type FROM renter_user;

GRANT SELECT, INSERT ON Users TO db_cp_renter;
REVOKE UPDATE, DELETE ON Users FROM renter_user;

-- --------------------------------------------------------------------------------------

CREATE ROLE db_cp_host WITH LOGIN PASSWORD 'host_password';

GRANT SELECT, INSERT, UPDATE ON Calendar TO db_cp_host;
REVOKE DELETE ON Calendar FROM db_cp_host;

GRANT SELECT, INSERT, UPDATE ON Host TO db_cp_host;
REVOKE DELETE ON Host FROM db_cp_host;

GRANT SELECT, INSERT ON Listings TO db_cp_host;
REVOKE UPDATE, DELETE ON Listings FROM db_cp_host;

GRANT SELECT, INSERT ON Listings_detailed TO db_cp_host;
REVOKE UPDATE, DELETE ON Listings_detailed FROM db_cp_host;

GRANT SELECT ON Room_type TO db_cp_host;
REVOKE INSERT, UPDATE, DELETE ON Room_type FROM db_cp_host;

GRANT SELECT ON Users TO db_cp_host;
REVOKE INSERT, UPDATE, DELETE ON Users FROM db_cp_host;
