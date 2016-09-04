require_relative './node'

n = Test::Node::Node.new
n.next = Test::Node::Node.new
n.next.next = Test::Node::Node.new
n.next.next.value = "hello"
p n.next.next.value