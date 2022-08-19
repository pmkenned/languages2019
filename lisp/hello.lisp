#!/usr/bin/env -S sbcl --script

(defun main ()
    (format t "Hello, world!~%"))
;    (write-line "Hello, World!")) ; alternative

(sb-ext:save-lisp-and-die "hello"
:executable t
:toplevel 'main)
