import * as dd from 'dd-models';

class Employee extends dd.Table {
  id = dd.pk(dd.int()).setDBName('emp_no');
  first_name = dd.varChar(50);
  last_name = dd.varChar(50);
  gender = dd.varChar(10);
  birth_date = dd.date();
  hire_date = dd.date();
}

export default dd.table(Employee, 'employees');
