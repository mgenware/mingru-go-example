import * as dd from 'dd-models';
import t from './employee';

const employeeColumns = [
  t.id,
  t.firstName,
  t.lastName,
  t.gender,
  t.birthDate,
  t.hireDate,
];

export class EmployeeTA extends dd.TA {
  // Select an employee by ID
  selectByID = dd.select(...employeeColumns).byID();
  // Select all employees
  selectAll = dd.selectAll(...employeeColumns);
  // Select an employee birth date by ID
  selectSig = dd.selectField(t.birthDate).byID();

  // Update an employee name by ID
  updateName = dd
    .updateOne()
    .setInputs(t.firstName, t.lastName)
    .byID();

  // Delete an employee by ID
  deleteByID = dd.deleteOne().byID();

  // Delete all employees by a criteria
  deleteByBirthDate = dd.deleteSome().where(t.birthDate.isEqualToInput());

  // Delete all employees
  deleteAll = dd.unsafeDeleteAll();

  // Insert a new employee
  insertUser = dd.insertOne().setInputs(...employeeColumns);
}

export default dd.ta(t, EmployeeTA);
