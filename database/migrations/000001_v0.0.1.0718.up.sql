ALTER TABLE "public"."co_company_employee"
    ADD COLUMN "remark" text;

COMMENT ON COLUMN "public"."co_company_employee"."remark" IS '备注';
