-- CreateTable
CREATE TABLE "role_groups_permissions" (
    "role_id" VARCHAR(255) NOT NULL,
    "group_permission_id" VARCHAR(255) NOT NULL,

    CONSTRAINT "role_groups_permissions_pkey" PRIMARY KEY ("role_id")
);

-- AddForeignKey
ALTER TABLE "role_groups_permissions" ADD CONSTRAINT "role_groups_permissions_role_id_fkey" FOREIGN KEY ("role_id") REFERENCES "roles"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "role_groups_permissions" ADD CONSTRAINT "role_groups_permissions_group_permission_id_fkey" FOREIGN KEY ("group_permission_id") REFERENCES "groups_permissions"("id") ON DELETE CASCADE ON UPDATE CASCADE;
