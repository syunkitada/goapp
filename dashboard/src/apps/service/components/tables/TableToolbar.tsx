import * as React from 'react';
import {connect} from 'react-redux';

import {Theme} from '@material-ui/core/styles/createMuiTheme';
import createStyles from '@material-ui/core/styles/createStyles';
import withStyles, {
  StyleRules,
  WithStyles,
} from '@material-ui/core/styles/withStyles';

import {fade} from '@material-ui/core/styles/colorManipulator';
import {lighten} from '@material-ui/core/styles/colorManipulator';

import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import Grid from '@material-ui/core/Grid';
import IconButton from '@material-ui/core/IconButton';
import Input from '@material-ui/core/Input';
import InputAdornment from '@material-ui/core/InputAdornment';
import Toolbar from '@material-ui/core/Toolbar';
import Tooltip from '@material-ui/core/Tooltip';

import Badge from '@material-ui/core/Badge';
import ErrorOutlineIcon from '@material-ui/icons/ErrorOutline';
import HighlightOffOutlinedIcon from '@material-ui/icons/HighlightOffOutlined';

import SearchIcon from '@material-ui/icons/Search';

import icon_utils from '../../../../modules/icon_utils';
import TablePagination from './TablePagination';

const styles = (theme: Theme): StyleRules =>
  createStyles({
    actions: {
      color: theme.palette.text.secondary,
    },
    buttonMargin: {
      marginBottom: theme.spacing(2),
      marginTop: theme.spacing(2),
    },
    highlight:
      theme.palette.type === 'light'
        ? {
            backgroundColor: lighten(theme.palette.secondary.light, 0.85),
            color: theme.palette.secondary.main,
          }
        : {
            backgroundColor: theme.palette.secondary.dark,
            color: theme.palette.text.primary,
          },
    leftButton: {
      marginTop: theme.spacing(1),
    },
    margin: {
      margin: theme.spacing(2),
    },
    root: {
      width: '100%',
    },
    search: {
      '&:hover': {
        backgroundColor: fade(theme.palette.common.white, 0.25),
      },
      backgroundColor: fade(theme.palette.common.white, 0.15),
      borderRadius: theme.shape.borderRadius,
      marginLeft: 0,
      marginRight: theme.spacing(2),
      position: 'relative',
      width: '100%',
      [theme.breakpoints.up('sm')]: {
        marginLeft: theme.spacing(3),
        width: 'auto',
      },
    },
    searchIcon: {
      alignItems: 'center',
      display: 'flex',
      height: '100%',
      justifyContent: 'center',
      pointerEvents: 'none',
      position: 'absolute',
      width: theme.spacing(9),
    },
    spacer: {
      flex: '1 1 100%',
    },
    table: {
      minWidth: 1020,
    },
    tableWrapper: {
      overflowX: 'auto',
    },
    title: {
      flex: '0 0 auto',
    },
  });

interface ITableToolbar extends WithStyles<typeof styles> {
  index;
  onChangeSearchInput;
  onChangeFilter;
  count;
  numSelected;
  rowsPerPage;
  page;
  onChangePage;
  onChangeRowsPerPage;
  onActionClick;
}

class TableToolbar extends React.Component<ITableToolbar> {
  public render() {
    const {
      classes,
      index,
      onChangeSearchInput,
      count,
      numSelected,
      rowsPerPage,
      page,
      onChangePage,
      onChangeRowsPerPage,
      onChangeFilter,
      onActionClick,
    } = this.props;
    console.log(onChangeFilter);

    const actionButtons: any[] = [];
    if (numSelected > 0) {
      if (index.SelectActions != null) {
        actionButtons.push(
          <Button key={-1} color="secondary">
            {numSelected} selected
          </Button>,
        );
        for (let i = 0, len = index.SelectActions.length; i < len; i++) {
          const action = index.SelectActions[i];
          actionButtons.push(
            <Tooltip key={i} title={action.Name}>
              <IconButton
                color="secondary"
                className={classes.marginButton}
                onClick={e => onActionClick(e, action.Name)}>
                {icon_utils.getIcon(action.Icon)}
              </IconButton>
            </Tooltip>,
          );
        }
      }
    } else {
      if (index.Actions != null) {
        for (let i = 0, len = index.Actions.length; i < len; i++) {
          const action = index.Actions[i];
          actionButtons.push(
            <Tooltip key={i} title={action.Name}>
              <IconButton
                color="primary"
                className={classes.marginButton}
                onClick={e => onActionClick(e, action.Name)}>
                {icon_utils.getIcon(action.Icon)}
              </IconButton>
            </Tooltip>,
          );
        }
      }
    }

    return (
      <Toolbar>
        <Grid container={true} justify="space-between" spacing={1}>
          <Grid item={true}>
            <div>
              <FormControl className={classes.margin}>
                <Input
                  id="input-with-icon-adornment"
                  placeholder="Search"
                  onChange={onChangeSearchInput}
                  startAdornment={
                    <InputAdornment position="start">
                      <SearchIcon />
                    </InputAdornment>
                  }
                />
              </FormControl>

              <IconButton
                value="bold"
                aria-label="bold"
                className={classes.leftButton}>
                <Badge badgeContent={4} color="secondary">
                  <ErrorOutlineIcon />
                </Badge>
              </IconButton>
              <IconButton
                value="bold"
                aria-label="bold"
                className={classes.leftButton}>
                <Badge badgeContent={4} color="secondary">
                  <HighlightOffOutlinedIcon />
                </Badge>
              </IconButton>
            </div>
          </Grid>
          <Grid item={true}>{actionButtons}</Grid>

          <Grid item={true}>
            <TablePagination
              count={count}
              rowsPerPage={rowsPerPage}
              page={page}
              onChangePage={onChangePage}
              onChangeRowsPerPage={onChangeRowsPerPage}
            />
          </Grid>
        </Grid>
      </Toolbar>
    );
  }
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
)(withStyles(styles, {withTheme: true})(TableToolbar));
