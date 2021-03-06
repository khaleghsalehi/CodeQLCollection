class AstNode extends @ruby_ast_node {
  string toString() { none() }
}

class Location extends @location {
  string toString() { none() }
}

from AstNode id, AstNode body, Location loc
where ruby_lambda_def(id, body) and ruby_ast_node_info(id, _, _, loc)
select id, body, loc
