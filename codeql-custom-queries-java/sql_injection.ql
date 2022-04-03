import java
import semmle.code.java.dataflow.DataFlow::DataFlow

class StringConcat extends AddExpr {
    StringConcat() { getType() instanceof TypeString }
  }

  
from MethodAccess ma, StringConcat stringConcat
where
  ma.getMethod().getName().regexpMatch("createQuery") and
  localFlow(exprNode(stringConcat), exprNode(ma.getArgument(0)))
select ma, "SQL query vulnerable to injection."