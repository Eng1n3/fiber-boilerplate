-- CreateTable
CREATE TABLE "permissions" (
    "id" VARCHAR(255) NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT,
    "path" VARCHAR(255) NOT NULL,
    "method" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) DEFAULT NULL,
    "deleted_at" TIMESTAMP(3) DEFAULT NULL,

    CONSTRAINT "permissions_pkey" PRIMARY KEY ("id")
);
