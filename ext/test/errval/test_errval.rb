require_relative './errval'

[10, 4, 0].each do |val|
  begin
    puts Test::Errval.flip(val)
  rescue => err
    puts "Got error '#{err}'"
  end
end
