import Panes from "./panes/Panes";
import Panels from "./panels/Panels";
import Table from "./table/Table";
import Notfound from "./core/Notfound";

function Render(input: any) {
    const { View } = input;
    console.log("Render", input);
    switch (View.Kind) {
        case "Panes":
            return Panes.Render(input);
        case "Panels":
            return Panels.Render(input);
        case "Table":
            return Table.Render(input);
        default:
            return Notfound.Render(input);
    }
}

const index = {
    Render
};
export default index;
