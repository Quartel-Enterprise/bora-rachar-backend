import { prisma } from "..";

// For soft-delete implementation reference:
// https://www.prisma.io/docs/concepts/components/prisma-client/middleware/soft-delete-middleware

export function registerSoftDeleteMiddleware() {
    // Behavior of read/update queries for deleted records
    prisma.$use(async (params, next) => {
        if (!params.args) {
            params.args = {};
        }

        // All view tables need to have "_ViewTable" suffix
        // DO NOT apply soft delete to the View Tables
        if (/_ViewTable$/.test(params.model)) {
            return next(params);
        }

        if (params.action === "findUnique" || params.action === "findFirst") {
            // Change to findFirst - you cannot filter
            // by anything except ID / unique with findUnique
            params.action = "findFirst";
            // Add 'deletedAt' filter
            // ID filter maintained

            if (params.args.where.deletedAt === undefined) {
                params.args.where["deletedAt"] = null;
            }
        }
        if (params.action === "findMany") {
            // Find many queries
            if (params.args?.where) {
                if (params.args.where.deletedAt === undefined) {
                    // Exclude deleted records if they have not been explicitly requested
                    params.args.where["deletedAt"] = null;
                }
            } else {
                params.args["where"] = { deletedAt: null };
            }
        }
        return next(params);
    });

    // Behavior of delete queries for all models
    prisma.$use(async (params, next) => {
        // Check incoming query type
        if (params.action == "delete") {
            // Delete queries
            // Overwrites action to a field update
            params.action = "update";
            params.args["data"] = { deletedAt: new Date() };
        }
        if (params.action == "deleteMany") {
            // Delete many queries
            params.action = "updateMany";
            if (params.args.data != undefined) {
                params.args.data["deletedAt"] = new Date();
            } else {
                params.args["data"] = { deletedAt: new Date() };
            }
        }

        return next(params);
    });
}
