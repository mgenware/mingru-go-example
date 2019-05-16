import * as mr from "mingru";
// Import model actions
import employeeTA from "./models/employeeTA";

(async () => {
  const actions = [employeeTA];
  const dialect = new mr.MySQL();
  // Build Go code to '../da/` directory
  await mr.build(actions, dialect, "../da/", { cleanBuild: true });
})();
