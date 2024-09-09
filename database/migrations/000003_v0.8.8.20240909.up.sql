ALTER TABLE "public"."co_company"
    ADD COLUMN "country_code" varchar(16),
  ADD COLUMN "region" varchar(16);

COMMENT ON COLUMN "public"."co_company"."country_code" IS '所属国家编码';

COMMENT ON COLUMN "public"."co_company"."region" IS '所属地区';

ALTER TABLE "public"."co_company_employee"
    ADD COLUMN "country_code" varchar(16),
  ADD COLUMN "region" varchar(16);

COMMENT ON COLUMN "public"."co_company_employee"."country_code" IS '所属国家编码';

COMMENT ON COLUMN "public"."co_company_employee"."region" IS '所属地区';ALTER TABLE "public"."co_company"
    ADD COLUMN "country_code" varchar(16),
  ADD COLUMN "region" varchar(16);

COMMENT ON COLUMN "public"."co_company"."country_code" IS '所属国家编码';

COMMENT ON COLUMN "public"."co_company"."region" IS '所属地区';

ALTER TABLE "public"."co_company_employee"
    ADD COLUMN "country_code" varchar(16),
  ADD COLUMN "region" varchar(16);

COMMENT ON COLUMN "public"."co_company_employee"."country_code" IS '所属国家编码';

COMMENT ON COLUMN "public"."co_company_employee"."region" IS '所属地区';