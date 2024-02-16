import { serverConfig } from "./config";
import buildApp from "./app";
import { registerPrismaMiddlewares } from "@database/prisma/middlewares";

process.env.TZ = serverConfig.timezone;

function logError(error: Error) {
    console.error(error);
    process.exit(1);
}

const init = async () => {
    registerPrismaMiddlewares();

    const app = buildApp();

    const port = serverConfig.port;

    app.listen(port, "::", () => {
        if (process.env.NODE_ENV !== "production") {
            console.log(`API http server running on port ${port}`);
        }
    });

    app.keepAliveTimeout = serverConfig.keepAliveTimeout;
    app.headersTimeout = serverConfig.headersTimeout;
};

init().catch((e: Error) => {
    logError(e);
});
