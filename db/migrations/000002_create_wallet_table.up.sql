create table "wallet"
(
    id          SERIAL PRIMARY KEY,
    member_id     INT         NOT NULL    REFERENCES "member" (id),
    wallet_name VARCHAR(255) NOT NULL default 'My Wallet',
    balance     DECIMAL(20, 0)        DEFAULT 0,
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ  NOT NULL DEFAULT now(),
    UNIQUE (member_id, wallet_name)
);


CREATE INDEX ON "wallet" (member_id);