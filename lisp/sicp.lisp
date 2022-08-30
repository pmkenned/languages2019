#!/usr/bin/env -S sbcl --script

(defun apply_func (f n)
    (funcall f n))

; passing global function to apply_func
(defun my_global_func (n)
    (1+ n))
(format t "~a~%" (apply_func #'my_global_func 1))

; passing local function to apply_func
(flet ((my_local_func (n) (1+ n)))
    (format t "~a~%" (apply_func #'my_local_func 1)))

; passing in a lambda function
(format t "~a~%" (apply_func (lambda (n) (1+ n)) 1))

(defun sum (term a next b)
    (if (> a b)
        0
        (+ (funcall term a)
           (sum term (funcall next a) next b))))

(flet ((ident (a) a))
    (defun sum-int (a b)
        (sum #'ident a #'1+ b)))

(flet ((square (a) (* a a)))
    (defun sum-sq (a b)
        (sum #'square a #'1+ b)))

(defun pi-sum (a b)
    (sum (lambda (i) (/ 1 (* i (+ i 2))))
         a
         (lambda (i) (+ i 4))
         b))

(format t "~a~%" (sum-int 1 4))
(format t "~a~%" (* 8 (pi-sum 1.0 50000.0)))
