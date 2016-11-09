require_relative './array'

include Test::Array

puts "# Reverse an array:"
p reverse_array(["one", "two", "three"])

puts "\n# Mutable string array:"
mutate_array(["one", "two", "three"]) do |arr|
  puts "Ruby has #{arr.inspect}"
  arr[1] = "zero"
  arr.push "nine"
end

puts "\n# Mutable int array:"
mutate_int_array((1..10).to_a) do |arr|
  puts "Ruby has #{arr.inspect}"
  arr.each.with_index do |el, i|
    arr[i] = el * 5
  end
end
