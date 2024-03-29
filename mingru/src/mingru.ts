import * as mr from 'mingru';
import source from './source.js';

(async () => {
  const outDir = '../da/';
  const builder = new mr.Builder(outDir, {
    cleanOutDir: true, // Cleans build directory on each build.
    tsOutDir: outDir + 'ts',
    createTableSQL: true,
  });

  await builder.build(source);
})();
