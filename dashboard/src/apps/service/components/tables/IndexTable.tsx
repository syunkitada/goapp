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
import Grid from '@material-ui/core/Grid';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import Menu from '@material-ui/core/Menu';
import MenuItem from '@material-ui/core/MenuItem';
import Popover from '@material-ui/core/Popover';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableRow from '@material-ui/core/TableRow';

import KeyboardArrowDownIcon from '@material-ui/icons/KeyboardArrowDown';

import FormDialog from '../dialogs/FormDialog';
import IndexTableHead from './IndexTableHead';
import TableToolbar from './TableToolbar';

import Badge from '@material-ui/core/Badge';
import IconButton from '@material-ui/core/IconButton';
import TextField from '@material-ui/core/TextField';

import actions from '../../../../actions';
import logger from '../../../../lib/logger';
import theme_utils from '../../../../lib/theme_utils';
import icon_utils from '../../../../modules/icon_utils';
import sort_utils from '../../../../modules/sort_utils';

import Tooltip from '@material-ui/core/Tooltip';
import Icon from '../../../../components/icons/Icon';

import SearchIcon from '@material-ui/icons/Search';

import Autocomplete from '@material-ui/lab/Autocomplete';

import {DateTimePicker} from '@material-ui/pickers';
import {MuiPickersUtilsProvider} from '@material-ui/pickers';

import DateFnsUtils from '@date-io/date-fns';

interface IIndexTable extends WithStyles<typeof styles> {
  auth;
  render;
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
  popoverHtml;
  popoverTarget;
  rowsPerPage;
  searchQueries: any;
  searchRegExp: any;
  anchorEl: any;
  actionName: any;
  filterMap: any;
}

class IndexTable extends React.Component<IIndexTable> {
  public state: IState = {
    actionName: null,
    anchorEl: null,
    data: [],
    filterMap: {},
    order: 'asc',
    orderBy: 0,
    page: 0,
    popoverHtml: null,
    popoverTarget: null,
    rowsPerPage: 5,
    searchQueries: {},
    searchRegExp: null,
    selected: [],
  };

  public componentWillMount() {
    const {routes} = this.props;
    const route = routes[routes.length - 1];
    const location = route.location;
    const queryStr = decodeURIComponent(location.search);
    try {
      const value = queryStr.match(new RegExp('[?&]q=({.*?})(&|$|#)'));
      if (value) {
        const searchQueries = JSON.parse(value[1]);
        this.setState({searchQueries});
      }
    } catch (e) {
      console.log('Ignored failed parse', queryStr);
    }
  }

  public render() {
    const {auth, routes, classes, index, data} = this.props;
    const {
      popoverTarget,
      popoverHtml,
      selected,
      anchorEl,
      rowsPerPage,
      page,
      searchRegExp,
      actionName,
      searchQueries,
      filterMap,
    } = this.state;
    logger.info('IndexTable', 'render', actionName, routes);
    console.log('DEBUG Table', index.DataKey, index.Columns, data);

    console.log(auth);
    const exButtons: any[] = [];
    const exInputs: any[] = [];
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
    let order = 'asc';
    let orderBy = 0;
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
        for (let j = 0, lenj = column.FilterValues.length; j < lenj; j++) {
          const filterValue = column.FilterValues[j];
          filterCounterMap[filterValue.Value] = 0;
        }
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
        if (column.Kind === 'Time') {
          if (c) {
            const time: any = new Date(c);
            if (!isNaN(time.getTime())) {
              row.push(time.toISOString());
            } else {
              row.push(time.toString());
            }
          } else {
            row.push('');
          }
        } else if (column.Kind === 'Action') {
          row.push('');
        } else if (column.Kind === 'Popover') {
          if (column.View) {
            const popoverData = {
              Data: d,
              Value: c,
            };
            row.push(popoverData);
          } else {
            const popoverData = {
              Data: null,
              Value: c,
            };
            row.push(popoverData);
          }
        } else {
          row.push(c);
        }

        if (column.FilterValues) {
          for (let j = 0, lenj = column.FilterValues.length; j < lenj; j++) {
            const filterValue = column.FilterValues[j];
            if (filterValue.Value === c) {
              filterCounterMap[filterValue.Value] += 1;
            }
          }
        }
      } // for (const column of columns) {

