import * as dd from 'mingru-models';

class Dept extends dd.Table {
  id = dd.pk(dd.char(4)).setDBName('dept_no');
  name = dd.varChar(40).setDBName('dept_name');
}

export default dd.table(Dept, 'departments');
