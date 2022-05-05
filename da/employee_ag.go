/******************************************************************************************
 * This file was automatically generated by mingru (https://github.com/mgenware/mingru)
 * Do not edit this file manually, your changes will be overwritten.
 ******************************************************************************************/

package da

import (
	"fmt"
	"time"

	"github.com/mgenware/mingru-go-lib"
)

type EmployeeAGType struct {
}

var Employee = &EmployeeAGType{}

// ------------ Actions ------------

func (mrTable *EmployeeAGType) DeleteByBirthDate(mrQueryable mingru.Queryable, birthDate time.Time) (int, error) {
	result, err := mrQueryable.Exec("DELETE FROM `employees` WHERE `birth_date` = ?", birthDate)
	return mingru.GetRowsAffectedIntWithError(result, err)
}

func (mrTable *EmployeeAGType) DeleteByID(mrQueryable mingru.Queryable, id int) error {
	result, err := mrQueryable.Exec("DELETE FROM `employees` WHERE `emp_no` = ?", id)
	return mingru.CheckOneRowAffectedWithError(result, err)
}

func (mrTable *EmployeeAGType) InsertUser(mrQueryable mingru.Queryable, id int, firstName string, lastName string, gender string, birthDate time.Time, hireDate time.Time) error {
	_, err := mrQueryable.Exec("INSERT INTO `employees` (`emp_no`, `first_name`, `last_name`, `gender`, `birth_date`, `hire_date`) VALUES (?, ?, ?, ?, ?, ?)", id, firstName, lastName, gender, birthDate, hireDate)
	return err
}

type EmployeeAGSelectAllResult struct {
	BirthDate time.Time
	FirstName string
	Gender    string
	HireDate  time.Time
	ID        int
	LastName  string
}

func (mrTable *EmployeeAGType) SelectAll(mrQueryable mingru.Queryable, page int, pageSize int) ([]EmployeeAGSelectAllResult, bool, error) {
	if page <= 0 {
		err := fmt.Errorf("Invalid page %v", page)
		return nil, false, err
	}
	if pageSize <= 0 {
		err := fmt.Errorf("Invalid page size %v", pageSize)
		return nil, false, err
	}
	limit := pageSize + 1
	offset := (page - 1) * pageSize
	max := pageSize
	rows, err := mrQueryable.Query("SELECT `emp_no`, `first_name`, `last_name`, `gender`, `birth_date`, `hire_date` FROM `employees` ORDER BY `hire_date` LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, false, err
	}
	result := make([]EmployeeAGSelectAllResult, 0, limit)
	itemCounter := 0
	defer rows.Close()
	for rows.Next() {
		itemCounter++
		if itemCounter <= max {
			var item EmployeeAGSelectAllResult
			err = rows.Scan(&item.ID, &item.FirstName, &item.LastName, &item.Gender, &item.BirthDate, &item.HireDate)
			if err != nil {
				return nil, false, err
			}
			result = append(result, item)
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, false, err
	}
	return result, itemCounter > len(result), nil
}

type EmployeeAGSelectAllWithLimitResult struct {
	BirthDate time.Time
	FirstName string
	Gender    string
	HireDate  time.Time
	ID        int
	LastName  string
}

func (mrTable *EmployeeAGType) SelectAllWithLimit(mrQueryable mingru.Queryable, page int, pageSize int) ([]EmployeeAGSelectAllWithLimitResult, bool, error) {
	if page <= 0 {
		err := fmt.Errorf("Invalid page %v", page)
		return nil, false, err
	}
	if pageSize <= 0 {
		err := fmt.Errorf("Invalid page size %v", pageSize)
		return nil, false, err
	}
	limit := pageSize + 1
	offset := (page - 1) * pageSize
	max := pageSize
	rows, err := mrQueryable.Query("SELECT `emp_no`, `first_name`, `last_name`, `gender`, `birth_date`, `hire_date` FROM `employees` ORDER BY `hire_date` LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, false, err
	}
	result := make([]EmployeeAGSelectAllWithLimitResult, 0, limit)
	itemCounter := 0
	defer rows.Close()
	for rows.Next() {
		itemCounter++
		if itemCounter <= max {
			var item EmployeeAGSelectAllWithLimitResult
			err = rows.Scan(&item.ID, &item.FirstName, &item.LastName, &item.Gender, &item.BirthDate, &item.HireDate)
			if err != nil {
				return nil, false, err
			}
			result = append(result, item)
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, false, err
	}
	return result, itemCounter > len(result), nil
}

type EmployeeAGSelectByIDResult struct {
	BirthDate time.Time
	FirstName string
	Gender    string
	HireDate  time.Time
	ID        int
	LastName  string
}

func (mrTable *EmployeeAGType) SelectByID(mrQueryable mingru.Queryable, id int) (EmployeeAGSelectByIDResult, error) {
	var result EmployeeAGSelectByIDResult
	err := mrQueryable.QueryRow("SELECT `emp_no`, `first_name`, `last_name`, `gender`, `birth_date`, `hire_date` FROM `employees` WHERE `emp_no` = ?", id).Scan(&result.ID, &result.FirstName, &result.LastName, &result.Gender, &result.BirthDate, &result.HireDate)
	if err != nil {
		return result, err
	}
	return result, nil
}

type EmployeeAGSelectPagedResult struct {
	BirthDate time.Time
	FirstName string
	Gender    string
	HireDate  time.Time
	ID        int
	LastName  string
}

func (mrTable *EmployeeAGType) SelectPaged(mrQueryable mingru.Queryable, page int, pageSize int) ([]EmployeeAGSelectPagedResult, bool, error) {
	if page <= 0 {
		err := fmt.Errorf("Invalid page %v", page)
		return nil, false, err
	}
	if pageSize <= 0 {
		err := fmt.Errorf("Invalid page size %v", pageSize)
		return nil, false, err
	}
	limit := pageSize + 1
	offset := (page - 1) * pageSize
	max := pageSize
	rows, err := mrQueryable.Query("SELECT `emp_no`, `first_name`, `last_name`, `gender`, `birth_date`, `hire_date` FROM `employees` ORDER BY `hire_date` LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, false, err
	}
	result := make([]EmployeeAGSelectPagedResult, 0, limit)
	itemCounter := 0
	defer rows.Close()
	for rows.Next() {
		itemCounter++
		if itemCounter <= max {
			var item EmployeeAGSelectPagedResult
			err = rows.Scan(&item.ID, &item.FirstName, &item.LastName, &item.Gender, &item.BirthDate, &item.HireDate)
			if err != nil {
				return nil, false, err
			}
			result = append(result, item)
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, false, err
	}
	return result, itemCounter > len(result), nil
}

func (mrTable *EmployeeAGType) SelectSig(mrQueryable mingru.Queryable, id int) (time.Time, error) {
	var result time.Time
	err := mrQueryable.QueryRow("SELECT `birth_date` FROM `employees` WHERE `emp_no` = ?", id).Scan(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (mrTable *EmployeeAGType) UpdateName(mrQueryable mingru.Queryable, id int, firstName string, lastName string) error {
	result, err := mrQueryable.Exec("UPDATE `employees` SET `first_name` = ?, `last_name` = ? WHERE `emp_no` = ?", firstName, lastName, id)
	return mingru.CheckOneRowAffectedWithError(result, err)
}