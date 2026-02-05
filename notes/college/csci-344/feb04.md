---
title: 'Class of February 4th'
date: 2026-02-04
---

# Evaluation Context

foo bar

```racket
#lang racket

; (+ 2 []) <- (* 3 4)

(lambda (x) (+ 2 x))
```

## New Operations

```racket
(define (a m n k)) (k (+ m n))

(mul 3 4 (lambda (r1)
(add 2 r1 (lambda (r2)
r2))))

(define (countk xs k)
    (cond ((null? xs) (k0))
          ((pair? xs) (countk (cdr xs) (lambda (n) (k (+ 1 n)))))
          (else "Bad List")))
```
