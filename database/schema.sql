CREATE TABLE sectors
(
    uuid  BINARY(16) PRIMARY KEY,
    alias VARCHAR(128) NOT NULL,
    name  VARCHAR(255) NOT NULL
);
ALTER TABLE sectors
    ADD CONSTRAINT unique_sectors_alias UNIQUE KEY (alias);

CREATE TABLE industries
(
    uuid  BINARY(16) PRIMARY KEY,
    alias VARCHAR(128) NOT NULL,
    name  VARCHAR(255) NOT NULL
);
ALTER TABLE industries
    ADD CONSTRAINT unique_industries_alias UNIQUE KEY (alias);

CREATE TABLE currencies
(
    uuid    BINARY(16) PRIMARY KEY,
    code    VARCHAR(3)   NOT NULL,
    name    VARCHAR(255) NOT NULL,
    display VARCHAR(32)  NOT NULL
);
ALTER TABLE currencies
    ADD CONSTRAINT unique_currencies_code UNIQUE KEY (code);

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



INSERT INTO sectors (uuid, alias, name)
VALUES (UUID_TO_BIN('04e182e6-7470-4dd7-bdf2-cb5e5599dc0e', true), 'communication-services', 'Communication Services'),
       (UUID_TO_BIN('fff90079-e45a-4c9d-b25c-03fd4bfb38aa', true), 'technology', 'Technology');

INSERT INTO industries (uuid, alias, name)
VALUES (UUID_TO_BIN('0990dc7e-310a-41d3-bc25-a9141e7343c7', true), 'internet-content-and-information',
        'Internet Content & Information'),
       (UUID_TO_BIN('a2f4b42a-14f5-4287-a16e-e1ad613579e5', true), 'solar', 'Solar');

INSERT INTO currencies (uuid, code, name, display)
VALUES (UUID_TO_BIN('f0dd9459-6123-465c-aec7-03b82c1c5856', true), 'USD', 'American Dollar', '$%s');

INSERT INTO countries (uuid, code, name)
VALUES (UUID_TO_BIN('7ca2f739-a0f3-4e5f-848a-54887e762d3a', true), 'US', 'United States of America');

INSERT INTO markets (uuid, code, name, currency_id)
VALUES (UUID_TO_BIN('a9da11f6-bb30-47a0-9f27-b52510f1cc6a', true), 'NQGB', 'Nasdaq Global',
        UUID_TO_BIN('f0dd9459-6123-465c-aec7-03b82c1c5856', true));


INSERT INTO companies (uuid, symbol, name, sector_id, industry_id,
                       address, country_id, phone, email, website,
                       facebook, twitter, linkedin, instagram, pinterest,
                       employees, description, founded, ceo, fiscal_year_end,
                       ipo_url, exchange_commission_url, logo_url)
