import * as dd from 'mingru-models';
import t from './dept';

export class DeptTA extends dd.TA {
  // Selects paged departments.
  selectPage = dd.selectPage().orderByAsc(t.id);
}

export default dd.ta(t, DeptTA);
