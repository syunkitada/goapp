import Panels from "./panels/Panels";
import Notfound from "./core/Notfound";

function Render(input: any) {
    const { id, View } = input;
    switch (View.Kind) {
        case "Panels":
            return Panels.Render(input);
        default:
            return Notfound.Render(input);
    }
    console.log("Render", input);
}

const index = {
    Render
};
export default index;
