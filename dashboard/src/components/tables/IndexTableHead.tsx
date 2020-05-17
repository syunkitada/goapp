import * as React from "react";
import { connect } from "react-redux";

import Checkbox from "@material-ui/core/Checkbox";
import TableCell from "@material-ui/core/TableCell";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import TableSortLabel from "@material-ui/core/TableSortLabel";
import Tooltip from "@material-ui/core/Tooltip";

interface IIndexTableHead {
    index;
    order;
    orderBy;
    columns;
    numSelected;
    rowCount;
    onSelectAllClick;
    onRequestSort;
}

class IndexTableHead extends React.Component<IIndexTableHead> {
    public render() {
        const {
            index,
            order,
            orderBy,
            columns,
            numSelected,
            rowCount,
            onSelectAllClick
        } = this.props;

        return (
            <TableHead>
                <TableRow>
                    {index.SelectActions != null &&
                    index.SelectActions.length > 0 ? (
                        <TableCell padding="checkbox">
                            <Checkbox
                                indeterminate={
                                    numSelected > 0 && numSelected < rowCount
                                }
                                checked={numSelected === rowCount}
                                onChange={onSelectAllClick}
                            />
                        </TableCell>
                    ) : null}
                    {columns.map(
                        column =>
                            column.Sortable ? (
                                <TableCell
                                    key={column.id}
                                    align={
                                        column.Align ? column.Align : "right"
                                    }
                                    padding={
                                        column.Padding
                                            ? column.Padding
                                            : "default"
                                    }
                                    sortDirection={
                                        orderBy === column.id ? order : false
                                    }
                                    style={{
                                        minWidth: column.MinWidth,
                                        width: column.Width
                                    }}
                                >
                                    <Tooltip
                                        title="Sort"
                                        placement={
                                            column.numeric
                                                ? "bottom-end"
                                                : "bottom-start"
                                        }
                                        enterDelay={300}
                                    >
                                        <TableSortLabel
                                            active={orderBy === column.id}
                                            direction={order}
                                            onClick={this.createSortHandler(
                                                column.id
                                            )}
                                        >
                                            {column.Name}
                                        </TableSortLabel>
                                    </Tooltip>
                                </TableCell>
                            ) : (
                                <TableCell
                                    key={column.id}
                                    align={
                                        column.Align ? column.Align : "right"
                                    }
                                    padding={
                                        column.Padding
                                            ? column.Padding
                                            : "default"
                                    }
                                    style={{
                                        minWidth: column.MinWidth,
                                        width: column.Width
                                    }}
                                >
                                    {column.Name}
                                </TableCell>
                            ),
                        this
                    )}
                </TableRow>
            </TableHead>
        );
    }

    private createSortHandler = property => event => {
        this.props.onRequestSort(event, property);
    };
}

function mapStateToProps(state, ownProps) {
    return {};
}

function mapDispatchToProps(dispatch, ownProps) {
    return {};
}

export default connect(mapStateToProps, mapDispatchToProps)(IndexTableHead);
