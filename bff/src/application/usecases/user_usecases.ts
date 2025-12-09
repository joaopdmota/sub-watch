import Usecase from "./usecase";

export default class UserUsesCases implements Usecase {
	execute (input: any): Promise<any> {
		throw new Error("Method not implemented.");
	}
}
