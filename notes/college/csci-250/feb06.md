---
title: Change of Execution
date: 2026-02-06T15:02:39.931771-05:00
---

# Important Dates

- Exam 1 next Wednesday in GOL-2400.
- Review exam session 10am on zoom.
- Experiment 2 due next Friday.

## Different Types

- Conditional means **branch** instruction
- Unconditional means **jump** instruction

- `goto` (jump to fixed)
- Function calls (jump to fixed)
- Function returns (jump to variable)
- If statements (branch to fixed)
- Switches (branch to variable)
- Loops (branch to fixed)
- Function calls through function pointers (jump to variable)

## How do they work?

1. Fetch the instruction from memory
2. Decode
3. Execute
4. Repeat back to step 1

## Program counter (PC)

1. Current instruction to execute
2. Next instruction to execute

### Fetching Instruction

This the way that MIPS fetches the instruction. It is possible for some ISAs to do things in the opposite order.

1. Read instruction
2. Increment PC

# goto in assembly

```mipsasm
jump inst
j    label  # PC = label
```

# J-Type Instruction

Jump type instruction.

<!-- DRAWING -->

At runtime we have the following computation for the PC:

<!-- DRAWING -->

## Function Call

Jump and link (saving the PC as the return address `$ra`):

```mipsasm
jal label # $ra = PC
          # PC = label
```

We can then do this to go back from the function:

```mipsasm
jr $ra
```

## If Statement

```mipsasm
beq rs,rt,label # if rs==rt then PC=label
```

If the value of register `$rs` is equal to `$rt` then jump to `label`.

We have `bne` too, which is the same as `beq` but it checks to see if the two registers are **not equal** instead:

```mipsasm
bne rs,rt,label # if rs!=rt then PC=label
```
