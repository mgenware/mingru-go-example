import * as dd from 'dd-models';
import t from './dept';

export class DeptTA extends dd.TA {
  selectPage = dd.selectPage().orderBy(t.id);
}

export default dd.ta(t, DeptTA);
