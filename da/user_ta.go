 /******************************************************************************************
 * This code was automatically generated by mingru (https://github.com/mgenware/mingru)
 * Do not edit this file manually, your changes will be overwritten.
 ******************************************************************************************/

package da

import (
	"github.com/mgenware/go-packagex/v5/dbx"
)

// TableTypeUser ...
type TableTypeUser struct {
}

// User ...
var User = &TableTypeUser{}

// ------------ Actions ------------

// UserTableSelectUserProfileResult ...
type UserTableSelectUserProfileResult struct {
	ID   uint64
	Name string
	Sig  *string
}

// SelectUserProfile ...
func (da *TableTypeUser) SelectUserProfile(queryable dbx.Queryable, id uint64) (*UserTableSelectUserProfileResult, error) {
	result := &UserTableSelectUserProfileResult{}
	err := queryable.QueryRow("SELECT `id`, `name`, `sig` FROM `user` WHERE `id` = ?", id).Scan(&result.ID, &result.Name, &result.Sig)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UserTableSelectAllUserProfilesResult ...
type UserTableSelectAllUserProfilesResult struct {
	ID   uint64
	Name string
	Sig  *string
}

// SelectAllUserProfiles ...
func (da *TableTypeUser) SelectAllUserProfiles(queryable dbx.Queryable) ([]*UserTableSelectAllUserProfilesResult, error) {
	rows, err := queryable.Query("SELECT `id`, `name`, `sig` FROM `user`")
	if err != nil {
		return nil, err
	}
	result := make([]*UserTableSelectAllUserProfilesResult, 0)
	defer rows.Close()
	for rows.Next() {
		item := &UserTableSelectAllUserProfilesResult{}
		err = rows.Scan(&item.ID, &item.Name, &item.Sig)
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
func (da *TableTypeUser) SelectSig(queryable dbx.Queryable, id uint64) (*string, error) {
	var result *string
	err := queryable.QueryRow("SELECT `sig` FROM `user` WHERE `id` = ?", id).Scan(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// UpdateUserProfile ...
func (da *TableTypeUser) UpdateUserProfile(queryable dbx.Queryable, id uint64, name string, sig *string) error {
	result, err := queryable.Exec("UPDATE `user` SET `name` = ?, `sig` = ? WHERE `id` = ?", name, sig, id)
	return dbx.CheckOneRowAffectedWithError(result, err)
}

// UpdateAllSigToEmpty ...
func (da *TableTypeUser) UpdateAllSigToEmpty(queryable dbx.Queryable) (int, error) {
	result, err := queryable.Exec("UPDATE `user` SET `sig` = ''")
	return dbx.GetRowsAffectedIntWithError(result, err)
}

// DeleteByID ...
func (da *TableTypeUser) DeleteByID(queryable dbx.Queryable, id uint64) error {
	result, err := queryable.Exec("DELETE FROM `user` WHERE `id` = ?", id)
	return dbx.CheckOneRowAffectedWithError(result, err)
}

// DeleteByName ...
func (da *TableTypeUser) DeleteByName(queryable dbx.Queryable, name string) (int, error) {
	result, err := queryable.Exec("DELETE FROM `user` WHERE `name` = ?", name)
	return dbx.GetRowsAffectedIntWithError(result, err)
}

// DeleteAll ...
func (da *TableTypeUser) DeleteAll(queryable dbx.Queryable) (int, error) {
	result, err := queryable.Exec("DELETE FROM `user`")
	return dbx.GetRowsAffectedIntWithError(result, err)
}

// InsertUser ...
func (da *TableTypeUser) InsertUser(queryable dbx.Queryable, name string, age int) (uint64, error) {
	result, err := queryable.Exec("INSERT INTO `user` (`sig`, `name`, `age`) VALUES ('', ?, ?)", name, age)
	return dbx.GetLastInsertIDUint64WithError(result, err)
}
