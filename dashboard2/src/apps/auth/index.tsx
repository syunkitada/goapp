import Login from "../../components/login/Login";
import Loading from "../../components/login/Loading";
import service from "../service";
import client from "../../client";

function loginWithToken() {
    Loading.Render({ id: "root" });

    client.loginWithToken({
        onSuccess: function (input: any) {
            service.init();
        },
        onError: function (input: any) {
            console.log("error", input);
            renderLoginView();
        }
    });
}

function login(input: any) {
    const { userName, password } = input;
    client.login({
        userName,
        password,
        onSuccess: function (input: any) {
            service.init();
        },
        onError: function (input: any) {
            renderLoginView();
        }
    });
}

function logout() {
    client.logout({
        onSuccess: function (input: any) {
            renderLoginView();
        },
        onError: function (input: any) {
            console.log("DEBUG onError");
        }
    });
}

function renderLoginView() {
    Login.Render({
        id: "root",
        onSubmit: function (input: any) {
            login(input);
        }
    });
}

function init() {
    loginWithToken();
}

const index = {
    init,
    logout
};
export default index;
