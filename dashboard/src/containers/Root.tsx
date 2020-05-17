import * as React from "react";
import { Provider } from "react-redux";

import store from "../store";

import Router from "./Router";

export default class Root extends React.Component {
    public render() {
        return (
            <Provider store={store}>
                <Router />
            </Provider>
        );
    }
}