      tableData.push(row);
    }
    console.log('DEBUG counterMap', filterCounterMap);

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
            onClick={event => this.handleActionClick(event, columnAction.Name)}>
            <ListItemIcon>{icon_utils.getIcon(columnAction.Icon)}</ListItemIcon>
            <ListItemText inset={true} primary={columnAction.Name} />
          </MenuItem>,
        );
      }
    }

    let actionDialog: any = null;
    if (actionName !== null) {
      let action: any = null;
      for (const a of index.Actions) {
        if (a.Name === actionName) {
          action = a;
          break;
        }
      }

      if (action === null) {
        for (const a of index.SelectActions) {
          if (a.Name === actionName) {
            action = a;
            break;
          }
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
                selected={selected}
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
        for (let j = 0, lenj = column.FilterValues.length; j < lenj; j++) {
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
                style={{opacity}}
                onClick={() =>
                  this.handleFilterClick(column.Name, filterValue.Value)
                }
                value={filterValue.Value}>
                <Badge
                  badgeContent={filterCounterMap[filterValue.Value]}
                  color="primary">
                  <Icon
                    name={filterValue.Icon}
                    style={{
                      color: theme_utils.getFgColor(
                        auth.theme,
                        filterValue.Icon,
                      ),
                    }}
                    marginDirection="left"
                  />
                </Badge>
              </IconButton>
            </Tooltip>,
          );
        }
      }
    }

    let exInputsForm: any = null;
    if (index.ExInputs) {
      for (let i = 0, len = index.ExInputs.length; i < len; i++) {
        const input = index.ExInputs[i];
        let defaultValue = searchQueries[input.Name];
        if (!defaultValue) {
          defaultValue = input.Default;
        }
        switch (input.Type) {
          case 'Selector':
            let selectorData: any;
            if (input.DataKey) {
              selectorData = data[input.DataKey];
            } else if (input.Data) {
              selectorData = input.Data;
            }
            const options: any = [];
            if (!selectorData) {
              continue;
            }
            for (let j = 0, lenj = selectorData.length; j < lenj; j++) {
              options.push(selectorData[j]);
            }

            const selectorProps = {
              getOptionLabel: option => option,
              options,
            };

            exInputs.push(
              <Grid item={true} key={input.Name}>
                <Autocomplete
                  {...selectorProps}
                  multiple={input.Multiple}
                  disableCloseOnSelect={true}
                  defaultValue={defaultValue}
                  onChange={(event, values) =>
                    this.handleSelectorChange(event, input.Name, values)
                  }
                  renderInput={params => (
                    <TextField
                      {...params}
                      size="small"
                      label={input.Name}
                      variant="outlined"
                    />
                  )}
                />
              </Grid>,
            );
            break;
          case 'Text':
            exInputs.push(
              <Grid item={true} key={input.Name}>
                <TextField
                  label={input.Name}
                  defaultValue={defaultValue}
                  variant="outlined"
                  size="small"
                  name={input.Name}
                  onChange={this.handleInputChange}
                />
              </Grid>,
            );
            break;
          case 'DateTime':
            exInputs.push(
              <Grid item={true} key={input.Name}>
                <MuiPickersUtilsProvider utils={DateFnsUtils}>
                  <DateTimePicker
                    label={input.Name}
                    inputVariant="outlined"
                    size="small"
                    value={defaultValue}
                    format="yyyy/MM/dd HH:mm"
                    showTodayButton={true}
                    onChange={(date: Date | null) =>
                      this.handleDateChange(input.Name, date)
                    }
                  />
                </MuiPickersUtilsProvider>
              </Grid>,
            );
            break;
        }
      }

      exInputsForm = (
        <form onSubmit={this.handleInputSubmit}>
          <Grid container={true} direction="row" spacing={1}>
            {exInputs}
            <Grid item={true}>
              <Button
                variant="outlined"
                type="submit"
                size="medium"
                color="primary"
                startIcon={<SearchIcon />}>
                Submit
              </Button>
            </Grid>
          </Grid>
        </form>
      );
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
          exButtons={exButtons}
          exInputsForm={exInputsForm}
        />
        <div className={classes.tableWrapper}>
          <Table
            className={classes.table}
            aria-labelledby="tableTitle"
            size="small">
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
                .map((n, rowIndex) => {
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
                    const column = columns[i];
                    if (column.Link) {
                      let link = column.Link;
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
                          align={i === 0 ? 'left' : 'right'}
                          key={i}
                          component="th"
                          scope="row"
                          style={{cursor: 'pointer'}}
                          onClick={e => {
                            this.handleLinkClick(e, link, n[i], column);
                          }}>
                          {n[i]}
                        </TableCell>,
                      );
                    } else if (column.Kind === 'Action') {
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
                    } else if (column.FilterValues) {
                      let filterButton: any = null;
                      const currentValue = filterMap[column.Name];
                      let isShowCells = true;
                      if (currentValue) {
                        isShowCells = false;
                      }
                      if (column.FilterValues) {
                        for (
                          let j = 0, lenj = column.FilterValues.length;
                          j < lenj;
                          j++
                        ) {
                          const filterValue = column.FilterValues[j];
                          if (filterValue.Value === n[i]) {
                            if (
                              currentValue &&
                              currentValue === filterValue.Value
                            ) {
                              isShowCells = true;
                            }
                            filterButton = (
                              <Button>
                                <Icon
                                  key={j}
                                  name={filterValue.Icon}
                                  style={{
                                    color: theme_utils.getFgColor(
                                      auth.theme,
                                      filterValue.Icon,
                                    ),
                                  }}
                                  marginDirection="left"
                                />
                                {n[i]}
                              </Button>
                            );
                            break;
                          }
                        }
                      }
                      if (!isShowCells) {
                        return;
                      }
                      cells.push(
                        <TableCell key={i} align={i === 0 ? 'left' : 'right'}>
                          {filterButton}
                        </TableCell>,
                      );
                    } else if (column.Kind === 'Popover') {
                      let tmpColor = column.Color;
                      let isInactive = true;
                      if (n[i].Data) {
                        const tmpData = n[i].Data[column.View.DataKey];
                        if (column.View.Kind === 'Table') {
                          if (tmpData && tmpData.length > 0) {
                            isInactive = false;
                          }
                        }
                      }
                      if (isInactive) {
                        tmpColor = column.InactiveColor;
                      }
                      cells.push(
                        <TableCell key={i} align={i === 0 ? 'left' : 'right'}>
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
                                    tmpColor,
                                  ),
                                }}
                                marginDirection="left"
                                onClick={e =>
                                  this.handlePopoverOpen(e, n[i], column.View)
                                }
                              />
                            }>
                            {n[i].Value}
                          </Button>
                        </TableCell>,
                      );
                    } else {
                      cells.push(
                        <TableCell key={i} align={i === 0 ? 'left' : 'right'}>
                          {n[i]}
                        </TableCell>,
                      );
                    }
                  } // for (let i = 0, len = columns.length; i < len; i++) {

                  return (
                    <TableRow hover={true} tabIndex={-1} key={rowIndex}>
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
        <Popover
          open={Boolean(popoverTarget)}
          anchorEl={popoverTarget}
          anchorOrigin={{
            horizontal: 'left',
            vertical: 'bottom',
          }}
          transformOrigin={{
            horizontal: 'left',
            vertical: 'top',
          }}
          disableRestoreFocus={true}
          onClose={this.handlePopoverClose}>
          {popoverHtml}
        </Popover>
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
    this.setState({actionName: null, selected: []});
  };

  private handleLinkClick = (event, link, value, column) => {
    const {routes} = this.props;
    const {searchQueries} = this.state;
    const route = routes[routes.length - 1];
    const params = route.match.params;
    params[column.LinkParam] = value;
    this.props.getQueries(
      column.LinkGetQueries,
      searchQueries,
      column.LinkSync,
      params,
    );
    route.history.push(link);
  };

  private handleFilterClick = (name, value) => {
    const {filterMap} = this.state;
    if (filterMap[name]) {
      delete filterMap[name];
    } else {
      filterMap[name] = value;
    }
    this.setState({filterMap});
  };

  private handleSelectorChange = (event, name, values) => {
    const {searchQueries} = this.state;
    searchQueries[name] = values;
    this.setState({searchQueries});
  };

  private handleInputSubmit = event => {
    const {routes} = this.props;
    const {searchQueries} = this.state;
    const route = routes[routes.length - 1];
    const location = route.location;
    console.log('debug handle submit sub');
    console.log('DEBUG location', location);
    console.log('DEBUG submit', searchQueries);
    event.preventDefault();
    // route.history.push(link);
    const queryStr = encodeURIComponent(JSON.stringify(searchQueries));
    location.search = 'q=' + queryStr;
    route.history.push(location);
  };

  private handleInputChange = event => {
    const {searchQueries} = this.state;
    searchQueries[event.target.name] = event.target.value;
    this.setState({searchQueries});
  };

  private handleDateChange = (name: string, date: Date | null) => {
    const {searchQueries} = this.state;
    searchQueries[name] = date;
    this.setState({searchQueries});
  };

  private handlePopoverOpen = (event, data, view) => {
    console.log('handlePopoverOpen', data, view);
    const {routes, render} = this.props;
    const html = render(routes, data, view);
    this.setState({popoverTarget: event.currentTarget, popoverHtml: html});
  };

  private handlePopoverClose = () => {
    this.setState({popoverTarget: null, popoverHtml: null});
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
    toolbarExButton: {
      marginTop: theme.spacing(1),
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
    getQueries: (queries, searchQueries, isSync, params) => {
      dispatch(
        actions.service.serviceGetQueries({
          isSync,
          params,
          queries,
          searchQueries,
        }),
      );
    },
  };
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(withStyles(styles)(IndexTable));
