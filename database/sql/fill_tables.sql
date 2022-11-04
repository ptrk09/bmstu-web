COPY listings FROM '/Users/temasarkisov/MyProjects/BMSTU/BMSTU-Web/database/data/data_csv_clean/listings.csv' DELIMITER ';' CSV HEADER;

COPY calendar (listing_id, date, available) FROM '/Users/temasarkisov/MyProjects/BMSTU/BMSTU-Web/database/data/data_csv_clean/calendar.csv' DELIMITER ';' CSV HEADER;

COPY apart_types FROM '/Users/temasarkisov/MyProjects/BMSTU/BMSTU-Web/database/data/data_csv_clean/apart_types.csv' DELIMITER ';' CSV HEADER;

COPY hosts FROM '/Users/temasarkisov/MyProjects/BMSTU/BMSTU-Web/database/data/data_csv_clean/hosts.csv' DELIMITER ';' CSV HEADER;