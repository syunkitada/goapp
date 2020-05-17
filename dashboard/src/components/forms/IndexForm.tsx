import * as React from "react";
import { connect } from "react-redux";

import { Theme } from "@material-ui/core/styles/createMuiTheme";
import createStyles from "@material-ui/core/styles/createStyles";
import withStyles, {
    StyleRules,
    WithStyles
} from "@material-ui/core/styles/withStyles";

import Button from "@material-ui/core/Button";
import CircularProgress from "@material-ui/core/CircularProgress";
import DialogActions from "@material-ui/core/DialogActions";
import DialogContent from "@material-ui/core/DialogContent";
import DialogContentText from "@material-ui/core/DialogContentText";
import DialogTitle from "@material-ui/core/DialogTitle";
import Grid from "@material-ui/core/Grid";
// import TextareaAutosize from '@material-ui/core/TextareaAutosize';
import TextField from "@material-ui/core/TextField";

import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText";
import DoneIcon from "@material-ui/icons/Done";

import green from "@material-ui/core/colors/green";
import red from "@material-ui/core/colors/red";

import actions from "../../actions";
import data_utils from "../../lib/data_utils";
import logger from "../../lib/logger";

import Icon from "../icons/Icon";

interface IIndexForm extends WithStyles<typeof styles> {
    targets;
    routes;
    data;
    indexData;
    selected;
    index;
    onClose;
    isSubmitting;
    title;
    service;
    submitQueries;
    submitSuccessMsg;
}

class IndexForm extends React.Component<IIndexForm> {
    public state = {
        fieldMap: {}
    };

    public render() {
        const {
            classes,
            index,
            selected,
            isSubmitting,
            submitSuccessMsg,
            title,
            onClose
        } = this.props;
        logger.info("IndexForm", "render", index, selected);

        if (submitSuccessMsg) {
            return (
                <div>
                    {title && (
                        <DialogTitle id="form-dialog-title">
                            {title}
                        </DialogTitle>
                    )}
                    <DialogContent className={classes.dialogContent}>
                        <DialogContentText>
                            {index.Description}
                        </DialogContentText>
                        {submitSuccessMsg}
                    </DialogContent>
                    <DialogActions>
                        <div
                            className={classes.wrapper}
                            style={{ width: "100%" }}
                        >
                            <Grid container={true}>
                                <Grid
                                    container={true}
                                    item={true}
                                    xs={6}
                                    justify="flex-start"
                                >
                                    {onClose && (
                                        <Button
                                            onClick={onClose}
                                            disabled={isSubmitting}
                                        >
                                            Cancel
                                        </Button>
                                    )}
                                </Grid>
                                <Grid
                                    container={true}
                                    item={true}
                                    xs={6}
                                    justify="flex-end"
                                >
                                    {isSubmitting ? (
                                        <CircularProgress
                                            size={24}
                                            className={classes.buttonProgress}
                                        />
                                    ) : (
                                        <Button
                                            variant="contained"
                                            color="primary"
                                            disabled={true}
                                            className={classes.button}
                                        >
                                            <Icon
                                                name={index.Icon}
                                                marginDirection={"right"}
                                            />
                                            Success
                                        </Button>
                                    )}
                                </Grid>
                            </Grid>
                        </div>
                    </DialogActions>
                </div>
            );
        }

        const fields = this.renderFields();

        return (
            <div>
                {title && (
                    <DialogTitle id="form-dialog-title">{title}</DialogTitle>
                )}
                <DialogContent className={classes.dialogContent}>
                    <DialogContentText>{index.Description}</DialogContentText>
                    {fields}
                </DialogContent>
                <DialogActions>
                    <div className={classes.wrapper} style={{ width: "100%" }}>
                        <Grid container={true}>
                            <Grid
                                container={true}
                                item={true}
                                xs={6}
                                justify="flex-start"
                            >
                                {onClose && (
                                    <Button
                                        onClick={onClose}
                                        disabled={isSubmitting}
                                    >
                                        Cancel
                                    </Button>
                                )}
                            </Grid>
                            <Grid
                                container={true}
                                item={true}
                                xs={6}
                                justify="flex-end"
                            >
                                {isSubmitting ? (
                                    <CircularProgress
                                        size={24}
                                        className={classes.buttonProgress}
                                    />
                                ) : (
                                    <Button
                                        variant="contained"
                                        color="primary"
                                        disabled={isSubmitting}
                                        className={classes.button}
                                        onClick={this.handleActionSubmit}
                                    >
                                        <Icon
                                            name={index.Icon}
                                            marginDirection={"right"}
                                        />
                                        {index.SubmitButtonName}
                                    </Button>
                                )}
                            </Grid>
                        </Grid>
                    </div>
                </DialogActions>
            </div>
        );
    }

