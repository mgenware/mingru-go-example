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

type EmploymentHistoryTAAGType struct {
}

var EmploymentHistoryTA = &EmploymentHistoryTAAGType{}

// ------------ Actions ------------

type EmploymentHistoryTAAGSelectAllHistoryResult struct {
	DepartmentName    string
	EmployeeFirstName string
	EmployeeLastName  string
	FromDate          time.Time
	ToDate            time.Time
}

func (mrTable *EmploymentHistoryTAAGType) SelectAllHistory(mrQueryable mingru.Queryable, page int, pageSize int) ([]EmploymentHistoryTAAGSelectAllHistoryResult, bool, error) {
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
	rows, err := mrQueryable.Query("SELECT `dept_emp`.`from_date`, `dept_emp`.`to_date`, `join_1`.`first_name`, `join_1`.`last_name`, `join_2`.`dept_name` FROM `dept_emp` AS `dept_emp` INNER JOIN `employees` AS `join_1` ON `join_1`.`emp_no` = `dept_emp`.`emp_no` INNER JOIN `departments` AS `join_2` ON `join_2`.`dept_no` = `dept_emp`.`dept_no` ORDER BY `dept_emp`.`emp_no`, `dept_emp`.`dept_no` LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, false, err
	}
	result := make([]EmploymentHistoryTAAGSelectAllHistoryResult, 0, limit)
	itemCounter := 0
	defer rows.Close()
	for rows.Next() {
		itemCounter++
		if itemCounter <= max {
			var item EmploymentHistoryTAAGSelectAllHistoryResult
			err = rows.Scan(&item.FromDate, &item.ToDate, &item.EmployeeFirstName, &item.EmployeeLastName, &item.DepartmentName)
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