export default interface Repository {
	execute (input: any): Promise<any>;
}