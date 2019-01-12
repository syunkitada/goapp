import { combineReducers } from 'redux'

import auth from '../apps/auth/modules';
import resource from '../apps/project/resource/modules';
import monitor from '../apps/project/monitor/modules';

export default {
  auth,
  resource,
  monitor,
};
