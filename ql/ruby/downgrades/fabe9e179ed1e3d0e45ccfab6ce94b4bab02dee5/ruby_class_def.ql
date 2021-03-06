class AstNode extends @ruby_ast_node {
  string toString() { none() }
}

class Location extends @location {
  string toString() { none() }
}

from AstNode id, AstNode name, Location loc
where ruby_class_def(id, name) and ruby_ast_node_info(id, _, _, loc)
select id, name, loc
