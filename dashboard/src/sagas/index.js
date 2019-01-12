import { all } from 'redux-saga/effects';
import auth from '../apps/auth/sagas';
import home from '../apps/home/sagas';
import resource from '../apps/project/resource/sagas';
import monitor from '../apps/project/monitor/sagas';

// redux-sagaのMiddlewareが rootSaga タスクを起動する

// function*は、Generatorオブジェクトを返すジェネレータ関数
// ジェネレーター関数を呼び出しても関数は直ぐには実行されません。代わりに、関数のためのiterator オブジェクトが返す
export default function* rootSaga() {
  yield all([
    auth.watchSyncState(),
    auth.watchLogin(),
    auth.watchLogout(),
    home.watchSyncState(),
    resource.watchSyncState(),
    monitor.watchSyncState(),
  ])
}
