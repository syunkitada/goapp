export function Render(input: any) {
    const { id } = input;
    console.log("DEBUG Panels.Render", input);

    $(`#${id}`).html(`
    Panels
    `);
}

const index = {
    Render
};
export default index;
