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
    uuid       BINARY(16) PRIMARY KEY,
    code       VARCHAR(16)  NOT NULL,
    name       VARCHAR(255) NOT NULL,
    country_id BINARY(16)   NOT NULL
);
ALTER TABLE markets
    ADD CONSTRAINT unique_markets_code UNIQUE KEY (code);
ALTER TABLE markets
    ADD CONSTRAINT foreign_markets_country FOREIGN KEY (country_id)
        REFERENCES countries (uuid) ON UPDATE CASCADE ON DELETE RESTRICT;


CREATE TABLE companies
(
    uuid   BINARY(16) PRIMARY KEY,
    symbol VARCHAR(16)  NOT NULL,
    name   VARCHAR(255) NOT NULL
);
ALTER TABLE companies
    ADD CONSTRAINT unique_companies_symbol UNIQUE KEY (symbol);


CREATE TABLE ipos
(
    uuid          BINARY(16) PRIMARY KEY,
    market_id     BINARY(16) NOT NULL,
    company_id    BINARY(16) NOT NULL,
    expected_date DATETIME   NULL DEFAULT NULL,
    created_at    TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP       DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
ALTER TABLE ipos
    ADD CONSTRAINT foreign_ipos_market FOREIGN KEY (market_id)
        REFERENCES markets (uuid) ON UPDATE CASCADE ON DELETE RESTRICT;
ALTER TABLE ipos
    ADD CONSTRAINT foreign_ipos_company FOREIGN KEY (company_id)
        REFERENCES companies (uuid) ON UPDATE CASCADE ON DELETE RESTRICT;



INSERT INTO countries (uuid, code, name)
VALUES (UUID_TO_BIN('5a1df32f-57c8-4e60-9ba6-96a12b741597', true), 'US', 'USA');

INSERT INTO markets (uuid, code, name, country_id)
VALUES (UUID_TO_BIN('7aed8bb9-8d1e-472d-930f-3f590977611e', true), 'NDX', 'Nasdaq', UUID_TO_BIN('5a1df32f-57c8-4e60-9ba6-96a12b741597', true));

INSERT INTO companies (uuid, symbol, name)
VALUES (UUID_TO_BIN('ccfd8377-309f-45ab-9dce-9855a6174af1', true), 'PINS', 'Pinterest');

INSERT INTO ipos (uuid, market_id, company_id, expected_date)
VALUES (UUID_TO_BIN('b09ca35a-b02b-49d1-b457-7f14e872060f', true),
        UUID_TO_BIN('7aed8bb9-8d1e-472d-930f-3f590977611e', true),
        UUID_TO_BIN('ccfd8377-309f-45ab-9dce-9855a6174af1', true),
        '2020-05-05 00:00:0')
