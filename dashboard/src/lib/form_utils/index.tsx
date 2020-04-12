import logger from "../logger";

const searchQueriesKey = "q";
const searchTextsKey = "t";

function getSearchQueries() {
    return getLocationSearch(searchQueriesKey);
}

function getSearchTexts() {
    return getLocationSearch(searchTextsKey);
}

function getLocationSearch(key: string) {
    const searchParams = new URLSearchParams(window.location.search);
    let data = {};
    if (searchParams.has(key)) {
        const value = searchParams.get(key);
        try {
            data = JSON.parse(String(value));
        } catch {
            logger.warning("Ignored failed parse", value);
        }
    }

    return data;
}

function setSearchQueries(route, obj) {
    setLocationSearch(searchQueriesKey, obj, route, true);
}

function setSearchTexts(route, obj) {
    setLocationSearch(searchTextsKey, obj, route, false);
}

function setLocationSearch(key, obj, route, isPush) {
    return new Promise(() => {
        const str = JSON.stringify(obj);
        const searchParams = new URLSearchParams(window.location.search);
        searchParams.set(key, str);
        const paramsStr = searchParams.toString();
        const link = window.location.pathname + "?" + paramsStr;
        if (isPush) {
            route.history.push(link);
        } else {
            console.log("DEBUG HOGE");
            // route.history.replace(link);  too slow
            window.history.replaceState(null, "", link);
        }
    });
}

function mergeDefaultInputsToFormData(index, rawData, formData) {
    const newFormData = {};
    const inputErrorMap = {};
    if (index.Inputs) {
        // Validate
        // フォーム入力がなく、デフォルト値がある場合はセットする
        for (let i = 0, len = index.Inputs.length; i < len; i++) {
            const input = index.Inputs[i];
            let value = formData[input.Name];
            switch (input.Type) {
                case "Text":
                    if (input.Require) {
                        if (!value || value === "") {
                            inputErrorMap[input.Name] = {
                                error: "This is required",
                                type: input.Kind
                            };
                        }
                        newFormData[input.Name] = { value };
                    }
                    break;

                case "Select":
                    if (!value) {
                        if (input.Default) {
                            value = input.Default;
                        } else {
                            let options = input.Options;
                            if (!options) {
                                options = [];
                                const d = rawData[input.DataKey];
                                if (d) {
                                    for (let j = 0, l = d.length; j < l; j++) {
                                        options.push(d[j].Name);
                                    }
                                } else {
                                    options.push("");
                                }
                            }
                            value = options[0];
                        }
                    }
                    newFormData[input.Name] = value;
                    break;
                case "DateTime":
                    newFormData[input.Name] = value;
                    break;
                default:
                    break;
            }
        }
    }
    return { newFormData, inputErrorMap };
}

export default {
    getSearchQueries,
    getSearchTexts,
    setSearchQueries,
    setSearchTexts,
    mergeDefaultInputsToFormData
};
