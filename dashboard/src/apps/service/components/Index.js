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
import IndexTable from './IndexTable'


const styles = theme => ({
  nested: {
    paddingLeft: theme.spacing.unit * 4,
  },
});


class Index extends Component {
  state = {
		expandedMap: {}
	};

  handleChange = (group, expanded) => {
		let { expandedMap } = this.state
		if (expandedMap[group] == expanded) {
			expandedMap[group] = null
		} else {
			expandedMap[group] = expanded
		}
    this.setState({
      expandedMap: expandedMap,
    });
  };

	renderIndex = (match, data, group, index) => {
		let html = []
		switch(index.Kind) {
			case "Msg":
				return (
					<div>{index.Name}</div>
				);
			case "Panels":
				console.log("DEBUGllalala Panels")
				console.log(group)
				for (const v of index.Panels) {
					console.log(v.Name)
					html.push(
						<ExpansionPanel key={v.Name} expanded={this.state.expandedMap[group] === v.Name} onChange={() => this.handleChange(group, v.Name)}>
							<ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
								<Typography variant="title">
									Index Table
								</Typography>
							</ExpansionPanelSummary>
							<ExpansionPanelDetails>
								{this.renderIndex(match, data, v.Name, v)}
							</ExpansionPanelDetails>
						</ExpansionPanel>
					)
				};
				break;
			case "Table":
				return (
					<IndexTable match={match} columns={index.Columns} data={data[index.DataKey]} />
				);
		}

		return (
			<div>
				{html}
			</div>
		)
	};

  render() {
    const {classes, match, auth, index} = this.props

    if (index.Index == null) {
      return null
    }
    let html = this.renderIndex(match, index.Data, index.Index.Name, index.Index)

    return (
      <div>
        { html }
      </div>
    );
  }
}

Index.propTypes = {
  classes: PropTypes.object.isRequired,
  auth: PropTypes.object.isRequired,
  index: PropTypes.object.isRequired,
}

function mapStateToProps(state, ownProps) {
  const auth = state.auth
  const index = state.service.index

  return {
    auth: auth,
    index: index,
  }
}

function mapDispatchToProps(dispatch, ownProps) {
  return {}
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(withStyles(styles)(Index))
