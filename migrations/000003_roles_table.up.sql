-- CreateTable
CREATE TABLE "roles" (
    "id" VARCHAR(255) NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "description" TEXT,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) DEFAULT NULL,
    "deleted_at" TIMESTAMP(3) DEFAULT NULL,

    CONSTRAINT "roles_pkey" PRIMARY KEY ("id")
);
