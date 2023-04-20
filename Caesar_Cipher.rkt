#lang scheme
(define (caesarPhire str shift)
  (cond
    ((= (string-length str) 0) '())
    ((not (char-alphabetic? (car (string->list str))))'())
    (else (cons (integer->char (+ (remainder (+ (char->integer (char-upcase(car (string->list str)))) shift -65) 26) 65)) (caesarPhire (list->string (cdr (string->list str))) shift))
  )))
(define (caesarPhireList list shift)
    (cond
    ((null? list) '())
    (else (cons (list->string (caesarPhire (car list) shift)) (caesarPhireList (cdr list) shift)))))
(caesarPhireList '("I" "LOVE" "CS!") 5)

