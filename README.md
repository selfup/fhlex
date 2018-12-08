# Fhlex

Flex radio hex CAT commands

## Interestings

- `ZZ` starts all commands
- `;` terminates all commands

#### Character to HEX

```
Z == 7A
M == 6D
D == 64
; == 3B
```

Please refer to:

1. docs/ascii_conversion_table.html
1. https://selfup.github.io/fhlex/ascii_conversion_table.html
1. https://www.eso.org/~ndelmott/ascii.html

## Flex CAT

A basic command to read USB mode:

`7A7A6D643B` aka `ZZMD;`

Should receive: `ZZMD01` if USB is on.

Here is the **PowerSDR** Flex CAT command reference guide:

https://www.flexradio.com/downloads/powersdr-cat-command-reference-guide/
