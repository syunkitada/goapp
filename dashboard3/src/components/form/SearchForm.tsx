export function Render(input: any) {
    const { id, View } = input;

    const inputs: any[] = [];
    for (let i = 0, len = View.Inputs.length; i < len; i++) {
        const input = View.Inputs[i];
        let defaultValue: any;
        switch (input.Type) {
            case "Select":
                inputs.push(`
                <div class="input-field col m6">
                    <input id="first_name" type="text" class="validate">
                        <label for="first_name">First Name</label>
                    </input>
                </div>
                `);
                break;
            case "Text":
                inputs.push(`
                <div class="input-field col m6">
                    <input id="first_name2" type="text" class="validate">
                        <label for="first_name2">First Name</label>
                    </input>
                </div>
                `);
                break;
            case "DateTime":
                inputs.push(`
                <div class="input-field col m3">
                    <input type="text" class="datepicker" placeholder="Date">
                </div>
                <div class="input-field col m3">
                    <input type="text" class="timepicker" placeholder="Time">
                </div>
                `);
                break;
            default:
                inputs.push(`<span>UnknownInput: ${input.Type}</span>`);
        }
    }

    $(`#${id}`).html(`
        <form class="col s12">
        <div class="row">${inputs.join("")}</div>
        </form>
    `);

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
