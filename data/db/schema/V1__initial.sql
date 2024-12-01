CREATE TABLE fam.users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    roles VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE fam.accounts (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    account_name VARCHAR(255) NOT NULL,
    account_type VARCHAR(50),
    balance NUMERIC(15, 2) DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE fam.transactions (
    id SERIAL PRIMARY KEY,
    account_id INT REFERENCES accounts(id),
    amount NUMERIC(15, 2) NOT NULL,
    transaction_type VARCHAR(10) CHECK (transaction_type IN ('credit', 'debit')),
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);