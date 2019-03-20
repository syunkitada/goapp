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

import Dialog from '@material-ui/core/Dialog';
import MuiDialogTitle from '@material-ui/core/DialogTitle';
import MuiDialogContent from '@material-ui/core/DialogContent';
import MuiDialogActions from '@material-ui/core/DialogActions';
import CloseIcon from '@material-ui/icons/Close';
import Button from '@material-ui/core/Button';


import ClickAwayListener from '@material-ui/core/ClickAwayListener';
import Grow from '@material-ui/core/Grow';
import Popper from '@material-ui/core/Popper';
import MenuItem from '@material-ui/core/MenuItem';
import MenuList from '@material-ui/core/MenuList';
import Popover from '@material-ui/core/Popover';


const DialogTitle = withStyles(theme => ({
  root: {
    borderBottom: `1px solid ${theme.palette.divider}`,
    margin: 0,
    padding: theme.spacing.unit * 2,
  },
  closeButton: {
    position: 'absolute',
    right: theme.spacing.unit,
    top: theme.spacing.unit,
    color: theme.palette.grey[500],
  },
}))(props => {
  const { children, classes, onClose } = props;
  return (
    <MuiDialogTitle disableTypography className={classes.root}>
      <Typography variant="h6">{children}</Typography>
      {onClose ? (
        <IconButton aria-label="Close" className={classes.closeButton} onClick={onClose}>
          <CloseIcon />
        </IconButton>
      ) : null}
    </MuiDialogTitle>
  );
});

const DialogContent = withStyles(theme => ({
  root: {
    margin: 0,
    padding: theme.spacing.unit * 2,
  },
}))(MuiDialogContent);

const DialogActions = withStyles(theme => ({
  root: {
    borderTop: `1px solid ${theme.palette.divider}`,
    margin: 0,
    padding: theme.spacing.unit,
  },
}))(MuiDialogActions);

const rows = [
  { id: 0, numeric: false, disablePadding: true, label: 'Name' },
  { id: 1, numeric: false, disablePadding: false, label: 'Region' },
  { id: 2, numeric: false, disablePadding: false, label: 'Kind' },
  { id: 3, numeric: false, disablePadding: false, label: 'UpdatedAt' },
  { id: 4, numeric: false, disablePadding: false, label: 'CreatedAt' },
  { id: 5, numeric: false, disablePadding: false, label: 'Action' },
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

class PhysicalResourcesTable extends Component {
  state = {
    order: 'asc',
    orderBy: 0,
    selected: [],
    data: [],
    page: 0,
    rowsPerPage: 5,
    searchRegExp: null,
    dialogOpen: false,
		menuOpen: false,
		openedActionMenu: -1,
		anchorEl: null,
  };

  handleClickOpen = () => {
    this.setState({
      open: true,
    });
  };

  handleClose = () => {
    this.setState({ open: false });
  };

	handleActionMenuToggle = (event) => {
			this.setState({
        anchorEl: event.currentTarget,
      });
	};

	handleClose = () => {
    this.setState({
         anchorEl: null,
    });
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
    const { match, classes, auth, datacenterIndex} = this.props
    const { order, orderBy, rowsPerPage, page, searchRegExp, menuOpen, dialogOpen, anchorEl } = this.state;

	  const open = Boolean(anchorEl);

    if (datacenterIndex == null || datacenterIndex.PhysicalResources == null) {
      return (
        <div>No Resources</div>
      )
    }
		console.log("DEBUBgllalalalala")
		console.log(menuOpen)

    const data = []
    for (let i in datacenterIndex.PhysicalResources) {
      let d = datacenterIndex.PhysicalResources[i]
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
                      <TableCell component="th" scope="row" padding="none">{n[0]}</TableCell>
                      <TableCell align="right">{n[1]}</TableCell>
                      <TableCell align="right">{n[2]}</TableCell>
                      <TableCell align="right">{n[3]}</TableCell>
                      <TableCell align="right">{n[4]}</TableCell>
                      <TableCell align="right">

														<Button
                              aria-owns={n[0] == this.state.openedActionMenu ? 'simple-popper' : undefined}
                              aria-haspopup="true"
                              variant="contained"
                                onClick={this.handleActionMenuToggle}
                            >
                              Action
                            </Button>
														  <Popover
                              id="simple-popper"
                              open={open}
                              anchorEl={anchorEl}
                              onClose={this.handleClose}
anchorOrigin={{
      vertical: 'bottom',
        horizontal: 'right',
      }}
                      transformOrigin={{
                            vertical: 'top',
                              horizontal: 'right',
                            }}
                            >
                              <Typography className={classes.typography}>The content of the Popover.</Typography>
                            </Popover>

											</TableCell>
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




        <Dialog
          onClose={this.handleClose}
          aria-labelledby="customized-dialog-title"
          open={this.state.dialogOpen}
        >
          <DialogTitle id="customized-dialog-title" onClose={this.handleClose}>
            Modal title
          </DialogTitle>
          <DialogContent>
            <Typography gutterBottom>
              Cras mattis consectetur purus sit amet fermentum. Cras justo odio, dapibus ac
              facilisis in, egestas eget quam. Morbi leo risus, porta ac consectetur ac, vestibulum
              at eros.
            </Typography>
            <Typography gutterBottom>
              Praesent commodo cursus magna, vel scelerisque nisl consectetur et. Vivamus sagittis
              lacus vel augue laoreet rutrum faucibus dolor auctor.
            </Typography>
            <Typography gutterBottom>
              Aenean lacinia bibendum nulla sed consectetur. Praesent commodo cursus magna, vel
              scelerisque nisl consectetur et. Donec sed odio dui. Donec ullamcorper nulla non metus
              auctor fringilla.
            </Typography>
          </DialogContent>
          <DialogActions>
            <Button onClick={this.handleClose} color="primary">
              Save changes
            </Button>
          </DialogActions>
        </Dialog>
      </div>
    );
  }
}

PhysicalResourcesTable.propTypes = {
  classes: PropTypes.object.isRequired,
};

function mapStateToProps(state, ownProps) {
  const auth = state.auth
  const datacenterIndex = ownProps.datacenterIndex

  return {
    auth: auth,
    datacenterIndex: datacenterIndex,
  }
}

function mapDispatchToProps(dispatch, ownProps) {
  return {}
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(withStyles(styles)(PhysicalResourcesTable));
