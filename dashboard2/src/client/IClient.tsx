export interface IClient {
    loginWithToken(input: any): void;
    login(input: any): void;
    logout(input: any): void;
    get_service_index(input: any): void;
    get_queries(input: any): void;
}
