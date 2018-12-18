import * as mr from 'mingru';
import actions from './actions';

(async () => {
  await mr.build(actions, new mr.MySQL(), './build/');
})();
