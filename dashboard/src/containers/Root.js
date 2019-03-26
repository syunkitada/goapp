import React, { Component } from 'react';
import { Provider} from 'react-redux';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

import configureStore from '../store/configureStore'

import NotFound from '../components/NotFound'
import AuthRoute from './AuthRoute'

import Auth from '../apps/auth'
import Login from '../apps/auth/components/Login'
import Service from '../apps/service'

const store = configureStore()

export default class Root extends Component {
  render() {
    return (
      <Provider store={store}>
        <Auth>
          <BrowserRouter>
            <Switch>
              <Route path="/login" component={Login} />
              <AuthRoute path="/Service/:service" component={Service} />
              <AuthRoute path="/Project/:project/:service" component={Service} />
              <Route component={NotFound} />
            </Switch>
          </BrowserRouter>
        </Auth>
      </Provider>
    )
  }
}
