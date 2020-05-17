import * as React from "react";

import { Theme } from "@material-ui/core/styles/createMuiTheme";
import createStyles from "@material-ui/core/styles/createStyles";
import withStyles, {
    StyleRules,
    WithStyles
} from "@material-ui/core/styles/withStyles";

import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableRow from "@material-ui/core/TableRow";

interface IField extends WithStyles<typeof styles> {
    field;
    data;
}

class Field extends React.Component<IField> {
    public render() {
        const { field, data, classes } = this.props;
        let html: any;
        let subData: any;
        if (!data) {
            return (
                <TableRow>
                    <TableCell>{field.Name}</TableCell>
                    <TableCell style={{ width: "100%" }}></TableCell>
                </TableRow>
            );
        }
        if (field.DataKey) {
            subData = data[field.DataKey];
        } else {
            subData = data[field.Name];
        }
        switch (field.Kind) {
            case "List":
                const fields: JSX.Element[] = [];
                for (let i = 0, len = subData.length; i < len; i++) {
                    const subDataItem = subData[i];
                    const subFields: JSX.Element[] = [];
                    for (let j = 0, lenj = field.Fields.length; j < lenj; j++) {
                        const subField = field.Fields[j];
                        subFields.push(
                            <Field
                                key={j}
                                classes={classes}
                                field={subField}
                                data={subDataItem}
                            />
                        );
                    }
                    const key = subDataItem[field.ListKey];
                    fields.push(
                        <Table key={key} size="small">
                            <TableBody>
                                <TableRow>
                                    <TableCell>{key}</TableCell>
                                    <TableCell>
                                        <Table size="small">
                                            <TableBody>{subFields}</TableBody>
                                        </Table>
                                    </TableCell>
                                </TableRow>
                            </TableBody>
                        </Table>
                    );
                }
                html = <div>{fields}</div>;
                break;
            default:
                html = <span>{subData}</span>;
                break;
        }
        return (
            <TableRow>
                <TableCell>{field.Name}</TableCell>
                <TableCell style={{ width: "100%" }}>{html}</TableCell>
            </TableRow>
        );
    }
}

const styles = (theme: Theme): StyleRules =>
    createStyles({
        root: {
            width: "100%"
        }
    });

export default withStyles(styles, { withTheme: true })(Field);
