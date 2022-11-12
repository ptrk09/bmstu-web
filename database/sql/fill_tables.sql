COPY apart_types FROM '/Users/temasarkisov/MyProjects/BMSTU/BMSTU-Web/database/data/data_csv_clean/apart_types.csv' DELIMITER ';' CSV HEADER;

COPY roles FROM '/Users/temasarkisov/MyProjects/BMSTU/bmstu-web/database/data/data_csv_clean/roles.csv' DELIMITER ';' CSV HEADER;

COPY users FROM '/Users/temasarkisov/MyProjects/BMSTU/bmstu-web/database/data/data_csv_clean/users.csv' DELIMITER ';' CSV HEADER;

COPY listings FROM '/Users/temasarkisov/MyProjects/BMSTU/BMSTU-Web/database/data/data_csv_clean/listings.csv' DELIMITER ';' CSV HEADER;

COPY listings_detailed (listing_id, description, neighbourhood, apart_type_id, price, minimum_nights) FROM '/Users/temasarkisov/MyProjects/BMSTU/bmstu-web/database/data/data_csv_clean/listings_detailed.csv' DELIMITER ';' CSV HEADER;

COPY listings_images FROM '/Users/temasarkisov/MyProjects/BMSTU/bmstu-web/database/data/data_csv_clean/listings_images.csv' DELIMITER ';' CSV HEADER;

COPY calendar (listing_id, date, available) FROM '/Users/temasarkisov/MyProjects/BMSTU/BMSTU-Web/database/data/data_csv_clean/calendar.csv' DELIMITER ';' CSV HEADER;
