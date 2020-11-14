import data from "../../data";
import data_utils from "../../lib/data_utils";

export function Render(input: any) {
    const { id, View } = input;
    console.log("DEBUG Panels.Render", input, View);

    // const locationData = data_utils.getLocationData();

    $(`#${id}`).html(`
    <div>
    table
    </div>
    `);
}

const index = {
    Render
};
export default index;
