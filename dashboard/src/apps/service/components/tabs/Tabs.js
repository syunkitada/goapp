import React, {Component} from 'react';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';

import { withStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import CoreTabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import AppBar from '@material-ui/core/AppBar';


class Tabs extends Component {
  state = {
    tabRoute: null,
		tabId: null,
	}

  handleChange = (event, tabId) => {
    const { index, routes } = this.props
    let beforeRoute = routes.slice(-2)[0]
    beforeRoute.history.push(beforeRoute.match.url + index.Tabs[tabId].Route)
  };

  render() {
    const { classes, render, routes, data, index } = this.props;

    let route = routes.slice(-1)[0]
    let beforeRoute = routes.slice(-2)[0]

    let tabs = []
    let tabContainer = null
    let tabId = 0;
    for (let i = 0, len = index.Tabs.length; i < len; i++) {
      let tab = index.Tabs[i];
      if (route.match.path === beforeRoute.match.path + tab.Route) {
        tabId = i;
        tabContainer = <Typography component="div">{render(routes, data, tab)}</Typography>
      }
      tabs.push(<Tab key={tab.Name} label={tab.Name} />)
    }

    return (
      <div className={classes.root}>
        <AppBar position="static" color="default">
          <CoreTabs
            value={tabId}
            onChange={this.handleChange}
            indicatorColor="primary"
            textColor="primary"
            variant="scrollable"
            scrollButtons="auto"
          >
            {tabs}
          </CoreTabs>
        </AppBar>
        {tabContainer}
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

Tabs.propTypes = {
  classes: PropTypes.object.isRequired,
  render: PropTypes.func.isRequired,
  routes: PropTypes.array.isRequired,
  data: PropTypes.object.isRequired,
  index: PropTypes.object.isRequired,
};

function mapStateToProps(state, ownProps) {
  return {}
}

function mapDispatchToProps(dispatch, ownProps) {
  return {}
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(withStyles(styles)(Tabs));
