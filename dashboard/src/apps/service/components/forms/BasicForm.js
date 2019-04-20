import React, {Component} from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';

import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableRow from '@material-ui/core/TableRow';
import Checkbox from '@material-ui/core/Checkbox';
import Button from '@material-ui/core/Button';
import Menu from '@material-ui/core/Menu';
import MenuItem from '@material-ui/core/MenuItem';
import KeyboardArrowDownIcon from '@material-ui/icons/KeyboardArrowDown';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';

import sort_utils from '../../../../modules/sort_utils'
import icon_utils from '../../../../modules/icon_utils'

class BasicForm extends Component {
  state = {
    order: 'asc',
    orderBy: 0,
    selected: [],
    data: [],
    page: 0,
    rowsPerPage: 5,
    searchRegExp: null,
    anchorEl: null,
		actionTarget: null,
		actionName: null,
  };

  render() {
    const { routes, classes, index, data} = this.props
    const { selected, anchorEl, order, orderBy, rowsPerPage, page, searchRegExp, actionName, actionTarget } = this.state;

    console.log("render BasicForm")
    console.log(index)
    console.log(data)

    return (
      <div className={classes.root}>
        TEST
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

BasicForm.propTypes = {
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
)(withStyles(styles)(BasicForm));
