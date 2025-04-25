---
title: 'Flip Flops'
---

# What is a Flip Flop IC (Integrated Circuit)?

A flip flop IC is a latch/switch that uses memory so that the current state and previous states of the circuit can be used to dictate which switches are activated.

## Basic Sequential Logic Circuit: SR Flip Flop

A SR Flip Flop (A.K.A SR Latch) has two inputs:

1. Set line
2. Reset line

It then only has one output that is either at 1 or 0.

## Triggers

Triggering is how a flip-flop knows how to change values. There are two main types of triggering:

1. Level Triggering: Can only change value when the clock is either at exactly 1 or 0. It cannot be in-between those two values.
   1. Value changes when the value changes from either 1 -> 0 or 0 -> 1.
2. Edge Triggering: Can only change values when hte clock is NOT at exactly 1 or 0. Basically the opposite of level triggering.
   1. Value changes when the state of the signal changes.

## Latch

- A latch is a type of basic sequential circuit.
- Is level-triggered.
- Responsive whenever the control signal is active.

## Latch v.s. Flip-Flop

- Latch:
  - Level-triggered (changes when the value changes from 1 or from 0).
  - Either is a 1 or a 0.
  - Asynchronous (instantaneous)
  - Used for simple data storage.
  - Examples:
- Flip-Flop:
  - Edge-triggered (changes when the state of the value changes).
  - Synchronous (clock-controlled).
  - Precise data storage and timing.
  - Examples: D Flip-Flip and JK Flip-Flop
