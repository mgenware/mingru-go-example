import * as dd from "dd-models";
import user from "./user";

const userTA = dd.actions(user);
// Select a user profile by ID
userTA.select("UserProfile", user.id, user.name).byID();
// Select all user profiles
userTA.selectAll("AllUserProfiles", user.id, user.name);
// Select a single field by ID
userTA.selectField("Sig", user.sig).byID();

// Update a row
userTA
  .updateOne("UserProfile")
  .setInputs(user.name, user.sig)
  .byID();

// Delete a row by ID
userTA.deleteOne("ByID").byID();

export default userTA;
