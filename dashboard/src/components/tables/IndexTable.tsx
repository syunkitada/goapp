import * as React from "react";
import { connect } from "react-redux";

import { Theme } from "@material-ui/core/styles/createMuiTheme";
import createStyles from "@material-ui/core/styles/createStyles";
import withStyles, {
    StyleRules,
    WithStyles
} from "@material-ui/core/styles/withStyles";

import Button from "@material-ui/core/Button";
import Checkbox from "@material-ui/core/Checkbox";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText";
import Menu from "@material-ui/core/Menu";
import MenuItem from "@material-ui/core/MenuItem";
import Paper from "@material-ui/core/Paper";
import Popover from "@material-ui/core/Popover";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableRow from "@material-ui/core/TableRow";

import Grid from "@material-ui/core/Grid";

import KeyboardArrowDownIcon from "@material-ui/icons/KeyboardArrowDown";

import FormDialog from "../dialogs/FormDialog";
import IndexTableHead from "./IndexTableHead";
import TableToolbar from "./TableToolbar";

import Badge from "@material-ui/core/Badge";
import IconButton from "@material-ui/core/IconButton";
import LinearProgress from "@material-ui/core/LinearProgress";

import actions from "../../actions";
import logger from "../../lib/logger";
import sort_utils from "../../lib/sort_utils";
import theme_utils from "../../lib/theme_utils";

import Tooltip from "@material-ui/core/Tooltip";
import Icon from "../icons/Icon";

import SearchForm from "../forms/SearchForm";

import data_utils from "../../lib/data_utils";

interface IIndexTable extends WithStyles<typeof styles> {
    actionName;
    auth;
    render;
    routes;
    index;
    data;
    subData;
    handleColumnLinkClick;
    setAction;
    resetAction;
    rootIndex;
}

interface IState {
    order;
    orderBy;
    selected: any[];
    data: any;
    subData: any;
    page;
    popoverHtml;
    popoverTarget;
    rowsPerPage;
    searchQueries: any;
    searchRegExp: any;
    anchorEl: any;
    filterMap: any;
}

class IndexTable extends React.Component<IIndexTable> {
    public state: IState = {
        anchorEl: null,
        data: [],
        subData: [],
        filterMap: {},
        order: "asc",
        orderBy: -1,
        page: 0,
        popoverHtml: null,
        popoverTarget: null,
        rowsPerPage: 0,
        searchQueries: {},
        searchRegExp: null,
        selected: []
    };

