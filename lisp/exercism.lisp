#!/usr/bin/env -S sbcl --script

(defpackage :my-package
    (:use :cl))

(defun categorize-number (l n)
    (cond
        ((evenp n) (cons (car l) (cons n (cdr l))))
        (t         (cons (cons n (car l)) (cdr l)))
    )
)

(defun partition-numbers (l)
    (reduce #'categorize-number l :initial-value '(nil))
)

(format t "~a~%" (partition-numbers '(1 2 3 4)))
(format t "~a~%" (equal '(nil) (partition-numbers 'nil)))
(format t "~a~%" (equal '((1)) (partition-numbers '(1))))
(format t "~a~%" (equal '(nil 2) (partition-numbers '(2))))
(format t "~a~%" (equal '((1)) (partition-numbers '(1))))
(format t "~a~%" (equal '(nil 2) (partition-numbers '(2))))
(format t "~a~%" (equal '((5 3 1) 6 4 2) (partition-numbers '(1 2 3 4 5 6))))

(format t "~a~%" (partition-numbers '(42 31 24 13)))
(format t "~a~%" '((13 31) . 42))

(defun foo (s1)
    (length s1)
)

(format t "~a~%" (foo "hello"))
