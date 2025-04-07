# Multiplexer and DeMultiplexers

## What is a Multiplexer (MUX)?

A multiplexer is a type of "switch" that takes in multiple inputs (sources) and a single output (destination). This allows us to have only **one** signal to pass through at a time. Here is a diagram of that basic process:

```txt
         S1   S0
          |   |
          v   v
         ________
I0 ---> |        \
I1 ---> |  MUX    )----> Y
I2 ---> |        /
I3 ---> |_______/

```

_I# are inputs, S# are select lines, and Y is the output_

## Select Lines

Select lines allow us to choose which input we want to have selected at a given time.

### Example of Selection

Based on the diagram from up above here is which inputs are selected based on the given select line values:

| S0  | S1  | Selected Input |
| --- | --- | -------------- |
| 0   | 0   | I0             |
| 0   | 1   | I1             |
| 1   | 0   | I2             |
| 1   | 1   | I3             |

## What is a DeMultiplexers (DEMUX)?

A DeMultiplexer is basically a multiplexer but in reverse. Here is a diagram of what that looks like:

```txt
         S1   S0
          |   |
          v   v
         __________
X -----> | DEMUX   \
         |         )----> Y0
         |         )----> Y1
         |         )----> Y2
         |_________/----> Y3
```

_X is the input, S# are the select lines, and Y# are the outputs_
