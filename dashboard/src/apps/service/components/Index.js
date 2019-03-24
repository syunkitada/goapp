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

import actions from '../../../actions'
import IndexTable from './tables/IndexTable'
import RoutePanels from './panels/RoutePanels'
import RouteTabs from './tabs/RouteTabs'


const styles = theme => ({
  nested: {
    paddingLeft: theme.spacing.unit * 4,
  },
});

function renderIndex(match, data, index) {
  console.log("DEBUG: Index.renderIndex: ", index.Kind, index.Name)
  switch(index.Kind) {
    case "Msg":
      return <div>{index.Name}</div>
    case "RoutePanels":
      return <RoutePanels render={renderIndex} match={match} data={data} index={index} />
    case "RouteTabs":
      return <RouteTabs render={renderIndex} match={match} data={data} index={index} />
    case "Table":
      return <IndexTable match={match} columns={index.Columns} data={data[index.DataKey]} />
    default:
      return <div>Unsupported Kind: {index.Kind}</div>
  }
}

class Index extends Component {
  componentWillMount() {
    console.log("DEBUG: Index will mount")
    const {match, getIndex} = this.props
    getIndex(match.params)
  }


  render() {
    const {classes, match, service, serviceName, projectName, auth, index, getIndex} = this.props
		console.log("reder Index", projectName, serviceName)

    if (service.serviceName != serviceName || service.projectName != projectName) {
      getIndex(match.params)
      return null
    }

    let state = null
    if (projectName) {
      state = service.projectServiceMap[projectName][serviceName]
    } else {
      state = service.serviceMap[serviceName]
    }

    console.log(state)
    if (state.isFetching) {
      return <div>Fetching...</div>
    }

    let html = renderIndex(match, state.Data, state.Index)

    return (
      <div>
        { html }
      </div>
    );
  }
}

Index.propTypes = {
  classes: PropTypes.object.isRequired,
  match: PropTypes.object.isRequired,
  auth: PropTypes.object.isRequired,
}

function mapStateToProps(state, ownProps) {
  console.log("DEBUG mapStateToProps")
  console.log(ownProps)
  const match = ownProps.match
  const auth = state.auth
  const service = state.service

  return {
    match: match,
    serviceName: match.params.service,
    projectName: match.params.project,
    auth: auth,
    service: service,
  }
}

function mapDispatchToProps(dispatch, ownProps) {
  return {
    getIndex: (params) => {
      dispatch(actions.service.serviceGetIndex(params));
    }
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(withStyles(styles)(Index))
