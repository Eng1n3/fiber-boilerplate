-- CreateTable
CREATE TABLE "role_permissions" (
    "role_id" VARCHAR(255) NOT NULL,
    "permission_id" VARCHAR(255) NOT NULL,

    CONSTRAINT "role_permissions_pkey" PRIMARY KEY ("role_id")
);

-- AddForeignKey
ALTER TABLE "role_permissions" ADD CONSTRAINT "role_permissions_role_id_fkey" FOREIGN KEY ("role_id") REFERENCES "roles"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "role_permissions" ADD CONSTRAINT "role_permissions_permission_id_fkey" FOREIGN KEY ("permission_id") REFERENCES "permissions"("id") ON DELETE CASCADE ON UPDATE CASCADE;