VALUES (UUID_TO_BIN('c2b71e7b-f9f9-4293-8271-77a4ce70c6f0', true), 'PINS', 'Pinterest',
        UUID_TO_BIN('04e182e6-7470-4dd7-bdf2-cb5e5599dc0e', true),
        UUID_TO_BIN('0990dc7e-310a-41d3-bc25-a9141e7343c7', true),
        '505 Brannan Street San Francisco, CA 94107', UUID_TO_BIN('7ca2f739-a0f3-4e5f-848a-54887e762d3a', true),
        '+1 415 762-7100', null, 'https://investor.pinterestinc.com',
        'https://www.facebook.com/pinterest', 'https://twitter.com/Pinterest',
        'https://www.linkedin.com/company/pinterest', 'https://www.instagram.com/pinterest',
        'https://www.pinterest.es/pinterest',
        2200,
        'Pinterest is an online product and idea discovery platform that helps users gather ideas on everything from recipes to cook to destinations to travel to. Founded in 2010, the platform consists of a largely female audience, at roughly two thirds of its more than 365 million monthly active users. The company generates revenue by selling digital ads and is now rolling out more in-platform e-commerce features.',
        2010, 'Ben Silbermann', 'Dec 31',
        'https://www.nasdaq.com/market-activity/ipos/overview?dealId=842469-89263', 'https://sec.report/Ticker/PINS',
        '/assets/images/pinterest-logo.jpg')/*,

       (UUID_TO_BIN('1da86fb0-44a8-48ed-92c5-ff05f43565a6', true), 'ARRY', 'Array Technologies',
        UUID_TO_BIN('fff90079-e45a-4c9d-b25c-03fd4bfb38aa', true),
        UUID_TO_BIN('a2f4b42a-14f5-4287-a16e-e1ad613579e5', true),
        '3901 Midway Place NE Albuquerque, NM 87109', UUID_TO_BIN('7ca2f739-a0f3-4e5f-848a-54887e762d3a', true),
        '(505) 437 001', 'investors@arraytechinc.com ', 'https://ir.arraytechinc.com/',
        null, 'https://twitter.com/arraytechinc', 'https://www.linkedin.com/company/array-technologies-inc',
        'https://www.instagram.com/arraytechinc', null,
        343,
        'We are one of the world’s largest manufacturers of ground-mounting systems used in solar energy projects. Our principal product is an integrated system of steel supports, electric motors, gearboxes and electronic controllers commonly referred to as a single-axis “tracker.” Trackers move solar panels throughout the day to maintain an optimal orientation to the sun, which significantly increases their energy production. Solar energy projects that use trackers generate up to 25% more energy and deliver a 22% lower levelized cost of energy (“LCOE”) than projects that use “fixed tilt” mounting systems, according to BloombergNEF. Trackers represent between 10% and 15% of the cost of constructing a ground-mounted solar energy project, and approximately 70% of all ground-mounted solar energy projects constructed in the U.S. during 2019 utilized trackers according to BloombergNEF and IHS Markit, respectively. Our trackers use a patented design that allows one motor to drive multiple rows of solar panels through articulated driveline joints. To avoid infringing on our U.S. patent, our competitors must use designs that we believe are inherently less efficient and reliable. For example, our largest competitor’s design requires one motor for each row of solar panels. As a result, we believe our products have greater reliability, lower installation costs, reduced maintenance requirements and competitive manufacturing costs. Our core U.S. patent on a linked-row, rotating gear drive system does not expire until February 5, 2030. We sell our products to engineering, procurement and construction firms (“EPCs”) that build solar energy projects and to large solar developers, independent power producers and utilities, often under master supply agreements or multi-year procurement contracts. In 2019, we derived 87%, 8% and 5% of our revenues from customers in the U.S., Australia and rest of the world, respectively. As of June 30, 2020, we had shipped more than 21 GWs of our trackers to customers worldwide. We are a U.S. company and our headquarters and principal manufacturing facility are in Albuquerque, New Mexico. ',
        1992, 'Jim Fusaro', 'Dec 31',
        'https://www.nasdaq.com/market-activity/ipos/overview?dealId=1128643-94122', 'https://sec.report/Ticker/ARRY',
        '/assets/images/array-technologies-logo.jpg')*/;

INSERT INTO ipos (uuid, alias, market_id, company_id, price_cents_from, price_cents_to, shares, expected_date)
VALUES (UUID_TO_BIN('28e29e39-06e1-4935-8d43-09fdf62ba7dc', true),
        'pinterest',
        UUID_TO_BIN('a9da11f6-bb30-47a0-9f27-b52510f1cc6a', true),
        UUID_TO_BIN('c2b71e7b-f9f9-4293-8271-77a4ce70c6f0', true),
        2200,
        null,
        47500000,
        '2019-04-18')/*,
       (UUID_TO_BIN('410ad8b5-f713-4de1-b67d-3658e4a89723', true),
        'array-technologies',
        UUID_TO_BIN('a9da11f6-bb30-47a0-9f27-b52510f1cc6a', true),
        UUID_TO_BIN('1da86fb0-44a8-48ed-92c5-ff05f43565a6', true),
        1900,
        null,
        75000000,
        '2020-10-15')*/
