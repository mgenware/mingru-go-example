import * as mr from 'mingru';
import userTA from './models/userTA';

(async () => {
  await mr.build([userTA], new mr.MySQL(), './build/');
})();
