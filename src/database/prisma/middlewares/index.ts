import { registerSoftDeleteMiddleware } from "./soft-delete";

export function registerPrismaMiddlewares() {
    registerSoftDeleteMiddleware();
}
