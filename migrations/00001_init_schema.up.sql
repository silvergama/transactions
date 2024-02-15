CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    document_number VARCHAR(20) NOT NULL
);

CREATE TABLE operation_types (
    id SERIAL PRIMARY KEY,
    description VARCHAR(255) NOT NULL
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    account_id INTEGER REFERENCES accounts(id),
    operation_type_id INTEGER REFERENCES operation_types(id),
    amount NUMERIC NOT NULL,
    event_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    -- Adicionando as chaves estrangeiras explícitas
    FOREIGN KEY (account_id) REFERENCES accounts(id),
    FOREIGN KEY (operation_type_id) REFERENCES operation_types(id)
);
