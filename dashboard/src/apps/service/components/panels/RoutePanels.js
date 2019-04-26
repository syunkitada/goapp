import React, {Component} from 'react';
import { connect } from 'react-redux';
import { Route } from 'react-router-dom';
import PropTypes from 'prop-types';

import ExpansionPanels from './ExpansionPanels';
import logger from '../../../../lib/logger';


class RoutePanels extends Component {
  render() {
    const { render, routes, data, index } = this.props

    let beforeRoute = routes.slice(-1)[0]
    logger.info('RoutePanels', 'render()', beforeRoute)
    console.log(index.Panels)

    return (
      <div>
      {index.Panels.map((v) =>
        <Route exact path={beforeRoute.match.path + v.Route} key={v.Name} render={props => {
          console.log(v.Route)
          const newRoutes = routes.slice(0)
          newRoutes.push(props)
          return (
            <ExpansionPanels render={render} routes={newRoutes} index={index} data={data} root={v} />
          )}
        } />
      )}
      </div>
    );
  }
}

RoutePanels.propTypes = {
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
)(RoutePanels);
