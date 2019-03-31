import React, {Component} from 'react';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';

import { withStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import ExpansionPanel from '@material-ui/core/ExpansionPanel';
import ExpansionPanelSummary from '@material-ui/core/ExpansionPanelSummary';
import ExpansionPanelDetails from '@material-ui/core/ExpansionPanelDetails';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';

import actions from '../../../../actions'


class ExpansionPanels extends Component {
  state = {
		expanded: null,
		expandedUrl: null,
	}

	componentWillMount() {
    const { routes, index } = this.props
    let route = routes.slice(-1)[0]
    let beforeRoute = routes.slice(-2)[0]

		for (let i = 0, len = index.Panels.length; i < len; i++) {
			let panel = index.Panels[i];
			if (route.match.path === beforeRoute.match.path + panel.Route) {
				if (panel.GetQueries) {
					this.props.getQueries(panel.GetQueries, route.match.params)
				}
				break
			}
		}
	}

  handleChange = (expandedPath, expandedUrl) => {
    this.setState({
      expanded: expandedPath,
      expandedUrl: expandedUrl,
    });
  };

  render() {
    const { classes, render, routes, data, index } = this.props
    let { expanded, expandedUrl } = this.state;
    console.log("ExpansionPanels.render()")

    let route = routes.slice(-1)[0]
    let beforeRoute = routes.slice(-2)[0]

    if (expanded === null) {
      expanded = route.match.path
    }

    if (expandedUrl !== null && expanded !== route.match.path && expandedUrl !== route.match.url) {
      expanded = route.match.path
    }

    return (
      <div className={classes.root}>
      { index.Panels.map((p) =>
            <ExpansionPanel key={p.Name} expanded={(() => {
              return expanded === beforeRoute.match.path + p.Route
            })()} onChange={() => this.handleChange(beforeRoute.match.path + p.Route, route.match.url)}>
              <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
                <Typography variant="title">
                  {p.Name} {route.match.params[p.Subname]}
                </Typography>
              </ExpansionPanelSummary>
              <ExpansionPanelDetails style={{padding: 0}}>
                {render(routes, data, p)}
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
		flexGrow: 1,
    width: '100%',
    backgroundColor: theme.palette.background.paper,
  },
});

ExpansionPanels.propTypes = {
  classes: PropTypes.object.isRequired,
  render: PropTypes.func.isRequired,
  routes: PropTypes.array.isRequired,
  data: PropTypes.object.isRequired,
  index: PropTypes.object.isRequired,
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
      dispatch(actions.service.serviceStartBackgroundSync(querys, params));
    }
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(withStyles(styles)(ExpansionPanels));
