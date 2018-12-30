 /******************************************************************************************
 * The code was automatically generated by mingru (https://github.com/mgenware/mingru)
 * Do not edit this file manually, your changes will be overwritten.
 ******************************************************************************************/

package da

import (
	"github.com/mgenware/go-packagex/database/sqlx"
)

// TableTypeUser ...
type TableTypeUser struct {
}

// User ...
var User = &TableTypeUser{}

// ------------ Actions ------------

// SelectUserProfileResult ...
type SelectUserProfileResult struct {
	UserID   uint64
	UserName string
	UserSig  *string
}

// SelectUserProfile ...
func (da *TableTypeUser) SelectUserProfile(queryable sqlx.Queryable, userID uint64) (*SelectUserProfileResult, error) {
	result := &SelectUserProfileResult{}
	err := queryable.QueryRow("SELECT `id`, `name`, `sig` FROM `user` WHERE `id` = ?", userID).Scan(&result.UserID, &result.UserName, &result.UserSig)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SelectAllUserProfilesResult ...
type SelectAllUserProfilesResult struct {
	UserID   uint64
	UserName string
	UserSig  *string
}

// SelectAllUserProfiles ...
func (da *TableTypeUser) SelectAllUserProfiles(queryable sqlx.Queryable) ([]*SelectAllUserProfilesResult, error) {
	rows, err := queryable.Query("SELECT `id`, `name`, `sig` FROM `user`")
	if err != nil {
		return nil, err
	}
	result := make([]*SelectAllUserProfilesResult, 0)
	defer rows.Close()
	for rows.Next() {
		item := &SelectAllUserProfilesResult{}
		err = rows.Scan(&item.UserID, &item.UserName, &item.UserSig)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SelectSig ...
func (da *TableTypeUser) SelectSig(queryable sqlx.Queryable, userID uint64) (*string, error) {
	var result *string
	err := queryable.QueryRow("SELECT `sig` FROM `user` WHERE `id` = ?", userID).Scan(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateUserProfile ...
func (da *TableTypeUser) UpdateUserProfile(queryable sqlx.Queryable, userID uint64, userName string, userSig *string) error {
	result, err := queryable.Exec("UPDATE `user` SET `name` = ?, `sig` = ? WHERE `id` = ?", userName, userSig, userID)
	return sqlx.CheckOneRowAffectedWithError(result, err)
}

// DeleteByID ...
func (da *TableTypeUser) DeleteByID(queryable sqlx.Queryable, userID uint64) error {
	result, err := queryable.Exec("DELETE FROM `user` WHERE `id` = ?", userID)
	return sqlx.CheckOneRowAffectedWithError(result, err)
}
