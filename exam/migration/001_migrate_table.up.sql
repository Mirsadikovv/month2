CREATE TABLE "user" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "age" INT,
    "phone" VARCHAR(50) UNIQUE NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);



CREATE TABLE IF NOT EXISTS "todo_list" (
    "task_id" UUID PRIMARY KEY,
    "task_name" VARCHAR,
    "task_time" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "task_comment" VARCHAR(50),
    "task_priority" VARCHAR,
    "user_id" UUID REFERENCES "user"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX "idx_sesion_user_id" ON "todo_list"("user_id");