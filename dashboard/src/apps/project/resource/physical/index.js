import React, {Component} from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { Route } from 'react-router-dom';

import { withStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import DashboardIcon from '@material-ui/icons/Dashboard';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import ExpansionPanel from '@material-ui/core/ExpansionPanel';
import ExpansionPanelSummary from '@material-ui/core/ExpansionPanelSummary';
import ExpansionPanelDetails from '@material-ui/core/ExpansionPanelDetails';
import Typography from '@material-ui/core/Typography';

import HostTable from './components/HostTable'
import IndexTable from './components/IndexTable'
import Index from './components/Index'
import Dashboard from '../../../../components/Dashboard'
import actions from '../../../../actions'


const styles = theme => ({
  nested: {
    paddingLeft: theme.spacing.unit * 4,
  },
});

class ProjectResourcePhysical extends Component {
  state = {
    expanded: "IndexPanel",
  };

  handleChange = panel => (event, expanded) => {
    this.setState({
      expanded: expanded ? panel : false,
    });
  };

  componentWillMount() {
    const {match, getIndex} = this.props
    getIndex(match.params.project)
  }

  render() {
    const {classes, match, auth, resourcePhysical} = this.props
    const { expanded } = this.state;
    if (!auth.user) {
      return null
    }

    const projectService = auth.user.Authority.ProjectServiceMap[match.params.project]

    if (!resourcePhysical.index) {
      console.log("!index.index")
      return (
        <Dashboard projectService={projectService} match={match}>
          <div>
            <h2>Resource.Physical</h2>
          </div>
        </Dashboard>
      );
    } else {
      console.log(resourcePhysical.index)
      return (
        <Dashboard projectService={projectService} match={match}>
          <Typography variant="display1">
            Resource.Physical
          </Typography>
          <Route exact path={match.path} render={props =>
            <Index {...props} index={resourcePhysical.index} />
          }/>
          <Route path={`${match.path}/:index`} render={props =>
            <Index {...props} index={resourcePhysical.index} />
          } />
        </Dashboard>
      );
    }
  }
}

ProjectResourcePhysical.propTypes = {
  classes: PropTypes.object.isRequired,
  auth: PropTypes.object.isRequired,
  resourcePhysical: PropTypes.object.isRequired,
  getIndex: PropTypes.func.isRequired
}

function mapStateToProps(state, ownProps) {
  const auth = state.auth
  const resourcePhysical = state.resourcePhysical

  return {
    auth: auth,
    resourcePhysical: resourcePhysical,
  }
}

function mapDispatchToProps(dispatch, ownProps) {
  return {
    getIndex: (projectName) => {
      dispatch(actions.resourcePhysical.resourcePhysicalGetIndex(projectName));
    }
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(withStyles(styles)(ProjectResourcePhysical))
