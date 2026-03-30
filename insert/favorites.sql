DROP TABLE IF EXISTS "public"."user_tasks";

-- Table Definition
CREATE TABLE "public"."user_tasks" (
    "user_id" int4 NOT NULL,
    "task_id" int4 NOT NULL,
    "relation_type" text NOT NULL CHECK (relation_type = ANY (ARRAY['favorite'::text, 'not-favorite'::text])),
    "time_added" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("user_id","task_id","relation_type")
);