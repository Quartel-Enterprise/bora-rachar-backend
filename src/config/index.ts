import * as dotenv from "dotenv";

dotenv.config();

const parseEnvStr = (env: string, defaultEnv?: string): string => {
    return process.env[env] || defaultEnv;
};

const parseEnvInt = (env: string, defaultEnv?: number): number => {
    if (env in process.env) {
        return parseInt(process.env[env], 10);
    }
    return defaultEnv;
};

interface AppConfig {
    env: string;
    timezone: string;
    port: number;
    origin: string;
    isProduction: boolean;
    isDevelopment: boolean;
    jwtSecret: string;
    keepAliveTimeout: number;
    headersTimeout: number;
    mysql: {
        host: string;
        port: string;
        username: string;
        password: string;
        database: string;
        dialect: "mysql"; // ! Should always be mysql
        synchronize: false; // ! Should always be false
    };
}

export const serverConfig: AppConfig = {
    env: parseEnvStr("NODE_ENV", "development"),
    timezone: parseEnvStr("TZ", "America/Sao Paulo"),
    port: parseEnvInt("SERVER_PORT", 3001),
    origin: parseEnvStr("SERVER_ORIGIN", "https://backoffice.melhorplano.net"),
    isProduction: process.env.NODE_ENV === "production",
    isDevelopment: process.env.NODE_ENV === "development",
    jwtSecret: parseEnvStr("JWT_SECRET", "secret-key"),
    keepAliveTimeout: parseEnvInt("SERVER_KEEP_ALIVE_TIMEOUT", 61 * 1000),
    headersTimeout: parseEnvInt("SERVER_HEADERS_TIMEOUT", 62 * 1000),
    mysql: {
        host: parseEnvStr("MYSQL_HOST", "127.0.0.1"),
        port: parseEnvStr("MYSQL_PORT", "3306"),
        username: parseEnvStr("MYSQL_USERNAME", "root"),
        password: parseEnvStr("MYSQL_PASSWORD", "root"),
        database: parseEnvStr("MYSQL_DATABASE", "DB_backoffice"),
        dialect: "mysql",
        synchronize: false,
    },
};
