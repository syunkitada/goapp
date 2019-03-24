import React, {Component} from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';
import { Route } from 'react-router-dom';

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

import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import ExpansionPanel from '@material-ui/core/ExpansionPanel';
import ExpansionPanelSummary from '@material-ui/core/ExpansionPanelSummary';
import ExpansionPanelDetails from '@material-ui/core/ExpansionPanelDetails';

import actions from '../../../../actions'


class ExpansionPanels extends Component {
  state = {
		expanded: null,
	}

	componentWillMount() {
    const { classes, auth, render, match, data, index, root, route } = this.props
		for (let i = 0, len = index.Panels.length; i < len; i++) {
			let panel = index.Panels[i];
			if (route.match.path === match.path + panel.Route) {
				if (panel.GetQueries) {
					this.props.getQueries(panel.GetQueries, route.match.params)
				}
				break
			}
		}
	}

  handleChange = (expanded) => {
    const { route } = this.props
    if (this.state.expanded == expanded) {
      this.setState({
        expanded: false,
      });
    } else if (this.state.expanded == null && route.match.path == expanded) {
      this.setState({
        expanded: false,
      });
    } else {
      this.setState({
        expanded: expanded,
      });
    }
  };

  render() {
    const { classes, auth, render, match, data, index, root, route } = this.props
    let { expanded } = this.state;
    console.log("DEBUG expansion panels")

    console.log(match)

    if (expanded === null) {
      expanded = route.match.path
    }

    return (
      <div>
      { index.Panels.map((p) =>
            <ExpansionPanel key={p.Name} expanded={
              expanded === match.path + p.Route
            } onChange={() => this.handleChange(match.path + p.Route)}>
              <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
                <Typography variant="title">
                  {p.Name} {route.match.params[p.Subname]}
                </Typography>
              </ExpansionPanelSummary>
              <ExpansionPanelDetails style={{padding: 0}}>
                {render(match, data, p)}
              </ExpansionPanelDetails>
            </ExpansionPanel>
          )
      }
      </div>
    );
  }
}

const styles = theme => ({
  root: {
		padding: 0,
	},
});

ExpansionPanels.propTypes = {
  classes: PropTypes.object.isRequired,
  render: PropTypes.func.isRequired,
};

function mapStateToProps(state, ownProps) {
  const auth = state.auth

  return {
    auth: auth,
  }
}

function mapDispatchToProps(dispatch, ownProps) {
  return {
    getQueries: (querys, params) => {
			console.log("DEBUG getQueries")
      dispatch(actions.service.serviceGetQueries(querys, params));
    }
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(withStyles(styles)(ExpansionPanels));