    public render() {
        const {
            auth,
            actionName,
            routes,
            resetAction,
            subData,
            classes,
            index
        } = this.props;
        let { data } = this.props;
        const {
            popoverTarget,
            popoverHtml,
            selected,
            anchorEl,
            rowsPerPage,
            page,
            searchRegExp,
            filterMap
        } = this.state;
        if (!data) {
            if (!subData) {
                logger.warning("Invalid data", index, data);

                return (
                    <Paper key={index.Name} className={classes.root}>
                        <div className={classes.tableWrapper}>
                            <br />
                            <LinearProgress color="secondary" />
                        </div>
                    </Paper>
                );
            }
            data = subData;
        }
        console.log("DEBUGRENDER", "IndexTable", performance.now(), index);

        logger.info("IndexTable", "render", actionName, routes);
        console.log("DEBUG Table", index.DataKey, index.Columns, data);
        let tmpSelected: any = null;
        if (resetAction) {
            tmpSelected = [];
        } else {
            tmpSelected = selected;
        }

        let tmpRowsPerPage = 5;
        if (rowsPerPage > 0) {
            tmpRowsPerPage = rowsPerPage;
        } else if (index.RowsPerPage) {
            tmpRowsPerPage = index.RowsPerPage;
        }

        console.log(auth);
        const exButtons: any[] = [];
        const columns = index.Columns;

        let isSelectActions = false;
        if (index.SelectActions) {
            isSelectActions = true;
        }

        const searchColumns: any[] = [];
        let order = "asc";
        let orderBy = -1;
        const filterCounterMap = {};
        for (let i = 0, len = columns.length; i < len; i++) {
            const column = columns[i];
            if (column.IsSearch) {
                searchColumns.push(column.Name);
            }
            if (column.Sort) {
                order = column.Sort;
                orderBy = i;
            }
            if (column.FilterValues) {
                for (
                    let j = 0, lenj = column.FilterValues.length;
                    j < lenj;
                    j++
                ) {
                    const filterValue = column.FilterValues[j];
                    filterCounterMap[filterValue.Value] = 0;
                }
            }
        }

        let isSkip = true;
        const tableData: any[] = [];
        for (let i = 0, len = data.length; i < len; i++) {
            const d = data[i];
            if (searchRegExp != null) {
                for (const c of searchColumns) {
                    if (searchRegExp.exec(d[c])) {
                        isSkip = false;
                        break;
                    }
                }
                if (isSkip) {
                    continue;
                }
                isSkip = true;
            }

            const row: any[] = [i];
            for (const column of columns) {
                let c: any;
                if (column.Key) {
                    c = d[column.Key];
                } else {
                    c = d[column.Name];
                }
                if (column.Kind === "Time") {
                    if (c) {
                        const time: any = new Date(c);
                        if (!isNaN(time.getTime())) {
                            row.push(time.toISOString());
                        } else {
                            row.push(time.toString());
                        }
                    } else {
                        row.push("");
                    }
                } else if (column.Kind === "Action") {
                    row.push("");
                } else if (column.Kind === "Popover") {
                    if (column.View) {
                        const popoverData = {
                            Data: d,
                            Value: c
                        };
                        row.push(popoverData);
                    } else {
                        const popoverData = {
                            Data: null,
                            Value: c
                        };
                        row.push(popoverData);
                    }
                } else {
                    row.push(c);
                }

                if (column.FilterValues) {
                    for (
                        let j = 0, lenj = column.FilterValues.length;
                        j < lenj;
                        j++
                    ) {
                        const filterValue = column.FilterValues[j];
                        if (filterValue.Value === c) {
                            filterCounterMap[filterValue.Value] += 1;
                        }
                    }
                }
            } // for (const column of columns) {

            tableData.push(row);
        }
        console.log("DEBUG counterMap", filterCounterMap);

        for (let i = 0, l = columns.length; i < l; i++) {
            columns[i].id = i;
        }
        columns[0].disablePadding = false;

        const columnActions: any[] = [];
        if (index.ColumnActions != null) {
            for (let i = 0, len = index.ColumnActions.length; i < len; i++) {
                const columnAction = index.ColumnActions[i];
                columnActions.push(
                    <MenuItem
                        key={columnAction.Name}
                        onClick={event =>
                            this.handleActionClick(event, columnAction.Name)
                        }
                    >
                        <ListItemIcon>
                            <Icon name={columnAction.Icon} />
                        </ListItemIcon>
                        <ListItemText
                            inset={true}
                            primary={columnAction.Name}
                        />
                    </MenuItem>
                );
            }
        }

        let actionDialog: any = null;
        if (actionName !== null) {
            let action: any = null;
            if (index.Actions) {
                for (const a of index.Actions) {
                    if (a.Name === actionName) {
                        action = a;
                        break;
                    }
                }
            }

            if (action === null) {
                if (index.SelectActions) {
                    for (const a of index.SelectActions) {
                        if (a.Name === actionName) {
                            action = a;
                            break;
                        }
                    }
                }
            }

            if (action === null) {
                if (index.ColumnActions) {
                    for (const a of index.ColumnActions) {
                        if (a.Name === actionName) {
                            action = a;
                            break;
                        }
                    }
                }
            }

            if (action === null) {
                actionDialog = null;
            } else {
                switch (action.Kind) {
                    case "Form":
                        actionDialog = (
                            <FormDialog
                                open={true}
                                data={data}
                                selected={tmpSelected}
                                action={action}
                                routes={routes}
                                onClose={this.handleActionDialogClose}
                            />
                        );
                        break;
                    default:
                        actionDialog = null;
                        break;
                }
            }
        }

        for (let i = 0, len = columns.length; i < len; i++) {
            const column = columns[i];
            if (column.FilterValues) {
                for (
                    let j = 0, lenj = column.FilterValues.length;
                    j < lenj;
                    j++
                ) {
                    const filterValue = column.FilterValues[j];

                    const currentValue = filterMap[column.Name];
                    let opacity = 1;
                    if (currentValue && currentValue === filterValue.Value) {
                        opacity = 0.7;
                    }

                    exButtons.push(
                        <Tooltip key={j} title={filterValue.Value}>
                            <IconButton
                                aria-label={column.Name}
                                className={classes.toolbarExButton}
                                style={{ opacity }}
                                onClick={() =>
                                    this.handleFilterClick(
                                        column.Name,
                                        filterValue.Value
                                    )
                                }
                                value={filterValue.Value}
                            >
                                <Badge
                                    badgeContent={
                                        filterCounterMap[filterValue.Value]
                                    }
                                    color="primary"
                                >
                                    <Icon
                                        name={filterValue.Icon}
                                        style={{
                                            color: theme_utils.getFgColor(
                                                auth.theme,
                                                filterValue.Icon
                                            )
                                        }}
                                        marginDirection="left"
                                    />
                                </Badge>
                            </IconButton>
                        </Tooltip>
                    );
                }
            }
        }

        let searchForm: any = null;
        if (index.SearchForm) {
            searchForm = (
                <Grid item={true} xs={12}>
                    <SearchForm
                        routes={routes}
                        index={index.SearchForm}
                        data={data}
                        onChange={this.handleChangeOnSearchForm}
                        onSubmit={this.handleSubmitOnSearchForm}
                    />
                </Grid>
            );
        }

        const indexLength = tableData.length;

        let emptyRows: number = 0;
        let cellItems: any;
        if (orderBy >= 0) {
            cellItems = sort_utils.stableSort(
                tableData,
                sort_utils.getSorting(order, orderBy)
            );
        } else {
            cellItems = tableData;
        }
        if (index.DisablePaging) {
            if (cellItems.length === 0) {
                emptyRows = 1;
            }
        } else {
            cellItems = cellItems.slice(
                page * tmpRowsPerPage,
                page * tmpRowsPerPage + tmpRowsPerPage
            );
            if (index.FixHeight) {
                emptyRows =
                    tmpRowsPerPage -
                    Math.min(
                        tmpRowsPerPage,
                        indexLength - page * tmpRowsPerPage
                    );
            } else {
                emptyRows = 1;
            }
        }

        return (
            <Paper key={index.Name} className={classes.root}>
                {!index.DisableToolbar && (
                    <TableToolbar
                        count={indexLength}
                        rowsPerPage={tmpRowsPerPage}
                        page={page}
                        index={index}
                        numSelected={tmpSelected.length}
                        onChangePage={this.handleChangePage}
                        onChangeRowsPerPage={this.handleChangeRowsPerPage}
                        onChangeSearchInput={this.handleChangeSearchInput}
                        onActionClick={this.handleActionClick}
                        exButtons={exButtons}
                        searchForm={searchForm}
                    />
                )}
                <div className={classes.tableWrapper}>
                    <Table
                        className={classes.table}
                        aria-labelledby="tableTitle"
                        size="small"
                    >
                        <IndexTableHead
                            index={index}
                            order={order}
                            orderBy={orderBy}
                            onSelectAllClick={this.handleSelectAllClick}
                            onRequestSort={this.handleRequestSort}
                            rowCount={indexLength}
                            columns={columns}
                            numSelected={tmpSelected.length}
                        />
                        <TableBody>
                            {cellItems.map((n, rowIndex) => {
                                const cells: any = [];
                                const key = n[0];
                                const isSelected = this.isSelected(key);

                                if (isSelectActions) {
                                    cells.push(
                                        <TableCell
                                            key={-1}
                                            padding="checkbox"
                                            onClick={event =>
                                                this.handleSelectClick(
                                                    event,
                                                    key
                                                )
                                            }
                                        >
                                            <Checkbox checked={isSelected} />
                                        </TableCell>
                                    );
                                }

                                for (
                                    let i = 0, len = columns.length;
                                    i < len;
                                    i++
                                ) {
                                    const column = columns[i];
                                    const cellValue = n[i + 1];
                                    let padding: any = "default";
                                    let align: any = "right";
                                    if (column.Padding) {
                                        padding = column.Padding;
                                    }
                                    if (column.Align) {
                                        align = column.Align;
                                    }

                                    if (column.LinkPath) {
                                        cells.push(
                                            <TableCell
                                                align={align}
                                                key={i}
                                                component="th"
                                                padding={padding}
                                                scope="row"
                                                style={{ cursor: "pointer" }}
                                                onClick={() => {
                                                    this.props.handleColumnLinkClick(
                                                        index,
                                                        column,
                                                        cellValue
                                                    );
                                                }}
                                            >
                                                {cellValue}
                                            </TableCell>
                                        );
                                    } else if (column.Kind === "Action") {
                                        cells.push(
                                            <TableCell
                                                key={i}
                                                align={align}
                                                padding={padding}
                                            >
                                                <Button
                                                    aria-owns={
                                                        anchorEl
                                                            ? "simple-menu"
                                                            : undefined
                                                    }
                                                    aria-haspopup="true"
                                                    variant="outlined"
                                                    onClick={e => {
                                                        this.handleActionMenuClick(
                                                            e,
                                                            key
                                                        );
                                                    }}
                                                >
                                                    Actions{" "}
                                                    <KeyboardArrowDownIcon />
                                                </Button>
                                            </TableCell>
                                        );
                                    } else if (column.FilterValues) {
                                        let filterButton: any = null;
                                        const currentValue =
                                            filterMap[column.Name];
                                        let isShowCells = true;
                                        if (currentValue !== undefined) {
                                            isShowCells = false;
                                        }
                                        if (column.FilterValues) {
                                            for (
                                                let j = 0,
                                                    lenj =
                                                        column.FilterValues
                                                            .length;
                                                j < lenj;
                                                j++
                                            ) {
                                                const filterValue =
                                                    column.FilterValues[j];
                                                if (
                                                    filterValue.Value ===
                                                    cellValue
                                                ) {
                                                    if (
                                                        currentValue !==
                                                            undefined &&
                                                        currentValue ===
                                                            filterValue.Value
                                                    ) {
                                                        isShowCells = true;
                                                    }

                                                    let tmpValue = cellValue;
                                                    if (
                                                        column.Kind === "Hidden"
                                                    ) {
                                                        tmpValue = "";
                                                    }

                                                    filterButton = (
                                                        <IconButton
                                                            component="span"
                                                            style={{
                                                                padding: 0
                                                            }}
                                                        >
                                                            <Icon
                                                                key={j}
                                                                name={
                                                                    filterValue.Icon
                                                                }
                                                                style={{
                                                                    color: theme_utils.getFgColor(
                                                                        auth.theme,
                                                                        filterValue.Color
                                                                    )
                                                                }}
                                                            />
                                                            {tmpValue}
                                                        </IconButton>
                                                    );
                                                    break;
                                                }
                                            }
                                        }
                                        if (!isShowCells) {
                                            return null;
                                        }
                                        cells.push(
                                            <TableCell
                                                key={i}
                                                align={align}
                                                padding={padding}
                                            >
                                                {filterButton}
                                            </TableCell>
                                        );
                                    } else if (column.Kind === "Popover") {
                                        let tmpColor = column.Color;
                                        let isInactive = true;
                                        if (cellValue.Data) {
                                            const tmpData =
                                                cellValue.Data[
                                                    column.View.DataKey
                                                ];
                                            if (column.View.Kind === "Table") {
                                                if (
                                                    tmpData &&
                                                    tmpData.length > 0
                                                ) {
                                                    isInactive = false;
                                                }
                                            }
                                        }
                                        if (isInactive) {
                                            tmpColor = column.InactiveColor;
                                        }
                                        cells.push(
                                            <TableCell
                                                key={i}
                                                align={align}
                                                padding={padding}
                                            >
                                                <Button
                                                    variant="outlined"
                                                    size="small"
                                                    className={classes.button}
                                                    startIcon={
                                                        <Icon
                                                            name={column.Icon}
                                                            style={{
                                                                color: theme_utils.getFgColor(
                                                                    auth.theme,
                                                                    tmpColor
                                                                )
                                                            }}
                                                            marginDirection="left"
                                                            onClick={e =>
                                                                this.handlePopoverOpen(
                                                                    e,
                                                                    cellValue,
                                                                    column.View
                                                                )
                                                            }
                                                        />
                                                    }
                                                >
                                                    {cellValue.Value}
                                                </Button>
                                            </TableCell>
                                        );
                                    } else if (column.Kind === "HideText") {
                                        cells.push(
                                            <TableCell
                                                key={i}
                                                align={align}
                                                padding={padding}
                                            >
                                                <Tooltip title={cellValue}>
                                                    <IconButton
                                                        component="span"
                                                        style={{ padding: 0 }}
                                                    >
                                                        <Icon name="info" />
                                                    </IconButton>
                                                </Tooltip>
                                            </TableCell>
                                        );
                                    } else {
                                        cells.push(
                                            <TableCell
                                                key={i}
                                                align={align}
                                                padding={padding}
                                            >
                                                {cellValue}
                                            </TableCell>
                                        );
                                    }
                                } // for (let i = 0, len = columns.length; i < len; i++) {

                                return (
                                    <TableRow
                                        hover={true}
                                        tabIndex={-1}
                                        key={rowIndex}
                                    >
                                        {cells}
                                    </TableRow>
                                );
                            })}
                            {emptyRows > 0 && (
                                <TableRow style={{ height: 49 * emptyRows }}>
                                    <TableCell colSpan={6} />
                                </TableRow>
                            )}
                        </TableBody>
                    </Table>
                </div>

                <Menu
                    anchorEl={anchorEl}
                    open={Boolean(anchorEl)}
                    onClose={this.handleClose}
                    transitionDuration={100}
                >
                    {columnActions}
                </Menu>

                {actionDialog}
                <Popover
                    open={Boolean(popoverTarget)}
                    anchorEl={popoverTarget}
                    anchorOrigin={{
                        horizontal: "left",
                        vertical: "bottom"
                    }}
                    transformOrigin={{
                        horizontal: "left",
                        vertical: "top"
                    }}
                    disableRestoreFocus={true}
                    onClose={this.handlePopoverClose}
                >
                    {popoverHtml}
                </Popover>
            </Paper>
        );
    }

