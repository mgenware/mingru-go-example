import * as dd from 'mingru-models';
import t from './employee';

export class EmployeeTA extends dd.TA {
  // Select an employee by ID.
  selectByID = dd.select().byID();
  // Select all employees.
  selectAll = dd.selectRows().orderByAsc(t.hireDate);
  // Select all employees with a limit.
  selectAllWithLimit = dd
    .selectRows()
    .limit()
    .orderByAsc(t.hireDate);
  // Select paged employees.
  selectPaged = dd.selectPage().orderByAsc(t.hireDate);
  // Select an employee birth date by ID.
  selectSig = dd.selectField(t.birthDate).byID();

  // Update an employee name by ID.
  updateName = dd
    .updateOne()
    .setInputs(t.firstName, t.lastName)
    .byID();

  // Delete an employee by ID.
  deleteByID = dd.deleteOne().byID();

  // Delete all employees by some criteria.
  deleteByBirthDate = dd.deleteSome().where(t.birthDate.isEqualToInput());

  // Insert a new employee.
  insertUser = dd.insertOne().setInputs();
}

export default dd.ta(t, EmployeeTA);
