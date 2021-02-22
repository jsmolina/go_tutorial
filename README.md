# go_tutorial

## an array of 10 pointers to string
If C has:
`char *str[20];`

Go has

`str [10]*string` 

## Vars auto initialize
unlike C, booleans will default false, floats to zero, strings to empty string.

`var ageofUniverse int` defaults to zero.

## short declaration
You could assign a value that type is inferred
ageofUniverse := 14e9

## multiple declares
```
var (
  ageofUniverse int
  lang = "go"
)
```

## declare and asign
var ageofuniverse uint64

ageofuniverse = 1382000000


## swaps (like python)
foo, bar := 1,2
foo, bar = bar, foo
