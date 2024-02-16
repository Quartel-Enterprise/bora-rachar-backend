import * as http from "http";
import * as express from "express";
import "reflect-metadata";
import { getRoutes } from "@routes/api/v1";
import { errorHandler } from "@controllers/middlewares/error-handler";
import { serverConfig } from "@config/index";
import { requestLogger } from "@utils/request-logger";

export default function buildApp() {
    const app = express();
    app.enable("trust proxy");

    if (serverConfig.isDevelopment) {
        app.use(requestLogger);
    }

    app.use(express.json({ limit: "10mb" }));
    app.use(express.urlencoded({ extended: true, limit: "10mb" }));

    app.use("/api", getRoutes());
    app.use(errorHandler);

    return http.createServer(app);
}
