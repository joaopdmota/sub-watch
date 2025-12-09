import Repository from "./repository";

export default class UserRepository implements Repository {
	execute (input: any): Promise<any> {
		throw new Error("Method not implemented.");
	}
}