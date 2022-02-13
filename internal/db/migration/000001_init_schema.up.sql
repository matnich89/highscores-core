CREATE TABLE "characters" (
                             "id" BIGSERIAL PRIMARY KEY,
                             "name" varchar UNIQUE NOT NULL,
                             "last_check" timestamp NOT NULL,
                             "created_at" timestamp NOT NULL
);

CREATE TABLE "xp_records" (
                             "id" BIGSERIAL PRIMARY KEY,
                             "skill" varchar NOT NULL,
                             "character_id" bigint UNIQUE NOT NULL,
                             "last_xp_count" bigint NOT NULL,
                             "last_level" int NOT NULL,
                             "last_position" bigint NOT NULL,
                             "updated_at" timestamp NOT NULL,
                             "created_at" timestamp NOT NULL
);

ALTER TABLE "xp_records" ADD FOREIGN KEY ("character_id") REFERENCES "characters" ("id");
