import HttpServer from "./server";

export default class MainController {
	constructor (readonly httpServer: HttpServer) {
		httpServer.on("get", "/", async function (params: any, body: any, headers: any) {
			return { status: "ok" };
		});
	}
}