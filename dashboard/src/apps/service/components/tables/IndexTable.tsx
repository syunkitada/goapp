import * as React from 'react';
import {connect} from 'react-redux';

import {Theme} from '@material-ui/core/styles/createMuiTheme';
import createStyles from '@material-ui/core/styles/createStyles';
import withStyles, {
  StyleRules,
  WithStyles,
} from '@material-ui/core/styles/withStyles';

import Button from '@material-ui/core/Button';
import Checkbox from '@material-ui/core/Checkbox';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import Menu from '@material-ui/core/Menu';
import MenuItem from '@material-ui/core/MenuItem';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableRow from '@material-ui/core/TableRow';

import KeyboardArrowDownIcon from '@material-ui/icons/KeyboardArrowDown';

import FormDialog from '../dialogs/FormDialog';
import IndexTableHead from './IndexTableHead';
import TableToolbar from './TableToolbar';

import actions from '../../../../actions';
import icon_utils from '../../../../modules/icon_utils';
import sort_utils from '../../../../modules/sort_utils';

interface IIndexTable extends WithStyles<typeof styles> {
  routes;
  index;
  data;
  getQueries;
}

interface IState {
  order;
  orderBy;
  selected: any[];
  data: any[];
  page;
  rowsPerPage;
  searchRegExp: any;
  anchorEl: any;
  actionName: any;
}

class IndexTable extends React.Component<IIndexTable> {
  public state: IState = {
    actionName: null,
    anchorEl: null,
    data: [],
    order: 'asc',
    orderBy: 0,
    page: 0,
    rowsPerPage: 5,
    searchRegExp: null,
    selected: [],
  };

