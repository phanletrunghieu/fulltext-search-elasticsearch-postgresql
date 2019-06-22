DROP TABLE IF EXISTS "posts";
DROP SEQUENCE IF EXISTS posts_id_seq;
CREATE SEQUENCE posts_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 START 1 CACHE 1;

CREATE TABLE "public"."posts" (
    "id" integer DEFAULT nextval('posts_id_seq') NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "title" text,
    "content" text,
    CONSTRAINT "posts_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "idx_posts_deleted_at" ON "public"."posts" USING btree ("deleted_at");

INSERT INTO "posts" ("created_at", "updated_at", "title", "content") VALUES ('2019-06-22 19:27:49.499825+00',	'2019-06-22 19:27:49.499825+00',	'Tôi là ai',	'Hiếu đẹp trai đây');