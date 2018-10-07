import { combineReducers } from 'redux'

import auth from '../apps/auth/modules';
import resource from '../apps/project/resource/modules';

export default {
  auth,
  resource,
};
