/* 员工增加手机号码 */
ALTER TABLE "public"."co_company_employee"
    ADD COLUMN "email" varchar(16);

COMMENT ON COLUMN "public"."co_company_employee"."email" IS '业务邮箱';


/* 员工邮箱长度修改 */
ALTER TABLE "public"."co_company_employee"
ALTER COLUMN "email" TYPE varchar(255) COLLATE "pg_catalog"."default";