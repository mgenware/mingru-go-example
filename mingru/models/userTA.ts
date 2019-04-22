import * as dd from "dd-models";
import user from "./user";

export class UserTA extends dd.TA {
  // Select a user profile by ID
  selectUserProfile = dd.select(user.id, user.name, user.sig).byID();
  // Select all user profiles
  selectAllUserProfiles = dd.selectAll(user.id, user.name, user.sig);
  // Select the single sig field by ID
  selectSig = dd.selectField(user.sig).byID();

  // Update an user profile by ID
  updateUserProfile = dd
    .updateOne()
    .setInputs(user.name, user.sig)
    .byID();

  // Update all user.sig to an empty string
  updateAllSigToEmpty = dd.unsafeUpdateAll().set(user.sig, dd.sql`''`);

  // Delete an user by ID
  deleteByID = dd.deleteOne().byID();

  // Delete all users by a specified name
  deleteByName = dd.deleteSome().where(user.name.isEqualToInput());

  // Delete all users
  deleteAll = dd.unsafeDeleteAll();

  // Insert a new user
  insertUser = dd
    .insertOne()
    .set(user.sig, dd.sql`''`)
    .setInputs(user.name, user.age);
}

export default dd.ta(user, UserTA);
