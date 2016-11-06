require_relative './blocks'

include Test::Blocks

# Basic block callback
total = do_with(5) {|v| v * 4 }
puts "5 * 4 = #{total}"
