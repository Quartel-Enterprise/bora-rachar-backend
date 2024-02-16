import { Request, Response } from "express";
import "express-async-errors";
import { ConstStatusCode } from "@backend-types/status-code";

// ? "_next" parameter is not used but is needed for "express-async-errors" package
// ? Removing it causes errors on every request
export function errorHandler(error: Error, request: Request, res: Response) {
    console.error(error);

    return res.status(ConstStatusCode.serverError).json({ message: "Internal Server Error" });
}
