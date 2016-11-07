require_relative './array'

puts "\n# Reverse an array:"
p Test::Array.reverse_array(["one", "two", "three"])

puts "\n# Mutable string array:"
Test::Array.mutate_array(["one", "two", "three"]) do |arr|
  puts "Ruby has #{arr.inspect}"
  arr[1] = "zero"
  arr.push "nine"
end

puts "\n# Mutable int array:"
Test::Array.mutate_int_array((1..10).to_a) do |arr|
  puts "Ruby has #{arr.inspect}"
  arr.each.with_index do |el, i|
    arr[i] = el * 5
  end
end
