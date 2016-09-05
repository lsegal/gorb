require_relative './fib'

include Test::Fib

class RubyFib
  def fib(n)
    n < 2 ? n : fib(n-1) + fib(n-2)
  end
end

p is_prime?(5)
p Fibonacci.new.red.r

require 'benchmark'
TIMES = 10
rf = RubyFib.new
gf = Fibonacci.new
Benchmark.bmbm do |x|
  x.report("golang") { TIMES.times { gf.fib(20) } }
  x.report("ruby") { TIMES.times { rf.fib(20) } }
end
