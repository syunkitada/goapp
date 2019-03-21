import { combineReducers } from 'redux'

import auth from '../apps/auth/modules';
import service from '../apps/service/modules';
import resourcePhysical from '../apps/project/resource/physical/modules';
import resourceVirtual from '../apps/project/resource/virtual/modules';
import monitor from '../apps/project/monitor/modules';
import sort_utils from './sort_utils';

export default {
  auth,
  service,
  resourcePhysical,
  resourceVirtual,
  monitor,
  sort_utils,
};
