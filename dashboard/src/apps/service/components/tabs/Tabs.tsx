import * as React from 'react';
import {connect} from 'react-redux';

import {Theme} from '@material-ui/core/styles/createMuiTheme';
import withStyles, {
  WithStyles,
  StyleRules,
} from '@material-ui/core/styles/withStyles';
import createStyles from '@material-ui/core/styles/createStyles';

import Typography from '@material-ui/core/Typography';
import CoreTabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import AppBar from '@material-ui/core/AppBar';

import logger from '../../../../lib/logger';

const styles = (theme: Theme): StyleRules =>
  createStyles({
    root: {
      flexGrow: 1,
      width: '100%',
      backgroundColor: theme.palette.background.paper,
    },
  });

interface ITabs extends WithStyles<typeof styles> {
  classes;
  render;
  routes;
  data;
  index;
}

class Tabs extends React.Component<ITabs> {
  state = {
    tabRoute: null,
    tabId: null,
  };

  handleChange = (event, tabId) => {
    const {index, routes} = this.props;
    let route = routes[routes.length - 1];
    let splitedPath = route.match.path.split('/');
    let splitedUrl = route.match.url.split('/');
    splitedUrl[splitedPath.indexOf(':' + index.TabParam)] =
      index.Tabs[tabId].Name;

    let lastIndex =
      route.match.path.split(index.Route)[0].split('/').length +
      index.Route.split('/').length;
    splitedUrl = splitedUrl.slice(0, lastIndex - 1);

    route.history.push(splitedUrl.join('/'));
  };

  render() {
    const {classes, render, routes, data, index} = this.props;
    logger.info(['Tabs', 'render()', routes]);

    let route = routes[routes.length - 1];

    let tabs: any[] = [];
    let tabContainer: any = null;
    let tabId = 0;
    for (let i = 0, len = index.Tabs.length; i < len; i++) {
      let tab = index.Tabs[i];
      if (route.match.params[index.TabParam] === tab.Name) {
        tabId = i;
        tabContainer = (
          <Typography component="div">{render(routes, data, tab)}</Typography>
        );
      }
      tabs.push(<Tab key={tab.Name} label={tab.Name} />);
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
            scrollButtons="auto">
            {tabs}
          </CoreTabs>
        </AppBar>
        {tabContainer}
      </div>
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
)(withStyles(styles)(Tabs));
