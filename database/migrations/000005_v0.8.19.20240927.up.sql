ALTER TABLE "public"."co_company"
ALTER COLUMN "region" TYPE varchar(128) COLLATE "pg_catalog"."default";
ALTER TABLE "public"."co_company_employee"
ALTER COLUMN "region" TYPE varchar(128) COLLATE "pg_catalog"."default";