CREATE TABLE sectors
(
    uuid  BINARY(16) PRIMARY KEY,
    alias VARCHAR(255) NOT NULL,
    name  VARCHAR(255) NOT NULL
);
ALTER TABLE sectors
    ADD CONSTRAINT unique_sectors_alias UNIQUE KEY (alias);

CREATE TABLE industries
(
    uuid  BINARY(16) PRIMARY KEY,
    alias VARCHAR(255) NOT NULL,
    name  VARCHAR(255) NOT NULL
);
ALTER TABLE industries
    ADD CONSTRAINT unique_industries_alias UNIQUE KEY (alias);

CREATE TABLE currencies
(
    uuid    BINARY(16) PRIMARY KEY,
    code    CHAR(3)   NOT NULL,
    name    VARCHAR(255) NOT NULL,
    display VARCHAR(32)  NOT NULL
);
ALTER TABLE currencies
    ADD CONSTRAINT unique_currencies_code UNIQUE KEY (code);

CREATE TABLE countries
(
    uuid BINARY(16) PRIMARY KEY,
    code CHAR(3)   NOT NULL,
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
    industry_id             BINARY(16)         NOT NULL,
    address                 VARCHAR(255)       NOT NULL,
    country_id              BINARY(16)         NOT NULL,
    phone                   VARCHAR(32)        NULL DEFAULT NULL,
    email                   VARCHAR(128)       NULL DEFAULT NULL,
    website                 VARCHAR(255)       NULL DEFAULT NULL,
    facebook                VARCHAR(255)       NULL DEFAULT NULL,
    twitter                 VARCHAR(255)       NULL DEFAULT NULL,
    linkedin                VARCHAR(255)       NULL DEFAULT NULL,
    instagram               VARCHAR(255)       NULL DEFAULT NULL,
    pinterest               VARCHAR(255)       NULL DEFAULT NULL,
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
    ADD CONSTRAINT foreign_companies_industry FOREIGN KEY (industry_id)
        REFERENCES industries (uuid) ON UPDATE CASCADE ON DELETE RESTRICT;

ALTER TABLE companies
    ADD CONSTRAINT foreign_companies_country FOREIGN KEY (country_id)
        REFERENCES countries (uuid) ON UPDATE CASCADE ON DELETE RESTRICT;

CREATE TABLE ipos
(
    uuid             BINARY(16) PRIMARY KEY,
    alias            VARCHAR(128)       NOT NULL,
    intro            VARCHAR(255)       NOT NULL,
    market_id        BINARY(16)         NOT NULL,
    company_id       BINARY(16)         NOT NULL,
    price_cents_from MEDIUMINT UNSIGNED,
    price_cents_to   MEDIUMINT UNSIGNED NULL DEFAULT NULL,
    shares           INT UNSIGNED       NULL DEFAULT NULL,
    expected_date    DATE,
    created_at       TIMESTAMP               DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE ipos
    ADD CONSTRAINT unique_ipos_alias UNIQUE KEY (alias);
ALTER TABLE ipos
    ADD CONSTRAINT foreign_ipos_market FOREIGN KEY (market_id)
        REFERENCES markets (uuid) ON UPDATE CASCADE ON DELETE RESTRICT;
ALTER TABLE ipos
    ADD CONSTRAINT foreign_ipos_company FOREIGN KEY (company_id)
        REFERENCES companies (uuid) ON UPDATE CASCADE ON DELETE RESTRICT;

