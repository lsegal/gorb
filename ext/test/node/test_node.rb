require_relative './node'

include Test::Node

root = Node.new('a', Node.new('b', Node.new('c', nil)))

n = root
loop do
  print "#{n.value} -> "
  break if n.end?
  n = n.next
end

puts 'END'
