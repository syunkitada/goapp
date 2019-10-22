import * as React from 'react';
import {connect} from 'react-redux';

import {Theme} from '@material-ui/core/styles/createMuiTheme';
import createStyles from '@material-ui/core/styles/createStyles';
import withStyles, {
  StyleRules,
  WithStyles,
} from '@material-ui/core/styles/withStyles';

import AppBar from '@material-ui/core/AppBar';
import Tab from '@material-ui/core/Tab';
import CoreTabs from '@material-ui/core/Tabs';
import Typography from '@material-ui/core/Typography';

import logger from '../../../../lib/logger';

const styles = (theme: Theme): StyleRules =>
  createStyles({
    root: {
      backgroundColor: theme.palette.background.paper,
      flexGrow: 1,
      width: '100%',
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
  public state = {
    tabId: null,
    tabRoute: null,
  };

  public render() {
    const {classes, render, routes, data, index} = this.props;
    logger.info('Tabs', 'render()', routes);

    const route = routes[routes.length - 1];

    const tabs: any[] = [];
    let tabContainer: any = null;
    let tabId = 0;
    for (let i = 0, len = index.Tabs.length; i < len; i++) {
      const tab = index.Tabs[i];
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

  private handleChange = (event, tabId) => {
    console.log('DEBUG tab change');
    const {index, routes} = this.props;
    const route = routes[routes.length - 1];
    const splitedPath = route.match.path.split('/');
    let splitedUrl = route.match.url.split('/');
    splitedUrl[splitedPath.indexOf(':' + index.TabParam)] =
      index.Tabs[tabId].Name;

    const lastIndex =
      route.match.path.split(index.Route)[0].split('/').length +
      index.Route.split('/').length;
    splitedUrl = splitedUrl.slice(0, lastIndex - 1);

    route.history.push(splitedUrl.join('/'));
  };
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
