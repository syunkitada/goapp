import * as React from 'react';
import {connect} from 'react-redux';

import Checkbox from '@material-ui/core/Checkbox';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import TableSortLabel from '@material-ui/core/TableSortLabel';
import Tooltip from '@material-ui/core/Tooltip';

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
      onSelectAllClick,
    } = this.props;

    return (
      <TableHead>
        <TableRow>
          {index.SelectActions != null && index.SelectActions.length > 0 ? (
            <TableCell padding="checkbox">
              <Checkbox
                indeterminate={numSelected > 0 && numSelected < rowCount}
                checked={numSelected === rowCount}
                onChange={onSelectAllClick}
              />
            </TableCell>
          ) : null}
          {columns.map(
            column => (
              <TableCell
                key={column.id}
                align={'right'}
                padding={column.disablePadding ? 'none' : 'default'}
                sortDirection={orderBy === column.id ? order : false}>
                <Tooltip
                  title="Sort"
                  placement={column.numeric ? 'bottom-end' : 'bottom-start'}
                  enterDelay={300}>
                  <TableSortLabel
                    active={orderBy === column.id}
                    direction={order}
                    onClick={this.createSortHandler(column.id)}>
                    {column.Name}
                  </TableSortLabel>
                </Tooltip>
              </TableCell>
            ),
            this,
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

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(IndexTableHead);
