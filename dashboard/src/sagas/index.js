import { all } from 'redux-saga/effects';
import auth from '../apps/auth/sagas';
import home from '../apps/home/sagas';
import resourcePhysical from '../apps/project/resource/physical/sagas';
import monitor from '../apps/project/monitor/sagas';
import service from '../apps/service/sagas';

// redux-sagaのMiddlewareが rootSaga タスクを起動する

// function*は、Generatorオブジェクトを返すジェネレータ関数
// ジェネレーター関数を呼び出しても関数は直ぐには実行されません。代わりに、関数のためのiterator オブジェクトが返す
export default function* rootSaga() {
  yield all([
    auth.watchSyncState(),
    auth.watchLogin(),
    auth.watchLogout(),
    home.watchSyncState(),
    resourcePhysical.watchGetIndex(),
    resourcePhysical.watchGetDatacenterIndex(),
    monitor.watchSyncState(),
    monitor.watchSyncIndexState(),
    service.watchGetIndex(),
  ])
}
