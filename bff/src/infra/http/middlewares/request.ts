export type HttpNextFunction = () => void;

export interface HttpResponse {
    statusCode: number;
    on(event: 'finish', listener: () => void): void;
}

export interface HttpRequest {
    method: string;
    url: string;
}

export type AbstractHttpMiddleware = (
    req: HttpRequest,
    res: HttpResponse,
    next: HttpNextFunction
) => void;