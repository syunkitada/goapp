import "./Form.css";

import data from "../../data";
import service from "../../apps/service";
import locationData from "../../data/locationData";
import Dashboard from "../core/Dashboard";
import Icon from "../icon/Icon";

export function Render(input: any) {
    const { View, useRootModal, selectedData, onSubmit } = input;
    let { id } = input;
    if (useRootModal) {
        id = Dashboard.RootModal.GetContentId();
    }
    const location = locationData.getLocationData();

    const keyPrefix = `${id}-BasicForm-`;
    const formId = `${keyPrefix}form`;
    const submitButtonWrapperId = `${keyPrefix}submitButtonWrapperId`;
    const submitButtonId = `${keyPrefix}submitButtonId`;
    const progressId = `${keyPrefix}progressId`;
    const fieldTextClass = `${keyPrefix}fieldText`;
    const fieldCheckClass = `${keyPrefix}fieldCheck`;
    const fieldParams: any = {};

    const fields: any = [];
    if (selectedData) {
        fields.push('<ul class="collection">');
        for (let i = 0, len = selectedData.length; i < len; i++) {
            const rdata = selectedData[i];
            fields.push(
                `<li class="collection-item">${rdata[View.SelectKey]}</li>`
            );
        }
        fields.push("</ul>");

        fieldParams.Items = selectedData;
    }

    if (View.Fields) {
        for (let i = 0, len = View.Fields.length; i < len; i++) {
            const field = View.Fields[i];
            let value = "";
            if (field.DefaultFunc) {
                value = field.DefaultFunc(data.service.data);
            } else if (field.Default) {
                value = field.Default;
            }
            if (field.Key) {
                fieldParams[field.Key] = null;
            } else {
                fieldParams[field.Name] = null;
            }
            const fieldId = `${keyPrefix}field${i}`;
            switch (field.Kind) {
                case "Text":
                    fields.push(`
                    <div class="row">
                      <div class="input-field col s12">
                        <input type="text" id="${fieldId}" class="validate ${fieldTextClass}" data-field-idx="${i}" value="${value}"/>
                        <label for="${fieldId}">${field.Name}</label>
                        <span class="helper-text" data-error="wrong" data-success="right"></span>
                      </div>
                    </div>
                    `);
                    break;
                case "Texts":
                    fields.push(`
                    <div class="row">
                      <div class="input-field col s12">
                        <textarea id="${fieldId}" class="materialize-textarea ${fieldTextClass}" data-field-idx="${i}">${value}</textarea>
                        <label for="${fieldId}">${field.Name}</label>
                        <span class="helper-text" data-error="wrong" data-success="right"></span>
                      </div>
                    </div>
                    `);
                    break;
                case "Password":
                    fields.push(`
                    <div class="row">
                      <div class="input-field col s12">
                        <input type="password" id="${fieldId}" class="validate ${fieldTextClass}" data-field-idx="${i}" value="${value}"/>
                        <label for="${fieldId}">${field.Name}</label>
                        <span class="helper-text" data-error="wrong" data-success="right"></span>
                      </div>
                    </div>
                    `);
                    break;
                case "Checkbox":
                    fields.push(`
                    <div class="row">
                      <div class="col s12">
                        <label>
                          <input type="checkbox" class="${fieldCheckClass}" data-field-idx="${i}" />
                          <span>Remember me</span>
                        </label>
                      </div>
                    </div>
                    `);
                    break;
                default:
                    fields.push(`UnknownField: ${field.Kind}`);
                    break;
            }
        }
    }

    function validateInputField(input: any) {
        const { elem } = input;
        const dataFieldIdx = elem.attr("data-field-idx");
        if (!dataFieldIdx) {
            return;
        }
        const field = View.Fields[parseInt(dataFieldIdx)];
        let val = elem.val();
        if (!val) {
            if (field.Default) {
                val = field.Default;
            } else {
                val = "";
            }
        }
        const len = val.length;
        let error = "";
        if (field.Required && len === 0) {
            error += "Please enter.";
        }
        if (field.Min && len < field.Min) {
            error += `Please enter ${field.Min} or more charactors. `;
        }
        if (field.Max && len > field.Max) {
            error += `Please enter ${field.Max} or less charactors. `;
        }

        if (field.RegExp) {
            const re = new RegExp(field.RegExp);
            if (val.length > 0 && !re.test(val)) {
                if (field.RegExpMsg) {
                    error += field.RegExpMsg + " ";
                } else {
                    error += "Invalid characters. ";
                }
            }
        }

        elem.parent().find(".helper-text").attr("data-error", error);
        let key = field.Key;
        if (!key) {
            key = field.Name;
        }
        if (error !== "") {
            elem.removeClass("valid").addClass("invalid");
            fieldParams[key] = null;
        } else {
            elem.removeClass("invalid").addClass("valid");
            fieldParams[key] = elem.val();
        }
    }

    function validateCheckField(input: any) {
        const { elem } = input;
        const dataFieldIdx = elem.attr("data-field-idx");
        if (!dataFieldIdx) {
            return;
        }
        const field = View.Fields[parseInt(dataFieldIdx)];
        let key = field.Key;
        if (!key) {
            key = field.Name;
        }
        fieldParams[key] = elem[0].checked;
    }

    $(`#${id}`).html(`
    <h1>${View.Name}</h1>
    <form class="col s12 form" id="${formId}">
      ${fields.join("")}
      <div id="${submitButtonWrapperId}"></div>

      <div id="${progressId}" class="progress">
        <div class="determinate" style="width: 0%"></div>
      </div>
    </form>
    `);

    function startProgress() {
        $(`#${progressId}`).html('<div class="indeterminate"></div>');
    }

    function stopProgress() {
        $(`#${progressId}`).html(
            '<div class="determinate" style="width: 0%"></div>'
        );
    }

    $(`.${fieldTextClass}`)
        .off("change")
        .on("change", function () {
            validateInputField({ elem: $(this) });
        })
        .off("keyup")
        .on("keyup", function () {
            validateInputField({ elem: $(this) });
        });

    function onSubmitInternal(e: any) {
        e.preventDefault();

        const inputs = $(`.${fieldTextClass}`);
        for (let i = 0, len = inputs.length; i < len; i++) {
            validateInputField({ elem: $(inputs[i]) });
        }

        const checks = $(`.${fieldCheckClass}`);
        for (let i = 0, len = checks.length; i < len; i++) {
            validateCheckField({ elem: $(checks[i]) });
        }

        let isValid = true;
        for (const value of Object.values(fieldParams)) {
            if (value === null) {
                isValid = false;
            }
        }
        if (!isValid) {
            console.log("debug submit");
            return;
        }

        if (onSubmit) {
            startProgress();
            onSubmit({
                params: fieldParams,
                onSuccess: function () {
                    stopProgress();
                    Dashboard.RootModal.Close();
                }
            });
        } else {
            startProgress();
            service.submitQueries({
                queries: [View.Action],
                location: location,
                params: fieldParams,
                onSuccess: function () {
                    stopProgress();
                    input.onSuccess();
                },
                onError: function () {
                    stopProgress();
                }
            });
        }
    }

    $(`#${formId}`).on("submit", onSubmitInternal);
    if (useRootModal) {
        Dashboard.RootModal.Init({ View, onSubmit: onSubmitInternal });
        Dashboard.RootModal.Open();
    } else {
        let submitButtonName = "Submit";
        if (View.SubmitButtonName) {
            submitButtonName = View.SubmitButtonName;
        }
        let icon = "";
        if (View.Icon) {
            icon = Icon.Html({ kind: View.Icon });
        }
        let buttonClass = "";
        let style = "";
        if (View.SubmitButtonStyle) {
            if (View.SubmitButtonStyle == "wide") {
                buttonClass = "btn-large";
                style = "width: 100%;";
            }
        }
        $(`#${submitButtonWrapperId}`).html(`
        <div class="row">
          <div class="input-field col s12">
            <button type="submit" style="display: none;"></button>
            <a id="${submitButtonId}" class="waves-effect waves-light btn btn-primary btn-block ${buttonClass}" style="${style}">${icon}${submitButtonName}</a>
          </div>
        </div>
        `);
        $(`#${submitButtonId}`).on("click", function (e: any) {
            e.preventDefault();
            onSubmitInternal(e);
        });
    }

    M.updateTextFields();
    if (View.Fields) {
        for (let i = 0, len = View.Fields.length; i < len; i++) {
            const field = View.Fields[i];
            const fieldId = `${keyPrefix}field${i}`;
            switch (field.Kind) {
                case "Texts":
                    M.textareaAutoResize($(`#${fieldId}`));
                    break;
            }
            if (i === 0) {
                $(`#${fieldId}`).focus();
            }
        }
    }
}

const index = {
    Render
};
export default index;
