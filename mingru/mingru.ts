import * as mr from 'mingru';
// Import model actions
import employeeTA from './models/employeeTA';
import deptTA from './models/deptTA';

(async () => {
  const actions = [employeeTA, deptTA];
  const dialect = new mr.MySQL();
  // Build Go code to '../da/` directory
  await mr.build(actions, dialect, '../da/', { cleanBuild: true });
})();
