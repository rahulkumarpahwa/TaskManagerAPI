CREATE TYPE task_status AS ENUM (
    'Completed',
    'Pending',
    'UnderProcess',
    'Skipped',
    'NotDone'
);

CREATE SEQUENCE IF NOT EXISTS tasks_id_seq;

-- Table Definition
CREATE TABLE "public"."tasks" (
    "id" int4 NOT NULL DEFAULT nextval ('tasks_id_seq'::regclass),
    "title" text NOT NULL,
    "description" text,
    "status" task_status NOT NULL DEFAULT 'Pending',
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);