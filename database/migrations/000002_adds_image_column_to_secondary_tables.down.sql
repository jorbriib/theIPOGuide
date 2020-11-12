ALTER TABLE sectors
    DROP COLUMN image_url,
    DROP COLUMN total_ipos;

ALTER TABLE countries
    DROP COLUMN image,
    DROP COLUMN total_ipos;

ALTER TABLE markets
    DROP COLUMN image,
    DROP COLUMN total_ipos;