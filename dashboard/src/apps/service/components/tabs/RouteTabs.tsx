import * as React from 'react';
import {connect} from 'react-redux';
import {Route} from 'react-router-dom';

import {Theme} from '@material-ui/core/styles/createMuiTheme';
import withStyles, {
  WithStyles,
  StyleRules,
} from '@material-ui/core/styles/withStyles';
import createStyles from '@material-ui/core/styles/createStyles';

import Tabs from './Tabs';

const styles = (theme: Theme): StyleRules =>
  createStyles({
    root: {
      width: '100%',
    },
  });

interface IRouteTabs extends WithStyles<typeof styles> {
  render;
  routes;
  data;
  index;
}

class RouteTabs extends React.Component<IRouteTabs> {
  render() {
    const {classes, render, routes, data, index} = this.props;

    let beforeRoute = routes.slice(-1)[0];

    return (
      <div className={classes.root}>
        {index.Tabs.map(v => (
          <Route
            exact={v.Route === ''}
            path={beforeRoute.match.path + v.Route}
            key={v.Name}
            render={props => {
              const newRoutes = routes.slice(0);
              newRoutes.push(props);
              return (
                <Tabs
                  render={render}
                  routes={newRoutes}
                  data={data}
                  index={index}
                  root={v}
                />
              );
            }}
          />
        ))}
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
)(withStyles(styles)(RouteTabs));