    private renderFields = () => {
        const { classes, data, selected, index, indexData } = this.props;
        const { fieldMap } = this.state;
        const fields: JSX.Element[] = [];

        if (selected) {
            const items = data[index.DataKey];
            const listItems: JSX.Element[] = [];
            for (let i = 0, len = selected.length; i < len; i++) {
                const s = selected[i];
                const item = items[s];
                const value = item[index.SelectKey];

                listItems.push(
                    <ListItem key={s}>
                        <ListItemIcon>
                            <DoneIcon />
                        </ListItemIcon>
                        <ListItemText primary={value} />
                    </ListItem>
                );
            }

            fields.push(<List key={"selected"}>{listItems}</List>);
        }

        if (!index.Fields) {
            return fields;
        }

        for (let i = 0, len = index.Fields.length; i < len; i++) {
            const field = index.Fields[i];
            let isError = false;
            let helperText = "";

            let autoFocus = false;
            if (i === 0) {
                autoFocus = true;
            }
            let disabled = false;

            let value = "";
            const fieldState = fieldMap[field.Name];
            if (fieldState) {
                if (fieldState.error !== "") {
                    isError = true;
                    helperText = fieldState.error;
                }
                value = fieldState.value;
            } else {
                if (indexData) {
                    value = indexData[field.Name];
                    if (!field.Updatable) {
                        disabled = true;
                    }
                } else if (field.DefaultFunc) {
                    value = field.DefaultFunc(data);
                }
            }

            switch (field.Kind) {
                case "text":
                    fields.push(
                        <TextField
                            id={field.Name}
                            key={i}
                            disabled={disabled}
                            label={field.Name}
                            autoFocus={autoFocus}
                            margin="dense"
                            type={field.Kind}
                            fullWidth={true}
                            onChange={event => {
                                this.handleTextFieldChange(event, field);
                            }}
                            value={value}
                            helperText={helperText}
                            error={isError}
                        />
                    );
                    break;

                case "texts":
                    fields.push(
                        <TextField
                            id={field.Name}
                            key={i}
                            disabled={disabled}
                            label={field.Name}
                            multiline={true}
                            rows={3}
                            autoFocus={autoFocus}
                            margin="dense"
                            type={field.Kind}
                            fullWidth={true}
                            onChange={event => {
                                this.handleTextFieldChange(event, field);
                            }}
                            value={value}
                            helperText={helperText}
                            error={isError}
                        />
                    );
                    break;

                case "password":
                    fields.push(
                        <TextField
                            id={field.Name}
                            key={i}
                            disabled={disabled}
                            label={field.Name}
                            autoFocus={autoFocus}
                            margin="dense"
                            type="password"
                            fullWidth={true}
                            onChange={event => {
                                this.handleTextFieldChange(event, field);
                            }}
                            value={value}
                            helperText={helperText}
                            error={isError}
                        />
                    );
                    break;

                case "select":
                    let options = field.Options;
                    if (!options) {
                        options = [];
                        const d = data[field.DataKey];
                        if (d) {
                            for (let j = 0, l = d.length; j < l; j++) {
                                options.push(d[j].Name);
                            }
                        } else {
                            options.push("");
                        }
                    }
                    if (!value || value === "") {
                        value = options[0];
                    }

                    fields.push(
                        <TextField
                            select={true}
                            key={i}
                            label={field.Name}
                            className={classes.textField}
                            disabled={disabled}
                            value={value}
                            onChange={event => {
                                this.handleSelectFieldChange(event, field);
                            }}
                            SelectProps={{
                                MenuProps: {
                                    className: classes.menu
                                },
                                native: true
                            }}
                            helperText="Please select"
                            margin="normal"
                            fullWidth={true}
                        >
                            {options.map(option => (
                                <option key={option} value={option}>
                                    {option}
                                </option>
                            ))}
                        </TextField>
                    );
                    break;
                default:
                    fields.push(<span>FieldNotFound</span>);
                    break;
            }
        }

        return fields;
    };

