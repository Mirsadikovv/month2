CREATE TABLEIF NOT EXISTS "user"(
    "id" UUID PRIMARY KEY
    "name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "phone" VARCHAR(50) UNIQUE NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMEST NOT NULL,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMEST NOT NULL

);

CREATE TABLE IF NOT EXISTS "session"(
    "id" UUID PRIMARY KEY,
    "user_id" UUID REFERENCES "user"("id"),
    "data" TEXT,
    "expires_at" TIMESTAMP,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMEST NOT NULL,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMEST NOT NULL

);


CREATE INDEX "idx_session_usesr_id" ON "session"("user_id");