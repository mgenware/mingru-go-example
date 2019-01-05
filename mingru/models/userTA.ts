import * as dd from 'dd-models';
import user from './user';

const userTA = dd.actions(user);
// Select a user profile by ID
userTA.select('UserProfile', user.id, user.name, user.sig).byID();
// Select all user profiles
userTA.selectAll('AllUserProfiles', user.id, user.name, user.sig);
// Select the single sig field by ID
userTA.selectField('Sig', user.sig).byID();

// Update an user profile by ID
userTA
  .updateOne('UserProfile')
  .setInputs(user.name, user.sig)
  .byID();

// Update all user.sig to an empty string
userTA.updateAll('AllSigToEmpty').set(user.sig, dd.sql`''`);

// Delete an user by ID
userTA.deleteOne('ByID').byID();

export default userTA;