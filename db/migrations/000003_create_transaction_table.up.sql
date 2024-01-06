CREATE TYPE "transaction_type" AS ENUM (
    'recharge',
    'gift',
    'withdrawal',
    'payment',
    'refund',
    'transfer'
    );

CREATE TABLE "transaction"
(
    id               SERIAL PRIMARY KEY,
    wallet_id        INT         NOT NULL REFERENCES "wallet" (id),
    amount           DECIMAL(20, 0),
    transaction_type transaction_type,
    description      VARCHAR(255),
    discount_code    VARCHAR(255),
    created_at       TIMESTAMPTZ NOT NULL DEFAULT now()
);


CREATE INDEX ON "transaction" (wallet_id);
CREATE INDEX ON "transaction" (transaction_type);
CREATE INDEX ON "transaction" (discount_code);