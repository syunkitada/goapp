import actionCreatorFactory from 'typescript-fsa';

const actionCreator = actionCreatorFactory();

export const authSyncState = actionCreator('AUTH_SYNC_STATE');

export const authLogin = actionCreator<{username: string; password: string}>(
  'AUTH_LOGIN',
);
export const authLoginSuccess = actionCreator<{username: any; authority: any}>(
  'AUTH_LOGIN_SUCCESS',
);
export const authLoginFailure = actionCreator<{error: any}>(
  'AUTH_LOGIN_FAILURE',
);

export const authLogout = actionCreator('AUTH_LOGOUT');
export const authLogoutSuccess = actionCreator('AUTH_LOGOUT_SUCCESS');
export const authLogoutFailure = actionCreator<{error: any}>(
  'AUTH_LOGOUT_FAILURE',
);
