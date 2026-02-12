---
title: feb09
date: 2026-02-09T15:04:32.304359-05:00
---

# if

In this example we are assuming that `s0` and `s1` are loaded correctly with the values we want.

```
if (s0 == s1) {
    // true
} else {
    // false
}
```

Here is what this code would look like in assembly:

```mipsasm
bne $s0, $s1, else
# true
j done

else:
    # false

done:
```

# Relative Comparisons

| Comparison | Difference Version |
| ---------- | ------------------ |
| `A < B`    | `A-B < 0`          |
| `A == B`   | `A-B == 0`         |
| `A > B`    | `A-B > 0`          |

# Set if less than

```mipsasm
slt rd,rs,rt # if rs < then rd = 1  else rd = 0
```

Example:

```mipsasm
slt $t0, $s0, $s1
bne $t0, $zero, foo
```

This would "translate" to this code in a high-level-language:

```
if (s0 < s1) goto foo
```

# Pseudo-Instructions we can't use

- `sgt`
- `sle`
- `sge`
- `blt`
- `ble`
- `bgt`
- `bge`

# For loops

```
for (s0 = 0; s0 < 10; s0+1) {
    // body
}
```

Here is a for loop in mips:

```mipsasm
li $s0, 0

loop:
    slti $t0, $s0, 10
    beq $t0, $zero, done
    # body
    addi $s0, $s0, 1
    j loop

done:
```

# Do while

```
do {
    # body
    s0++
} while (s0 < 10);
```

```mipsasm
li $s0, 0

loop:
    # body
    addi $s0, $s0, 1
    slti $t0, $s0, 10
    bne  $t0, $zero, loop
```

Cleaner but isn't the same as a normal while loop. For example for the loop up above if `$s0 > 10` then it would go on forever cause it's only checking when it reaches 10.

# While

```mipsasm
loop:
    slti $t0, $s0, 10
    beq  $t0, $zero, done

    # body

    addi $s0, $s0, 1
    bne $t0, $zero, loop
    j loop

done:
```

# Switch Statements

```
switch (s0) {
    case 0:
        // code for 0
        break
    case 1:
        // code for 1
        break
    case 2:
        // code for 2
        break
    default:
        // default code

}
```
