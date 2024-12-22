ALTER TABLE "public"."merchant_employee"
    ADD COLUMN "country_code" varchar(16),
  ADD COLUMN "region" varchar(128),
  ADD COLUMN "email" varchar(255),
  ADD COLUMN "weixin_account" varchar(128),
  ADD COLUMN "address" varchar(512);

COMMENT ON COLUMN "public"."merchant_employee"."country_code" IS '所属国家编码';

COMMENT ON COLUMN "public"."merchant_employee"."region" IS '所属地区';

COMMENT ON COLUMN "public"."merchant_employee"."email" IS '业务邮箱';

COMMENT ON COLUMN "public"."merchant_employee"."weixin_account" IS '微信号';

COMMENT ON COLUMN "public"."merchant_employee"."address" IS '地址';

ALTER TABLE "public"."operator_employee"
    ADD COLUMN "country_code" varchar(16),
  ADD COLUMN "region" varchar(128),
  ADD COLUMN "email" varchar(255),
  ADD COLUMN "weixin_account" varchar(128),
  ADD COLUMN "address" varchar(512);

COMMENT ON COLUMN "public"."operator_employee"."country_code" IS '所属国家编码';

COMMENT ON COLUMN "public"."operator_employee"."region" IS '所属地区';

COMMENT ON COLUMN "public"."operator_employee"."email" IS '业务邮箱';

COMMENT ON COLUMN "public"."operator_employee"."weixin_account" IS '微信号';

COMMENT ON COLUMN "public"."operator_employee"."address" IS '地址';

ALTER TABLE "public"."co_companyr_employee"
    ADD COLUMN "country_code" varchar(16),
  ADD COLUMN "region" varchar(128),
  ADD COLUMN "email" varchar(255),
  ADD COLUMN "weixin_account" varchar(128),
  ADD COLUMN "address" varchar(512);

COMMENT ON COLUMN "public"."co_company_employee"."country_code" IS '所属国家编码';

COMMENT ON COLUMN "public"."co_company_employee"."region" IS '所属地区';

COMMENT ON COLUMN "public"."co_company_employee"."email" IS '业务邮箱';

COMMENT ON COLUMN "public"."co_company_employee"."weixin_account" IS '微信号';

COMMENT ON COLUMN "public"."co_company_employee"."address" IS '地址';