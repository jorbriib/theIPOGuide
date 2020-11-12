ALTER TABLE sectors
    ADD COLUMN image_url VARCHAR(255) AFTER name,
    ADD COLUMN total_ipos INT DEFAULT 0;

ALTER TABLE countries
    ADD COLUMN image VARCHAR(255) AFTER name,
    ADD COLUMN total_ipos INT DEFAULT 0;

ALTER TABLE markets
    ADD COLUMN image VARCHAR(255) AFTER name,
    ADD COLUMN total_ipos INT DEFAULT 0;