require 'net/http'
require 'uri'
require 'open3'

Open3.popen3("ruby #{File.dirname(__FILE__)}/start_httpserv.rb") do |_, _, _, thr|
  #puts "PID=#{thr.pid}"
  sleep 0.2
  puts Net::HTTP.get(URI('http://localhost:8080/gorb'))
  system "kill -9 #{thr.pid}"
end

