# Battle

There can be only one :)

A battle is setup, random fighters are generated, they duke it out and
eventually there is a winner.


## Docs? Tests?!

There are no docs, and almost no comments. There are no tests. This is a toy
project.


## Example

```
$ COUNT=3 go run cmd/battle/main.go
Touchbead (H:708 P:158 L:0.54 V:0 TD:152) hit Thunderrune (H:198 P:169 L:0.66 V:0 TD:0) doing 152 damage
Thunderrune (H:198 P:169 L:0.66 V:0 TD:154) hit Touchbead (H:554 P:158 L:0.54 V:0 TD:152) doing 154 damage
Touchbead (H:554 P:158 L:0.54 V:0 TD:310) hit Thunderrune (H:40 P:169 L:0.66 V:0 TD:154) doing 158 damage
Coyotenight (H:618 P:184 L:0.94 V:0 TD:184) hit Touchbead (H:370 P:158 L:0.54 V:0 TD:310) doing 184 damage
Coyotenight (H:618 P:184 L:0.94 V:0 TD:368) hit Touchbead (H:186 P:158 L:0.54 V:0 TD:310) doing 184 damage
Thunderrune (H:40 P:169 L:0.66 V:0 TD:323) hit Touchbead (H:17 P:158 L:0.54 V:0 TD:310) doing 169 damage
Thunderrune (H:40 P:169 L:0.66 V:1 TD:492) hit Touchbead (H:-152 P:158 L:0.54 V:0 TD:310) doing 169 damage
    Touchbead is dead! :(
    2/3 fighters
Thunderrune (H:40 P:169 L:0.66 V:1 TD:646) hit Coyotenight (H:464 P:184 L:0.94 V:0 TD:368) doing 154 damage
Thunderrune (H:40 P:169 L:0.66 V:1 TD:804) hit Coyotenight (H:306 P:184 L:0.94 V:0 TD:368) doing 158 damage
Coyotenight (H:306 P:184 L:0.94 V:1 TD:552) hit Thunderrune (H:-144 P:169 L:0.66 V:1 TD:804) doing 184 damage
    Thunderrune is dead! :(
    1/3 fighters
Coyotenight (H:306 P:184 L:0.94 V:1 TD:552) wins!
```

## License

The [MIT License (MIT)](LICENCE.md)

Copyright (c) 2017 Scott Barr
