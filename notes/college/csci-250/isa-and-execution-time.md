---
title: ISA and Execution Time
date: 2026-01-16T13:38:34.8472-05:00
---

# ISA Spectrum

<!-- DRAWING -->

- **Complex Instruction Set**: Instruction set where there exists instructions that do a large number of things for the programmer.
- **Reduced Instruction Set**: Instruction set where all instructions do "simple" things.

## Position on Spectrum

- Position on spectrum doesn't correlate with number of instructions.
- Only thing that matters is how complex the instructions are.
- Example: x64 instruction set is turing complete with just the move instruction. If the ISA was just that instruction it would still be a CISC.

# Execution Time

<!-- DRAWING -->

- **IC**: instruction count
- **CPI**: (cycles per instruction) - weighted average of CPI for all instructions in ISA (what we will be doing for this class)
- **CCT**: clock cycle time
- This formula will be used to compare ISAs.

# Deeper Dive into CISC v.s. RISC

CCT (clock cycle time) is only ever affected by the size of the processor.

## CISC (Complex Instruction Set)

- Smaller IC
- Higher CPI
- CCT _not affected_

## RISC (Reduced Instruction Set)

- Higher IC
- Smaller CPI
- CCT _not affected_

# Logical Look at a Computer

<!-- DRAWING -->
