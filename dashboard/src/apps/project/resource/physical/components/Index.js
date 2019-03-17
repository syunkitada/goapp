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
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import AppBar from '@material-ui/core/AppBar';

import IndexTable from './IndexTable'
import HostTable from './HostTable'
import LogTable from './LogTable'
import actions from '../../../../../actions'


const styles = theme => ({
  nested: {
    paddingLeft: theme.spacing.unit * 4,
  },
});

function TabContainer(props) {
  return (
    <Typography component="div" style={{ padding: 8 * 3 }}>
      {props.children}
    </Typography>
  );
}

class Index extends Component {
  state = {
    expanded: null,
    tabId: 0,
  };

  handleChange = panel => (event, expanded) => {
    this.setState({
      expanded: expanded ? panel : false,
    });
  };

  handleChangeTab = (event, tabId) => {
    this.setState({ tabId });
  };

  componentWillMount() {
    const {match, getDatacenterIndex} = this.props
    if (match.params.index) {
      getDatacenterIndex(match.params.project, match.params.index)
    }
  }

  render() {
    const {classes, match, auth, index, datacenterIndex} = this.props
    let { expanded, tabId } = this.state;

    if (!auth.user) {
      return null
    }

    const projectService = auth.user.Authority.ProjectServiceMap[match.params.project]

    if (!index.Datacenters) {
      console.log("!monitor.monitor")
      return (
        <div>
          NotFound
        </div>
      );
    } else {
      var selectedIndexHtml = ""
      if (match.params.index) {
        selectedIndexHtml = ": " + match.params.index
        if (expanded === null) {
          expanded = "HostPanel"
        }
      } else {
        if (expanded === null) {
          expanded = "IndexPanel"
        }
      }

      console.log("DEBUG Index")
      console.log(datacenterIndex)

      return (
        <div>
          <ExpansionPanel expanded={expanded === 'IndexPanel'} onChange={this.handleChange('IndexPanel')}>
            <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
              <Typography variant="title">
                Index Table
                {selectedIndexHtml}
              </Typography>
            </ExpansionPanelSummary>
            <ExpansionPanelDetails>
              <IndexTable match={match} index={index} />
            </ExpansionPanelDetails>
          </ExpansionPanel>

          <Paper>
            <AppBar position="static" color="default">
              <Tabs
                value={tabId}
                onChange={this.handleChangeTab}
                indicatorColor="primary"
                textColor="primary"
                variant="scrollable"
                scrollButtons="auto"
              >
                <Tab label="Resources" />
                <Tab label="Racks" />
                <Tab label="Floors" />
              </Tabs>
            </AppBar>
            {tabId === 0 && <TabContainer>
              PhysicalResources
            </TabContainer>}
            {tabId === 1 && <TabContainer>
              Racks
            </TabContainer>}
            {tabId === 2 && <TabContainer>
              Floors
            </TabContainer>}
          </Paper>

        </div>
      );
    }
  }
}

Index.propTypes = {
  classes: PropTypes.object.isRequired,
  auth: PropTypes.object.isRequired,
  index: PropTypes.object.isRequired,
  getDatacenterIndex: PropTypes.func.isRequired
}

function mapStateToProps(state, ownProps) {
  const auth = state.auth
  const index = ownProps.index
  const datacenterIndex = state.resourcePhysical.datacenterIndex

  return {
    auth: auth,
    index: index,
    datacenterIndex: datacenterIndex,
  }
}

function mapDispatchToProps(dispatch, ownProps) {
  return {
    getDatacenterIndex: (projectName, datacenterName) => {
      dispatch(actions.resourcePhysical.resourcePhysicalGetDatacenterIndex(projectName, datacenterName));
    }
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(withStyles(styles)(Index))
