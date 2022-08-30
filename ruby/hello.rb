#!/usr/bin/env ruby

# arrays, hashes
# booleans, operators
# string templates
# symbols
# method chaining
# ! and ? methods
# string comparison, concat, substr
# to_s, to_i
# default parameters
# attr_accessor
# def initialize
# inheritance? interfaces?
# variadic methods?

# keywords:
#  __ENCODING__ __LINE__ __FILE__
#  BEGIN END
#  
#  alias
#  module class
#  super self
#  def undef
#  defined?
#  
#  true false nil
#  not and or
#  
#  begin end
#  unless if elsif else then until for in while do case when break next redo return
#  yield
#  
#  ensure rescue retry

class Foo
    attr_accessor :age
    def initialize()
        @age = 3
    end
end

def fact(n)
    if n == 1
        return 1
    end
    return n*fact(n-1)
end

puts "#{__LINE__}: #{"hello there"[0..5]}"

# and has lower precedence than &&
if '' and 0     then puts "empty string and 0 are true" end
if not nil      then puts "nil is false" end
if ''.empty?    then puts "empty string is empty" end
if true         then puts "true is true" end

puts fact(10)

my_hash = {}
my_hash["a"] = 7
my_hash["b"] = "hi"

my_hash.each { |k, v|
    puts "#{k.ljust(30)}: #{v}"
}

my_array = []
my_array << 1 << 2 << 3
my_array.concat [4, 5, 6]
my_array.push *[7, 8]
my_array = my_array + [9, 10]
my_array += [11, 12]
(my_array << [13, 14]).flatten!
my_array[my_array.length, 0] = 15
my_array.insert(my_array.length, 16)
my_array.reverse!
puts "my_array: #{my_array}"
my_array.sort!
puts "my_array: #{my_array}"

f = Foo.new
puts f.age

puts ":hello.class:             #{:hello.class}"
puts "method(:fact).class:      #{method(:fact).class}"
puts "method(:method).class:    #{method(:method).class}"
puts "Method.class:             #{Method.class}"
puts "Symbol.class:             #{Symbol.class}"
puts "Foo.class:                #{Foo.class}"
puts "Class.class:              #{Class.class}"
puts "\"hello\".class:            #{"hello".class}"
puts "123.class:                #{123.class}"
puts "nil.class:                #{nil.class}"
puts "false.class:              #{false.class}"
puts "true.class:               #{true.class}"
