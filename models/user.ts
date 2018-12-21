import * as dd from 'dd-models';

class User extends dd.Table {
  id = dd.pk();
  name = dd.varChar(200);
  sig = dd.text().nullable;
}

export default dd.table(User);
