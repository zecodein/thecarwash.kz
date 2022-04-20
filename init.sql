CREATE TABLE IF NOT EXISTS "entity"(
    "entity_id" BIGSERIAL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "address" TEXT NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS "user"(
    "user_id" BIGSERIAL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "number" TEXT NOT NULL UNIQUE,
    "access" TEXT NOT NULL DEFAULT 'basic_user',
    "password" TEXT NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS "service"(
    "service_id" BIGSERIAL PRIMARY KEY,
    "type" TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS "price"(
    "price_id" BIGSERIAL PRIMARY KEY,
    "entity_id" BIGSERIAL NOT NULL REFERENCES "entity"("entity_id") ON DELETE CASCADE,
    "service_id" BIGSERIAL NOT NULL REFERENCES "service"("service_id") ON DELETE CASCADE,
    "price" BIGSERIAL NOT NULL
);

CREATE TABLE IF NOT EXISTS "order"(
    "order_id" BIGSERIAL PRIMARY KEY,
    "user_id" BIGSERIAL NOT NULL REFERENCES "user"("user_id") ON DELETE CASCADE,
    "entity_id" BIGSERIAL NOT NULL REFERENCES "entity"("entity_id") ON DELETE CASCADE,
    "service_id" BIGSERIAL NOT NULL REFERENCES "service"("service_id") ON DELETE CASCADE,
    "price_id" BIGSERIAL NOT NULL REFERENCES "price"("price_id") ON DELETE CASCADE,
    "status" TEXT NOT NULL,
    "payment" TEXT NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NOT NULL
);