    private isSelected = id => {
        const { resetAction } = this.props;
        if (resetAction) {
            return false;
        }
        return this.state.selected.indexOf(id) !== -1;
    };

    private handleSelectClick = (event, id) => {
        const { selected } = this.state;
        const selectedIndex = selected.indexOf(id);
        let newSelected: any[] = [];

        if (selectedIndex === -1) {
            newSelected = newSelected.concat(selected, id);
        } else if (selectedIndex === 0) {
            newSelected = newSelected.concat(selected.slice(1));
        } else if (selectedIndex === selected.length - 1) {
            newSelected = newSelected.concat(selected.slice(0, -1));
        } else if (selectedIndex > 0) {
            newSelected = newSelected.concat(
                selected.slice(0, selectedIndex),
                selected.slice(selectedIndex + 1)
            );
        }
        console.log("TODO debug selected", newSelected);

        this.setState({ selected: newSelected });
        this.props.setAction(null);
    };

    private handleSelectAllClick = event => {
        const { data } = this.props;
        if (event.target.checked) {
            this.setState(state => ({ selected: data.map((v, i) => i) }));
            return;
        }

        this.setState({ selected: [] });
        this.props.setAction(null);
    };

    private handleRequestSort = (event, property) => {
        const orderBy = property;
        let order = "desc";

        if (this.state.orderBy === property && this.state.order === "desc") {
            order = "asc";
        }

        this.setState({ order, orderBy });
    };

