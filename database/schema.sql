CREATE TABLE sectors
(
    uuid BINARY(16) PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE currencies
(
    uuid    BINARY(16) PRIMARY KEY,
    code    VARCHAR(3)   NOT NULL,
    name    VARCHAR(255) NOT NULL,
    display VARCHAR(32)  NOT NULL
);

CREATE TABLE countries
(
    uuid BINARY(16) PRIMARY KEY,
    code VARCHAR(3)   NOT NULL,
    name VARCHAR(255) NOT NULL

);
ALTER TABLE countries
    ADD CONSTRAINT unique_countries_code UNIQUE KEY (code);


CREATE TABLE markets
(
    uuid        BINARY(16) PRIMARY KEY,
    code        VARCHAR(16)  NOT NULL,
    name        VARCHAR(255) NOT NULL,
    currency_id BINARY(16)   NOT NULL
);
ALTER TABLE markets
    ADD CONSTRAINT unique_markets_code UNIQUE KEY (code);
ALTER TABLE markets
    ADD CONSTRAINT foreign_markets_currency FOREIGN KEY (currency_id)
        REFERENCES currencies (uuid) ON UPDATE CASCADE ON DELETE RESTRICT;

CREATE TABLE companies
(
    uuid                    BINARY(16) PRIMARY KEY,
    symbol                  VARCHAR(16)        NOT NULL,
    name                    VARCHAR(255)       NOT NULL,
    sector_id               BINARY(16)         NOT NULL,
    address                 VARCHAR(255)       NOT NULL,
    country_id              BINARY(16)         NOT NULL,
    phone                   VARCHAR(32)        NULL DEFAULT NULL,
    email                   VARCHAR(128)       NULL DEFAULT NULL,
    website                 VARCHAR(255)       NULL DEFAULT NULL,
    employees               MEDIUMINT UNSIGNED NULL DEFAULT NULL,
    description             TEXT               NOT NULL,
    founded                 YEAR               NOT NULL,
    ceo                     VARCHAR(255)       NULL DEFAULT NULL,
    fiscal_year_end         VARCHAR(64)        NULL DEFAULT NULL,
    ipo_url                 VARCHAR(255)       NULL DEFAULT NULL,
    exchange_commission_url VARCHAR(255)       NULL DEFAULT NULL,
    logo_url                VARCHAR(255)       NOT NULL
);
ALTER TABLE companies
    ADD CONSTRAINT unique_companies_symbol UNIQUE KEY (symbol);

ALTER TABLE companies
    ADD CONSTRAINT foreign_companies_sector FOREIGN KEY (sector_id)
        REFERENCES sectors (uuid) ON UPDATE CASCADE ON DELETE RESTRICT;
ALTER TABLE companies
    ADD CONSTRAINT foreign_companies_country FOREIGN KEY (country_id)
        REFERENCES countries (uuid) ON UPDATE CASCADE ON DELETE RESTRICT;

CREATE TABLE ipos
(
    uuid             BINARY(16) PRIMARY KEY,
    market_id        BINARY(16)         NOT NULL,
    company_id       BINARY(16)         NOT NULL,
    price_cents_from MEDIUMINT UNSIGNED NULL DEFAULT NULL,
    price_cents_to   MEDIUMINT UNSIGNED NULL DEFAULT NULL,
    shares           INT UNSIGNED       NULL DEFAULT NULL,
    expected_date    DATETIME           NULL DEFAULT NULL,
    created_at       TIMESTAMP               DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP               DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

ALTER TABLE ipos
    ADD CONSTRAINT foreign_ipos_market FOREIGN KEY (market_id)
        REFERENCES markets (uuid) ON UPDATE CASCADE ON DELETE RESTRICT;
ALTER TABLE ipos
    ADD CONSTRAINT foreign_ipos_company FOREIGN KEY (company_id)
        REFERENCES companies (uuid) ON UPDATE CASCADE ON DELETE RESTRICT;



INSERT INTO sectors (uuid, name)
VALUES (UUID_TO_BIN('04e182e6-7470-4dd7-bdf2-cb5e5599dc0e', true), 'Communication Services');

INSERT INTO currencies (uuid, code, name, display)
VALUES (UUID_TO_BIN('f0dd9459-6123-465c-aec7-03b82c1c5856', true), 'USD', 'American Dollar', '$%s');

INSERT INTO countries (uuid, code, name)
VALUES (UUID_TO_BIN('7ca2f739-a0f3-4e5f-848a-54887e762d3a', true), 'US', 'USA');

INSERT INTO markets (uuid, code, name, currency_id)
VALUES (UUID_TO_BIN('a9da11f6-bb30-47a0-9f27-b52510f1cc6a', true), 'NDX', 'Nasdaq',
        UUID_TO_BIN('f0dd9459-6123-465c-aec7-03b82c1c5856', true));


INSERT INTO companies (uuid, symbol, name, sector_id,
                       address, country_id, phone, email, website,
                       employees, description, founded, ceo, fiscal_year_end,
                       ipo_url, exchange_commission_url, logo_url)
VALUES (UUID_TO_BIN('c2b71e7b-f9f9-4293-8271-77a4ce70c6f0', true), 'PINS', 'Pinterest',
        UUID_TO_BIN('04e182e6-7470-4dd7-bdf2-cb5e5599dc0e', true),
        '505 Brannan Street San Francisco, CA 94107', UUID_TO_BIN('7ca2f739-a0f3-4e5f-848a-54887e762d3a', true),
        '492929893', 'email@email.com', 'https://investor.pinterestinc.com',
        3929, 'Pinterest is a company...', 2012, 'Tomas Cook', 'March 31',
        'http://nasdaq.com/pins', 'http://sec.com/pins/', 'https://s23.q4cdn.com/958601754/files/design/pinterest-logo-footer.png');

INSERT INTO ipos (uuid, market_id, company_id, price_cents_from, price_cents_to, shares, expected_date, created_at,
                  updated_at)
VALUES (UUID_TO_BIN('28e29e39-06e1-4935-8d43-09fdf62ba7dc', true),
        UUID_TO_BIN('a9da11f6-bb30-47a0-9f27-b52510f1cc6a', true),
        UUID_TO_BIN('c2b71e7b-f9f9-4293-8271-77a4ce70c6f0', true),
        2400,
        2500,
        3000000,
        '2020-10-10 00:00:00',
        '2020-10-01 00:00:00',
        '2020-10-05 00:00:00')
