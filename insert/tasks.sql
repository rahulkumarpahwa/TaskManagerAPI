CREATE TYPE task_status AS ENUM (
    'Completed',
    'Pending',
    'UnderProcess',
    'Skipped',
    'UnCompleted'
);

CREATE SEQUENCE IF NOT EXISTS tasks_id_seq;

-- Table Definition
CREATE TABLE "public"."tasks" (
    "id" int4 NOT NULL DEFAULT nextval ('tasks_id_seq'::regclass),
    "title" text NOT NULL,
    "description" text,
    "status" task_status NOT NULL DEFAULT 'UnCompleted',
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "modified_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "user_id" int4 NOT NULL,
    "is_favorite" boolean NOT NULL DEFAULT false, 
    CONSTRAINT tasks_user_id_fkey
    FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE CASCADE,
    PRIMARY KEY ("id") 
);