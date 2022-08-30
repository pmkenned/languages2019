#!/usr/bin/env lua

-- keywords:
--  and or not
--  true false nil
--  if then else elseif end
--  for repeat while until do break
--  in
--  function return
--  local
-- types:
--  string
--  number
--  function
--  boolean
--  nil
--  userdata
--  thread
--  table

function fact (n)
    if n == 0 then
        return 1
    else
        return n * fact(n-1)
    end
end

function vec2_length (x, y)
    local hyp_squared = x^2 + y^2
    return math.sqrt(hyp_squared)
end

-- print("enter a number:")
-- a = io.read("*number")
a = 5
print(fact(a))
a = 5e2
print(vec2_length(3, 4))
b = {1,2,3}
print("hello\n", a, "hi", arg[0], b)
print(type("hello"))
print(type(10))
print(type(true))
print(type(nil))
print(type(type))
if (0 and "") then
    print("0 and empty string are both true")
end
if ((not nil) and (not false)) then
    print("nil and false are both false")
end

s = "hello\10"
s2 = string.gsub(s, "h", "H") -- strings are immutable
print(s, s2)

string_literal = [=[
this is a string literal
it contains nested [[ braces ]]
and code
function foo(a, b, c)
    return 0
end
]=]

print(string_literal)
print("10"+1)
name = "Paul"
print("hi, " .. name)

if (not (10 == "10")) then
    print('10 != "10"')
end

if (10 == tonumber("10")) then
    print("10 == 10")
end

if (tostring(10) == "10") then
    print('"10" == "10"')
end
