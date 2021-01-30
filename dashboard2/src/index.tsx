import React from "react";
import ReactDOM from "react-dom";

import auth from "./appui/src/apps/auth";
import service from "./appui/src/apps/service";
import provider from "./appui/src/provider";
import app from "./app";

$(function () {
    provider.register(new app.Provider());

    if ("auth" in window) {
        service.init();
    } else {
        auth.init();
    }
});