    private handleChangePage = (event, page) => {
        this.setState({ page });
    };

    private handleChangeRowsPerPage = event => {
        this.setState({ rowsPerPage: event.target.value });
    };

    private handleChangeSearchInput = event => {
        let searchRegExp: any = null;
        if (event.target.value !== "") {
            searchRegExp = new RegExp(event.target.value, "i");
        }
        data_utils.setFilterParamsSearch(event.target.value);
        this.setState({ searchRegExp });
    };

    private handleActionMenuClick = (event, key) => {
        this.setState({ anchorEl: event.currentTarget, actionTarget: key });
    };

    private handleClose = () => {
        this.setState({ anchorEl: null, selected: [] });
    };

    private handleActionClick = (event, actionName) => {
        const { setAction } = this.props;
        setAction(actionName);
    };

    private handleActionDialogClose = () => {
        this.props.setAction({ actionName: null });
        this.setState({ selected: [] });
    };

    private handleFilterClick = (name, value) => {
        const { filterMap } = this.state;
        if (filterMap[name] !== undefined) {
            delete filterMap[name];
        } else {
            filterMap[name] = value;
        }
        this.setState({ filterMap });
    };

    private handleChangeOnSearchForm = (event, searchQuery) => {
        console.log("TODO handleChangeOnSearchForm");
    };

