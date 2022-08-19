: fib ( n -- fib_n )
  dup
  0 >
  if
    >r
    0 1
    r>
    0 do dup rot + loop
    drop
  else
    drop
    0
  then
;

: fib_to_n 0 do i fib . loop ;

20 fib_to_n
.s
