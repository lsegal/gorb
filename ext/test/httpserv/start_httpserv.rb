require_relative './httpserv'

Test::Httpserv.serve(":8080") do |url|
  puts "Got request for #{url}"
  "Hello <code>#{url}</code>!"
end
