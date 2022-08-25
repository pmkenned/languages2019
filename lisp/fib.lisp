#!/usr/bin/env -S sbcl --script

; O(fib), terrible
(defun fib (n)
    (if (< n 2)
        n
        (+ (fib (- n 1)) (fib (- n 2)))
    )
)

(format t "~a~%" (fib (parse-integer (nth 1 sb-ext:*posix-argv*))))
