---
title: feb09
date: 2026-02-09T09:17:47.883826-05:00
---

# Important Dates

- Assignment 1B due Friday

# Message Passing

```racket
(define fnil
    (lambda (message)
        (cond ((eq? message 'null?) #t)
              ((eq? message 'pair?) #f)
              (else (error message "not an operation on fnil")))))

(define (fnull? obj) (and (procedure? obj) (obj 'fnull)))

(define (fpair? obj) (and (procedure? obj) (obj 'pair?)))

(fnull? fnil) ; #t
(fpair? fnil) ; #f
```

# Lexical Analysis

Example:

```
let x = 99, y = 1 in x + y
```

- Words: `in`, `x`, `99`
- Phrase: `x + y`
  - Made up of words
