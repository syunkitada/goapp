import data from "../../data";

export function Render(input: any) {
    const { id, View } = input;
    console.log("DEBUG Panels.Render", input, View, data);

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
