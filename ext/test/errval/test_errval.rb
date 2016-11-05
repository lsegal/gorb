require_relative './errval'

puts Test::Errval.flip(10)
puts Test::Errval.flip(4)

begin
  puts Test::Errval.flip(0)
rescue => err
  puts "Got error '#{err}'"
end
