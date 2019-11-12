import * as mr from 'mingru';
import tables from './models/models';
import actions from './actions/actions';

(async () => {
  const dialect = new mr.MySQL();
  const builder = new mr.Builder(dialect, '../da/', {
    cleanBuild: true, // Cleans build directory on each build
  });

  await builder.build(async () => {
    builder.buildActions(actions);
    builder.buildCreateTableSQLFiles(tables);
  });
})();