    private handleTextFieldChange = (event, field) => {
        const { fieldMap } = this.state;

        const text = event.target.value;
        let error = "";
        const len = text.length;

        if (field.Min) {
            if (len < field.Min) {
                error += `Please enter ${field.Min} or more charactors. `;
            } else if (len > field.Max) {
                error += `Please enter ${field.Max} or less charactors. `;
            }
        }

        if (field.RegExp) {
            const re = new RegExp(field.RegExp);
            if (!re.test(text)) {
                if (field.RegExpMsg) {
                    error += field.RegExpMsg + " ";
                } else {
                    error += "Invalid characters. ";
                }
            }
        }

        fieldMap[field.Name] = {
            error,
            type: field.Kind,
            value: event.target.value
        };

        this.setState({ fieldMap });
    };

    private handleSelectFieldChange = (event, field) => {
        const { fieldMap } = this.state;
        fieldMap[field.Name] = {
            error: null,
            type: field.Kind,
            value: event.target.value
        };
        this.setState({ fieldMap });
    };

    private handleActionSubmit = () => {
        const {
            index,
            data,
            indexData,
            selected,
            routes,
            submitQueries
        } = this.props;
        const { fieldMap } = this.state;
        const route = routes.slice(-1)[0];
        const fieldData = {};

        if (index.Fields) {
            console.log("DEBUG index.Fields", index.Fields);
            // Validate
            // フォーム入力がなく、デフォルト値がある場合はセットする
            for (let i = 0, len = index.Fields.length; i < len; i++) {
                const field = index.Fields[i];

                const fieldState = fieldMap[field.Name];
                let value = "";
                if (fieldState) {
                    value = fieldState.value;
                } else {
                    if (indexData) {
                        value = indexData[field.Name];
                    }
                }

                switch (field.Kind) {
                    case "text":
                    case "texts":
                        if (field.Require) {
                            if (!value || value === "") {
                                fieldMap[field.Name] = {
                                    error: "This is required",
                                    type: field.Kind,
                                    value: ""
                                };
                            }
                            fieldData[field.Name] = { value };
                        } else {
                            if (value) {
                                fieldData[field.Name] = { value };
                            } else {
                                fieldData[field.Name] = {
                                    value: field.Default
                                };
                            }
                        }
                        break;

                    case "select":
                        if (!value) {
                            let options = field.Options;
                            if (!options) {
                                options = [];
                                const d = data[field.DataKey];
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
                        fieldData[field.Name] = { value };
                        break;
                    default:
                        break;
                }
            }
        }

        const items: any[] = [];
        if (index.SelectKey && selected) {
            const tmpItems = data[index.DataKey];
            for (let i = 0, len = selected.length; i < len; i++) {
                items.push(tmpItems[selected[i]]);
            }
            if (items.length === 0) {
                // TODO handle unknown error
                return;
            }
        } else {
            items.push({});
        }

        for (const key in fieldMap) {
            if (fieldMap[key].error && fieldMap[key].error !== "") {
                this.setState({ fieldMap });
                return;
            }
        }

        submitQueries(route, items, fieldData);
    };
}

function mapStateToProps(state, ownProps) {
    const { service } = state;
    const { isSubmitting, submitSuccessMsg } = state.service;

    const data = data_utils.getDataFromState(state);
    const indexData = data[ownProps.index.DataKey];
    return {
        data,
        indexData,
        isSubmitting,
        service,
        submitSuccessMsg
    };
}

function mapDispatchToProps(dispatch, ownProps) {
    const { index } = ownProps;
    return {
        submitQueries: (route, items, fieldMap) => {
            dispatch(
                actions.service.serviceSubmitQueries({
                    fieldMap,
                    index,
                    items,
                    location: route
                })
            );
        }
    };
}

const styles = (theme: Theme): StyleRules =>
    createStyles({
        button: {
            margin: theme.spacing(1)
        },
        buttonFailed: {
            "&:hover": {
                backgroundColor: red[700]
            },
            backgroundColor: red[500]
        },
        buttonProgress: {
            color: green[500],
            left: "50%",
            marginLeft: -12,
            marginTop: -12,
            position: "absolute",
            top: "50%"
        },
        buttonSuccess: {
            "&:hover": {
                backgroundColor: green[700]
            },
            backgroundColor: green[500]
        },
        dialogContent: {
            minWidth: 400
        },
        fabProgress: {
            color: green[500],
            left: -6,
            position: "absolute",
            top: -6,
            zIndex: 1
        },
        root: {
            width: "100%"
        },
        wrapper: {
            margin: theme.spacing(1),
            position: "relative"
        }
    });

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(withStyles(styles, { withTheme: true })(IndexForm));
