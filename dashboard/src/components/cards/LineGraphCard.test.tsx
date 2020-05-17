import React from "react";
import { render, unmountComponentAtNode } from "react-dom";
import { act } from "react-dom/test-utils";
import { createStore } from "redux";

import reducers from "../../reducers/";
import LineGraphCard from "./LineGraphCard";

let container: any;

beforeEach(() => {
    container = document.createElement("div");
    document.body.appendChild(container);
});

afterEach(() => {
    unmountComponentAtNode(container);
    container.remove();
    container = null;
});

it("renders with or without a name", () => {
    const store = createStore(reducers);

    act(() => {
        const data = {
            Name: "test",
            Values: [{ name: "hoge", uv: 4000, pv: 2400, amt: 2400 }],
            Keys: ["pv", "uv"]
        };
        render(<LineGraphCard store={store} data={data} />, container);
    });
    // expect(container.textContent).toBe("test");
});
