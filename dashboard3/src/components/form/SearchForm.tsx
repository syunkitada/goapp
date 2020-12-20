export function Render(input: any) {
    const { id, View } = input;
    const keyPrefix = `${id}-SearchForm-`;
    const inputClass = `${keyPrefix}input`;
    const selectClass = `${keyPrefix}select`;

    const inputs: any[] = [];
    for (let i = 0, len = View.Inputs.length; i < len; i++) {
        const input = View.Inputs[i];
        let defaultValue: any;
        switch (input.Type) {
            case "Select":
                console.log("DEBUG input", input);
                let options = [];
                for (let i = 0, len = input.Data.length; i < len; i++) {
                    const option = input.Data[i];
                    options.push(
                        `<option value="${option}">${option}</option>`
                    );
                }
                const currentOption = input.Default;
                inputs.push(`
                <div class="input-field col m2">
                    <select name="${input.Name}" class="${selectClass}">
                        <option value="${currentOption}" disabled selected>${currentOption}</option>
                        ${options.join("")}
                    </select>
                    <label>${input.Name}</label>
                </div>
                `);
                break;
            case "Text":
                inputs.push(`
                <div class="input-field col m6">
                    <input id="first_name2" class="${inputClass}" type="text" class="validate">
                        <label for="first_name2">First Name</label>
                    </input>
                </div>
                `);
                break;
            case "DateTime":
                inputs.push(`
                <div class="input-field col m2">
                    <input type="text" class="${inputClass} datepicker" placeholder="Date">
                </div>
                <div class="input-field col m2">
                    <input type="text" class="${inputClass} timepicker" placeholder="Time">
                </div>
                `);
                break;
            default:
                inputs.push(`<span>UnknownInput: ${input.Type}</span>`);
        }
    }

    const formId = `${keyPrefix}form`;
    const searchButtonId = `${keyPrefix}searchInput`;

    $(`#${id}`).html(`
        <form class="col s12" id="${formId}">
            <div class="row">${inputs.join("")}
                <div class="input-field col m2">
                <a class="waves-effect waves-light btn" id="${searchButtonId}"><i class="material-icons right">search</i>Search</a>
                </div>
            </div>
        </form>
    `);

    $(`.${selectClass}`).formSelect();

    $(`#${searchButtonId}`).on("click", function () {
        const inputs = $(`.${inputClass}`);
        const selects = $(`.${selectClass}`);
        for (let i = 0, len = inputs.length; i < len; i++) {
            const input = $(inputs[i]);
            const val = input.val();
            console.log("DEBUG input val", val);
        }
        for (let i = 0, len = selects.length; i < len; i++) {
            const select = $(selects[i]);
            const val = select.val();
            console.log("DEBUG select val", val);
        }
    });

    $(".datepicker").datepicker({
        autoClose: true,
        onSelect: function (e: any) {
            console.log("DEBUG onCloseEnd datepicker", e);
        }
    });
    $(".timepicker").timepicker({
        autoClose: true,
        onCloseEnd: function (e: any) {
            console.log("DEBUG onCloseEnd datepicker", e);
        }
    });
}

const index = {
    Render
};
export default index;
