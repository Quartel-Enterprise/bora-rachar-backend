import { Router } from "express";
import { buildRouter } from "../../utils";

export function getRoutes(): Router {
    return buildRouter((router) => {
        router.use("/auth", () => {});
        return router;
    });
}
