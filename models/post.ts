import * as dd from 'dd-models';
import user from './user';

class Post extends dd.Table {
  id = dd.pk();
  // Foreign key to user.id
  user_id = user.id;
  title = dd.varChar(500);
  content = dd.text();
  cmtCount = dd.setName('cmt_c', dd.unsignedInt(0));
}

export default dd.table(Post);
