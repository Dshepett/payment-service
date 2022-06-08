CREATE TABLE payments
(
    id         SERIAL PRIMARY KEY,
    user_id    INT,
    user_email VARCHAR(255),
    amount     INT,
    currency   VARCHAR(10),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    status     VARCHAR(10)
);