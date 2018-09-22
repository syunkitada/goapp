import React, { Component } from 'react';
import { Provider} from 'react-redux';
import { BrowserRouter, Route, Link, Switch, Redirect, withRouter } from 'react-router-dom';

import configureStore from '../store/configureStore'

import About from '../components/About'
import Home from '../components/Home'
import NotFound from '../components/NotFound'
import Login from './Login'
import Logout from './Logout'
import AuthenticatedRoute from './AuthenticatedRoute'
import User from '../components/User'
import App from './App'

const store = configureStore()

export default class Root extends Component {
  render() {
    return (
      <Provider store={store}>
        <App>
        <BrowserRouter>
          <Switch>
            <Route path="/login" component={Login} />
            <AuthenticatedRoute path="/" component={Home} />
            <AuthenticatedRoute path="/user" component={User} />
            <Route component={NotFound} />
          </Switch>
        </BrowserRouter>
        </App>
      </Provider>
    )
  }
}
