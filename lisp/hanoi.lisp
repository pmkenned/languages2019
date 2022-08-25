#!/usr/bin/env -S sbcl --script

(defun hanoi (n from to spare)
    (when (> n 0)
        (hanoi (- n 1) from spare to)
        (format t "~a ~a~%" from to)
        (hanoi (- n 1) spare to from)
    )
)

(hanoi 4 0 1 2)