    private handleSubmitOnSearchForm = (event, index, searchQuery) => {
        console.log("TODO handleSubmitOnSearchForm");
    };

    private handlePopoverOpen = (event, data, view) => {
        const { routes, render } = this.props;
        const html = render(routes, data, view);
        this.setState({
            popoverTarget: event.currentTarget,
            popoverHtml: html
        });
    };

    private handlePopoverClose = () => {
        this.setState({ popoverTarget: null, popoverHtml: null });
    };
}

const styles = (theme: Theme): StyleRules =>
    createStyles({
        actions: {
            color: theme.palette.text.secondary
        },
        root: {
            // margin: theme.spacing.unit * 2,
            width: "100%"
        },
        spacer: {
            flex: "1 1 100%"
        },
        table: {
            width: "100%"
        },
        tableWrapper: {
            overflowX: "auto",
            width: "100%"
        },
        title: {
            flex: "0 0 auto"
        },
        toolbarExButton: {
            marginTop: theme.spacing(1)
        }
    });

function mapStateToProps(state, ownProps) {
    const auth = state.auth;
    const actionName = state.service.actionName;
    const resetAction = state.service.resetAction;
    const rootIndex = state.service.index;
    const data = data_utils.getIndexDataFromState(state, ownProps.index);

    return {
        actionName,
        auth,
        data,
        resetAction,
        rootIndex
    };
}

function mapDispatchToProps(dispatch, ownProps) {
    return {
        handleColumnLinkClick: (index, column, value) => {
            const params = {};
            params[column.LinkKey] = value;

            const location = {
                Path: column.LinkPath,
                Params: params,
                SearchQueries: {}
            };
            console.log("DEBUG TODO handleColumnLinkClick", location, value);
            dispatch(
                actions.service.serviceGetQueries({
                    index,
                    location,
                    searchQueries: null
                })
            );
        },
        setAction: actionName => {
            dispatch(
                actions.service.serviceSetAction({
                    actionName
                })
            );
        }
    };
}

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(withStyles(styles)(IndexTable));
