import * as mr from 'mingru';
import tables from './models/models.js';
import actions from './actions/actions.js';

(async () => {
  const dialect = new mr.MySQL();
  const outDir = '../da/';
  const builder = new mr.Builder(dialect, outDir, {
    cleanOutDir: true, // Cleans build directory on each build.
    tsOutDir: outDir + 'ts',
  });

  await builder.buildAsync(async () => {
    await builder.buildActionsAsync(actions);
    await builder.buildCreateTableSQLFilesAsync(tables);
  });
})();
