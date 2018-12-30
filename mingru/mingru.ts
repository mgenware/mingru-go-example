import * as mr from 'mingru';
import userTA from './models/userTA';

(async () => {
  const actions = [userTA];
  await mr.build(actions, new mr.MySQL(), '../da/');
})();