  public render() {
    const {routes, classes, index, data} = this.props;
    const {
      selected,
      anchorEl,
      order,
      orderBy,
      rowsPerPage,
      page,
      searchRegExp,
      actionName,
    } = this.state;

    const columns = index.Columns;
    let rawData = data[index.DataKey];

    if (!rawData) {
      rawData = [];
    }

    let isSelectActions = false;
    if (index.SelectActions) {
      isSelectActions = true;
    }

    const currentRoute = routes.slice(-1)[0];
    const beforeRoute = routes.slice(-2)[0];

    const searchColumns: any[] = [];
    for (let i = 0, len = columns.length; i < len; i++) {
      if (columns[i].IsSearch) {
        searchColumns.push(columns[i].Name);
      }
    }

    let isSkip = true;
    const tableData: any[] = [];
    for (let i = 0, len = rawData.length; i < len; i++) {
      const d = rawData[i];
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

      const row: any[] = [];
      for (const column of columns) {
        const c = d[column.Name];
        if (column.Type === 'Time') {
          const time = new Date(c.seconds * 1000);
          row.push(time.toISOString());
        } else if (column.Type === 'Action') {
          row.push('');
        } else {
          row.push(c);
        }
      }
      tableData.push(row);
    }

    for (let i = 0, l = columns.length; i < l; i++) {
      columns[i].id = i;
    }

    const columnActions: any[] = [];
    if (index.ColumnActions != null) {
      for (let i = 0, len = index.ColumnActions.length; i < len; i++) {
        const columnAction = index.ColumnActions[i];
        columnActions.push(
          <MenuItem
            key={columnAction.Name}
            onClick={event => this.handleActionClick(event, columnAction.Name)}>
            <ListItemIcon>{icon_utils.getIcon(columnAction.Icon)}</ListItemIcon>
            <ListItemText inset={true} primary={columnAction.Name} />
          </MenuItem>,
        );
      }
    }

    let action: any = null;
    let actionDialog: any = null;
    if (actionName !== null) {
      for (const a of index.Actions) {
        if (a.Name === actionName) {
          action = a;
          break;
        }
      }
      if (action === null) {
        for (const a of index.ColumnActions) {
          if (a.Name === actionName) {
            action = a;
            break;
          }
        }
      }
      if (action === null) {
        actionDialog = null;
      } else {
        switch (action.Kind) {
          case 'Form':
            actionDialog = (
              <FormDialog
                open={true}
                data={data}
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

    const indexLength = tableData.length;
    const emptyRows =
      rowsPerPage - Math.min(rowsPerPage, indexLength - page * rowsPerPage);

    return (
      <div className={classes.root}>
        <TableToolbar
          count={indexLength}
          rowsPerPage={rowsPerPage}
          page={page}
          index={index}
          numSelected={selected.length}
          onChangePage={this.handleChangePage}
          onChangeRowsPerPage={this.handleChangeRowsPerPage}
          onChangeSearchInput={this.handleChangeSearchInput}
          onActionClick={this.handleActionClick}
        />
        <div className={classes.tableWrapper}>
          <Table className={classes.table} aria-labelledby="tableTitle">
            <IndexTableHead
              index={index}
              order={order}
              orderBy={orderBy}
              onSelectAllClick={this.handleSelectAllClick}
              onRequestSort={this.handleRequestSort}
              rowCount={indexLength}
              columns={columns}
              numSelected={selected.length}
            />
            <TableBody>
              {sort_utils
                .stableSort(tableData, sort_utils.getSorting(order, orderBy))
                .slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
                .map(n => {
                  const cells: any = [];
                  const key = n[0];
                  const isSelected = this.isSelected(key);

                  if (isSelectActions) {
                    cells.push(
                      <TableCell
                        key={-1}
                        padding="checkbox"
                        onClick={event => this.handleSelectClick(event, key)}>
                        <Checkbox checked={isSelected} />
                      </TableCell>,
                    );
                  }

                  for (let i = 0, len = columns.length; i < len; i++) {
                    if (columns[i].Link) {
                      let link = columns[i].Link;
                      const splitedLink = link.split('/');
                      const splitedNextLink: any[] = [];
                      const baseUrl = beforeRoute.match.url;
                      for (
                        let j = 0, linkLen = splitedLink.length;
                        j < linkLen;
                        j++
                      ) {
                        let path = splitedLink[j];
                        if (path.indexOf(':') === 0) {
                          const pathKey = path.slice(1);
                          const tmppath = currentRoute.match.params[pathKey];
                          if (tmppath) {
                            path = tmppath;
                          } else {
                            path = n[parseInt(pathKey, 10)];
                          }
                        }
                        splitedNextLink.push(path);
                      }
                      link = baseUrl + '/' + splitedNextLink.join('/');
                      cells.push(
                        <TableCell
                          align="right"
                          key={i}
                          component="th"
                          scope="row"
                          padding="none"
                          style={{cursor: 'pointer'}}
                          onClick={e => {
                            this.handleLinkClick(e, link, n[i], columns[i]);
                          }}>
                          {n[i]}
                        </TableCell>,
                      );
                    } else if (columns[i].Type === 'Action') {
                      cells.push(
                        <TableCell key={i} align="right">
                          <Button
                            aria-owns={anchorEl ? 'simple-menu' : undefined}
                            aria-haspopup="true"
                            variant="outlined"
                            onClick={e => {
                              this.handleActionMenuClick(e, key);
                            }}>
                            Actions <KeyboardArrowDownIcon />
                          </Button>
                        </TableCell>,
                      );
                    } else {
                      cells.push(
                        <TableCell key={i} align="right">
                          {n[i]}
                        </TableCell>,
                      );
                    }
                  }
                  return (
                    <TableRow hover={true} tabIndex={-1} key={key}>
                      {cells}
                    </TableRow>
                  );
                })}
              {emptyRows > 0 && (
                <TableRow style={{height: 49 * emptyRows}}>
                  <TableCell colSpan={6} />
                </TableRow>
              )}
            </TableBody>
          </Table>

          <Menu
            anchorEl={anchorEl}
            open={Boolean(anchorEl)}
            onClose={this.handleClose}
            transitionDuration={100}>
            {columnActions}
          </Menu>

          {actionDialog}
        </div>
      </div>
    );
  }

  private isSelected = id => {
    return this.state.selected.indexOf(id) !== -1;
  };

  private handleSelectClick = (event, id) => {
    const {selected} = this.state;
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
        selected.slice(selectedIndex + 1),
      );
    }

    this.setState({selected: newSelected});
  };

  private handleSelectAllClick = event => {
    const {index, data} = this.props;
    const columns = index.Columns;
    const keyColumn = columns[0].Name;
    const rawData = data[index.DataKey];
    if (event.target.checked) {
      this.setState(state => ({selected: rawData.map(n => n[keyColumn])}));
      return;
    }

    this.setState({selected: []});
  };

  private handleRequestSort = (event, property) => {
    const orderBy = property;
    let order = 'desc';

    if (this.state.orderBy === property && this.state.order === 'desc') {
      order = 'asc';
    }

    this.setState({order, orderBy});
  };

  private handleChangePage = (event, page) => {
    this.setState({page});
  };

  private handleChangeRowsPerPage = event => {
    this.setState({rowsPerPage: event.target.value});
  };

  private handleChangeSearchInput = event => {
    let searchRegExp: any = null;
    if (event.target.value !== '') {
      searchRegExp = new RegExp(event.target.value, 'i');
    }
    this.setState({searchRegExp});
  };

  private handleActionMenuClick = (event, key) => {
    this.setState({anchorEl: event.currentTarget, actionTarget: key});
  };

  private handleClose = () => {
    this.setState({anchorEl: null});
  };

  private handleActionClick = (event, actionName) => {
    this.setState({actionName});
  };

  private handleActionDialogClose = () => {
    this.setState({actionName: null});
  };

  private handleLinkClick = (event, link, value, column) => {
    const {routes} = this.props;
    const route = routes[routes.length - 1];
    const params = route.match.params;
    params[column.LinkParam] = value;
    this.props.getQueries(
      column.LinkGetQueries,
      column.LinkSync,
      route.match.params,
    );
    route.history.push(link);
  };
}

const styles = (theme: Theme): StyleRules =>
  createStyles({
    actions: {
      color: theme.palette.text.secondary,
    },
    root: {
      // margin: theme.spacing.unit * 2,
      width: '100%',
    },
    spacer: {
      flex: '1 1 100%',
    },
    table: {
      width: '100%',
    },
    tableWrapper: {
      overflowX: 'auto',
    },
    title: {
      flex: '0 0 auto',
    },
  });

function mapStateToProps(state, ownProps) {
  const auth = state.auth;

  return {
    auth,
  };
}

function mapDispatchToProps(dispatch, ownProps) {
  return {
    getQueries: (queries, isSync, params) => {
      dispatch(
        actions.service.serviceGetQueries({
          isSync,
          params,
          queries,
        }),
      );
    },
  };
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(withStyles(styles)(IndexTable));
