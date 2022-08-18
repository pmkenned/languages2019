." hello, world!" CR

0 Value fd-in
0 Value fd-out
: open-input ( addr u -- ) r/o open-file throw to fd-in ;
: open-output ( addr u -- ) w/o create-file throw to fd-out ;
: close-input ( -- ) r/o fd-in close-file throw ;
: close-output ( -- ) w/o fd-out close-file throw ;

256 Constant max-line
Create line-buffer max-line 2 + allot
: scan-file ( addr u -- )
  line-buffer max-line fd-in read-line drop drop drop
  begin
    line-buffer max-line fd-in read-line throw
  while
      >r 2dup line-buffer r> compare 0=
    until
  else
    drop
  then
  2drop ;

: copy-file ( -- )
  begin
    line-buffer max-line fd-in read-line throw
  while
    line-buffer swap fd-out write-line throw
  repeat ;

s" in.txt" open-input
s" out.txt" open-output

s" hello you" scan-file

copy-file

close-input
close-output
