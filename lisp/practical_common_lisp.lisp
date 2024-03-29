#!/usr/bin/env -S sbcl --script

; left off at "Improving the User Interaction"
; https://gigamonkeys.com/book/practical-a-simple-database.html

(defun make-cd (title artist rating ripped)
  (list :title title :artist artist :rating rating :ripped ripped))

(defvar *db* nil)
(defun add-record (cd) (push cd *db*))

;(defun dump-db ()
;  (dolist (cd *db*)
;    (format t "~{~a:~10t~a~%~}~%" cd)))
(defun dump-db ()
  (format t "~{~{~a:~10t~a~%~}~%~}" *db*))

(add-record (make-cd "Roses" "Kathy Mattea" 7 t))
(add-record (make-cd "Fly" "Dixie Chicks" 8 t))
(add-record (make-cd "Home" "Dixie Chicks" 9 t))

(dump-db)
