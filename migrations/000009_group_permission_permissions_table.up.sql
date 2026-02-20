-- CreateTable
CREATE TABLE "group_permission_permissions" (
    "group_permission_id" TEXT NOT NULL,
    "permission_id" TEXT NOT NULL,

    CONSTRAINT "group_permission_permissions_pkey" PRIMARY KEY ("group_permission_id")
);

-- AddForeignKey
ALTER TABLE "group_permission_permissions" ADD CONSTRAINT "group_permission_permissions_group_permission_id_fkey" FOREIGN KEY ("group_permission_id") REFERENCES "groups_permissions"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "group_permission_permissions" ADD CONSTRAINT "group_permission_permissions_permission_id_fkey" FOREIGN KEY ("permission_id") REFERENCES "permissions"("id") ON DELETE CASCADE ON UPDATE CASCADE;
