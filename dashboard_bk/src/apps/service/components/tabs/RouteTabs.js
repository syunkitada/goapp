import React, {Component} from 'react';
import { connect } from 'react-redux';
import { Route } from 'react-router-dom';
import PropTypes from 'prop-types';

import { withStyles } from '@material-ui/core/styles';

import Tabs from './Tabs'

class RouteTabs extends Component {
  render() {
    const { classes, render, routes, data, index } = this.props

    let beforeRoute = routes.slice(-1)[0]

    return (
      <div className={classes.root}>
      {index.Tabs.map((v) =>
        <Route exact={v.Route === ""} path={beforeRoute.match.path + v.Route} key={v.Name} render={props => {
          const newRoutes = routes.slice(0)
          newRoutes.push(props)
          return <Tabs render={render} routes={newRoutes} data={data} index={index} root={v} />
        }
        } />
      )}
      </div>
    );
  }
}

const styles = theme => ({
  root: {
    width: '100%',
  },
});

RouteTabs.propTypes = {
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
)(withStyles(styles)(RouteTabs));
