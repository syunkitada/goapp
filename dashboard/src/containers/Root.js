import React, { Component } from 'react';
import { Provider} from 'react-redux';
import { BrowserRouter, Route, Link, Switch, Redirect, withRouter } from 'react-router-dom';

import configureStore from '../store/configureStore'

import NotFound from '../components/NotFound'
import Login from './Login'
import Logout from './Logout'
import AuthRoute from './AuthRoute'
import Dashboard from '../components/Dashboard'

import Auth from '../apps/auth'
import Home from '../apps/home'
import Chat from '../apps/chat'
import Datacenter from '../apps/datacenter'
import Ticket from '../apps/ticket'
import Wiki from '../apps/wiki'
import ProjectHome from '../apps/project/home'
import ProjectResource from '../apps/project/resource'

const store = configureStore()

export default class Root extends Component {
  render() {
    return (
      <Provider store={store}>
        <Auth>
          <BrowserRouter>
            <Switch>
              <Route path="/login" component={Login} />
              <AuthRoute path="/Home" component={Home} />
              <AuthRoute path="/Chat" component={Chat} />
              <AuthRoute path="/Datacenter" component={Datacenter} />
              <AuthRoute path="/Ticket" component={Ticket} />
              <AuthRoute path="/Wiki" component={Wiki} />
              <AuthRoute path="/Project/:project/Home" component={ProjectHome} />
              <AuthRoute path="/Project/:project/Resource" component={ProjectResource} />
              <Route component={NotFound} />
            </Switch>
          </BrowserRouter>
        </Auth>
      </Provider>
    )
  }
}
