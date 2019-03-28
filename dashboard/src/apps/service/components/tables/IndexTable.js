import React, {Component} from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';

import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableRow from '@material-ui/core/TableRow';

import Button from '@material-ui/core/Button';
import Menu from '@material-ui/core/Menu';
import MenuItem from '@material-ui/core/MenuItem';

import IndexTableHead from './IndexTableHead'
import TableToolbar from './TableToolbar'
import sort_utils from '../../../../modules/sort_utils'

class IndexTable extends Component {
  state = {
    order: 'asc',
    orderBy: 0,
    selected: [],
    data: [],
    page: 0,
    rowsPerPage: 5,
    searchRegExp: null,
    anchorEl: null,
  };

  handleRequestSort = (event, property) => {
    const orderBy = property;
    let order = 'desc';

    if (this.state.orderBy === property && this.state.order === 'desc') {
      order = 'asc';
    }

    this.setState({ order, orderBy });
  };

  handleChangePage = (event, page) => {
    this.setState({ page });
  };

  handleChangeRowsPerPage = event => {
    this.setState({ rowsPerPage: event.target.value });
  };

  handleChangeSearchInput = event => {
    let searchRegExp = null
    if (event.target.value !== '') {
      searchRegExp = new RegExp(event.target.value, 'i');
    }
    this.setState({ searchRegExp: searchRegExp });
  };

  handleClick = event => {
    this.setState({ anchorEl: event.currentTarget });
  };

  handleClose = () => {
    this.setState({ anchorEl: null });
  };

  render() {
    const { routes, classes, columns, data} = this.props
    const { anchorEl, order, orderBy, rowsPerPage, page, searchRegExp } = this.state;

    if (!data) {
      return null
    }

    let beforeRoute = routes.slice(-2)[0]

    let searchColumns = []
    for (let i = 0, len = columns.length; i < len; i++) {
      if (columns[i].IsSearch) {
        searchColumns.push(columns[i].Name)
      }
    }

    let isSkip = true
    const tableData = []
    for (let d of data) {
      if (searchRegExp != null) {
        for (let c of searchColumns) {
          if (searchRegExp.exec(d[c])) {
            isSkip = false
            break
          }
        }
        if (isSkip) {
          continue
        }
        isSkip = true
      }

      let row = []
      for (let column of columns) {
        let c = d[column.Name]
        if (column.Type === "Time") {
          let time = new Date(c.seconds * 1000)
          row.push(time.toISOString())
        } else if (column.Type === "Action") {
          row.push("")
        } else {
          row.push(c)
        }
      }
      tableData.push(row)
    }

    for (let i = 0, l = columns.length; i < l; i++) {
      columns[i].id = i
    }

    const indexLength = tableData.length
    const emptyRows = rowsPerPage - Math.min(rowsPerPage, indexLength - page * rowsPerPage);

    return (
      <div className={classes.root}>
        <TableToolbar
          count={indexLength}
          rowsPerPage={rowsPerPage}
          page={page}
          onChangePage={this.handleChangePage}
          onChangeRowsPerPage={this.handleChangeRowsPerPage}
          onChangeSearchInput={this.handleChangeSearchInput}
        />
        <div className={classes.tableWrapper}>
          <Table className={classes.table} aria-labelledby="tableTitle">
            <IndexTableHead
              order={order}
              orderBy={orderBy}
              onSelectAllClick={this.handleSelectAllClick}
              onRequestSort={this.handleRequestSort}
              rowCount={indexLength}
              columns={columns}
            />
            <TableBody>
              {sort_utils.stableSort(tableData, sort_utils.getSorting(order, orderBy))
                .slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
                .map(n => {
                  let cells = []
                  for (let i = 0, len = columns.length; i < len; i++) {
                    if (columns[i].Link) {
                      cells.push(
                        <TableCell key={i} component="th" scope="row" padding="none">
                          <Link to={`${beforeRoute.match.url}${columns[i].Link}/${n[0]}`}>{n[i]}</Link>
                        </TableCell>
                      )
                    } else if (columns[i].Type === "Action") {
                      cells.push(
                        <TableCell key={i} align="right">
                          <Button
                            aria-owns={anchorEl ? 'simple-menu' : undefined}
                            aria-haspopup="true"
                            variant="outlined"
                            onClick={this.handleClick}
                          >
                            Open Actions
                          </Button>
                        </TableCell>
                      )
                    } else {
                      cells.push(
                        <TableCell key={i} align="right">{n[i]}</TableCell>
                      )
                    }
                  }
                  return (
                    <TableRow
                      hover
                      tabIndex={-1}
                      key={n[0]}
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

          <Menu
            anchorEl={anchorEl}
            open={Boolean(anchorEl)}
            onClose={this.handleClose}
            transitionDuration={100}
          >
            <MenuItem onClick={this.handleClose}>Profile</MenuItem>
            <MenuItem onClick={this.handleClose}>My account</MenuItem>
            <MenuItem onClick={this.handleClose}>Logout</MenuItem>
          </Menu>
        </div>
      </div>
    );
  }
}

const styles = theme => ({
  root: {
    // margin: theme.spacing.unit * 2,
    width: '100%',
  },
  table: {
    width: '100%',
  },
  tableWrapper: {
    overflowX: 'auto',
  },
  margin: {
    // margin: theme.spacing.unit,
  },
  spacer: {
    flex: '1 1 100%',
  },
  actions: {
    color: theme.palette.text.secondary,
  },
  title: {
    flex: '0 0 auto',
  },
});

IndexTable.propTypes = {
  classes: PropTypes.object.isRequired,
};

function mapStateToProps(state, ownProps) {
  const auth = state.auth

  return {
    auth: auth,
  }
}

function mapDispatchToProps(dispatch, ownProps) {
  return {}
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(withStyles(styles)(IndexTable));
