import * as React from 'react';
import {connect} from 'react-redux';

import {Theme} from '@material-ui/core/styles/createMuiTheme';
import withStyles, {
  WithStyles,
  StyleRules,
} from '@material-ui/core/styles/withStyles';
import createStyles from '@material-ui/core/styles/createStyles';

import Toolbar from '@material-ui/core/Toolbar';
import {lighten} from '@material-ui/core/styles/colorManipulator';

import {fade} from '@material-ui/core/styles/colorManipulator';

import Input from '@material-ui/core/Input';
import InputAdornment from '@material-ui/core/InputAdornment';
import FormControl from '@material-ui/core/FormControl';
import Tooltip from '@material-ui/core/Tooltip';
import SearchIcon from '@material-ui/icons/Search';
import Button from '@material-ui/core/Button';
import IconButton from '@material-ui/core/IconButton';

import Grid from '@material-ui/core/Grid';

import icon_utils from '../../../../modules/icon_utils';
import TablePagination from './TablePagination';

const styles = (theme: Theme): StyleRules =>
  createStyles({
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
    buttonMargin: {
      marginTop: theme.spacing.unit * 2,
      marginBottom: theme.spacing.unit * 2,
    },
    highlight:
      theme.palette.type === 'light'
        ? {
            color: theme.palette.secondary.main,
            backgroundColor: lighten(theme.palette.secondary.light, 0.85),
          }
        : {
            color: theme.palette.text.primary,
            backgroundColor: theme.palette.secondary.dark,
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
    search: {
      position: 'relative',
      borderRadius: theme.shape.borderRadius,
      backgroundColor: fade(theme.palette.common.white, 0.15),
      '&:hover': {
        backgroundColor: fade(theme.palette.common.white, 0.25),
      },
      marginRight: theme.spacing.unit * 2,
      marginLeft: 0,
      width: '100%',
      [theme.breakpoints.up('sm')]: {
        marginLeft: theme.spacing.unit * 3,
        width: 'auto',
      },
    },
    searchIcon: {
      width: theme.spacing.unit * 9,
      height: '100%',
      position: 'absolute',
      pointerEvents: 'none',
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'center',
    },
  });

interface ITableToolbar extends WithStyles<typeof styles> {
  index;
  onChangeSearchInput;
  count;
  numSelected;
  rowsPerPage;
  page;
  onChangePage;
  onChangeRowsPerPage;
  onActionClick;
}

class TableToolbar extends React.Component<ITableToolbar> {
  render() {
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
      onActionClick,
    } = this.props;

    const actionButtons: any[] = [];
    if (numSelected > 0) {
      if (index.SelectActions != null) {
        actionButtons.push(
          <Button key={-1} color="secondary">
            {numSelected} selected
          </Button>,
        );
        for (let i = 0, len = index.SelectActions.length; i < len; i++) {
          let action = index.SelectActions[i];
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
          let action = index.Actions[i];
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
        <Grid container justify="space-between" spacing={24}>
          <Grid item>
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
            </div>
          </Grid>
          <Grid item>{actionButtons}</Grid>

          <Grid item>
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
