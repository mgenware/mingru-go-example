import * as dd from 'dd-models';

class User extends dd.Table {
  id = dd.pk();
  name = dd.varChar(100);
  sig = dd.text().nullable;
  age = dd.int();
}

export default dd.table(User);
