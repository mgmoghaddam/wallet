CREATE TABLE "member"
(
    id         SERIAL PRIMARY KEY,
    first_name VARCHAR(20),
    last_name  VARCHAR(20),
    email      VARCHAR(50),
    phone      VARCHAR(20) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);


CREATE INDEX ON "member" (phone);