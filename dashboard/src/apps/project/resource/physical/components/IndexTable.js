import React, {Component} from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';

import classNames from 'classnames';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TablePagination from '@material-ui/core/TablePagination';
import TableRow from '@material-ui/core/TableRow';
import TableSortLabel from '@material-ui/core/TableSortLabel';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Paper from '@material-ui/core/Paper';
import Checkbox from '@material-ui/core/Checkbox';
import IconButton from '@material-ui/core/IconButton';
import Tooltip from '@material-ui/core/Tooltip';
import DeleteIcon from '@material-ui/icons/Delete';
import NotificationsOffIcon from '@material-ui/icons/NotificationsOff';
import FilterListIcon from '@material-ui/icons/FilterList';
import { lighten } from '@material-ui/core/styles/colorManipulator';

import { fade } from '@material-ui/core/styles/colorManipulator';
import FirstPageIcon from '@material-ui/icons/FirstPage';
import KeyboardArrowLeft from '@material-ui/icons/KeyboardArrowLeft';
import KeyboardArrowRight from '@material-ui/icons/KeyboardArrowRight';
import LastPageIcon from '@material-ui/icons/LastPage';

import InputBase from '@material-ui/core/InputBase';
import Input from '@material-ui/core/Input';
import InputLabel from '@material-ui/core/InputLabel';
import InputAdornment from '@material-ui/core/InputAdornment';
import FormControl from '@material-ui/core/FormControl';
import SearchIcon from '@material-ui/icons/Search';

import Grid from '@material-ui/core/Grid';

import Badge from '@material-ui/core/Badge';
import ShoppingCartIcon from '@material-ui/icons/ShoppingCart';


import CheckCircleIcon from '@material-ui/icons/CheckCircle';
import CheckCircleOutlineIcon from '@material-ui/icons/CheckCircleOutline';
import WarningIcon from '@material-ui/icons/Warning';
import ErrorIcon from '@material-ui/icons/Error';
import ErrorOutlineIcon from '@material-ui/icons/ErrorOutline';
import NotificationImportantIcon from '@material-ui/icons/NotificationImportant';
import NotificationsNoneIcon from '@material-ui/icons/NotificationsNone';

import IndexTableHead         from '../../../../../components/tables/IndexTableHead'
import TableToolbar from '../../../../../components/tables/TableToolbar'
import sort_utils from '../../../../../modules/sort_utils'

const rows = [
  { id: 0, numeric: false, disablePadding: true, label: 'Name' },
  { id: 1, numeric: false, disablePadding: false, label: 'Region' },
  { id: 2, numeric: false, disablePadding: false, label: 'Kind' },
  { id: 3, numeric: false, disablePadding: false, label: 'UpdatedAt' },
  { id: 4, numeric: false, disablePadding: false, label: 'CreatedAt' },
];

const styles = theme => ({
  root: {
    width: '100%',
  },
  table: {
    minWidth: 1020,
  },
  tableWrapper: {
    overflowX: 'auto',
  },
  margin: {
    margin: theme.spacing.unit * 2,
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

class IndexTable extends Component {
  state = {
    order: 'asc',
    orderBy: 0,
    selected: [],
    data: [],
    page: 0,
    rowsPerPage: 5,
    searchRegExp: null,
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
    if (event.target.value != '') {
      searchRegExp = new RegExp(event.target.value, 'i');
    }
    this.setState({ searchRegExp: searchRegExp });
  };

  render() {
    const { match, classes, auth, index} = this.props
    const { order, orderBy, rowsPerPage, page, searchRegExp } = this.state;

    const data = []
    for (let i in index.Datacenters) {
      let d = index.Datacenters[i]
      if (searchRegExp && !searchRegExp.exec(d.Name) && !searchRegExp.exec(d.Region)) {
        continue
      }
      let updatedAt = new Date(d.UpdatedAt.seconds * 1000)
      let createdAt = new Date(d.CreatedAt.seconds * 1000)
      data.push([d.Name, d.Region, d.Kind, updatedAt.toISOString(), createdAt.toISOString()])
    }

    const indexLength = data.length
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
              rows={rows}
            />
            <TableBody>
              {sort_utils.stableSort(data, sort_utils.getSorting(order, orderBy))
                .slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
                .map(n => {
                  return (
                    <TableRow
                      hover
                      role="checkbox"
                      tabIndex={-1}
                      key={n[0]}
                    >
                      <TableCell component="th" scope="row" padding="none">
                        <Link to={`${match.url}/${n[0]}`}>{n[0]}</Link>
                      </TableCell>
                      <TableCell align="right">{n[1]}</TableCell>
                      <TableCell align="right">{n[2]}</TableCell>
                      <TableCell align="right">{n[3]}</TableCell>
                      <TableCell align="right">{n[4]}</TableCell>
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
      </div>
    );
  }
}

IndexTable.propTypes = {
  classes: PropTypes.object.isRequired,
};

function mapStateToProps(state, ownProps) {
  const auth = state.auth
  const index = ownProps.index

  return {
    auth: auth,
    index: index,
  }
}

function mapDispatchToProps(dispatch, ownProps) {
  return {}
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(withStyles(styles)(IndexTable));
