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
3 fighters enter the arena.
    Rabbitseed (H:957 P:124 L:0.39 V:0 TD:0)
    Ogreeast (H:395 P:129 L:0.43 V:0 TD:0)
    Coyotequasar (H:785 P:114 L:0.67 V:0 TD:0)

Rabbitseed (H:957 P:124 L:0.39 V:0 TD:124) hit Coyotequasar (H:661 P:114 L:0.67 V:0 TD:0) doing 124 damage
Ogreeast (H:395 P:129 L:0.43 V:0 TD:99) hit Coyotequasar (H:562 P:114 L:0.67 V:0 TD:0) doing 99 damage
Coyotequasar (H:562 P:114 L:0.67 V:0 TD:111) hit Rabbitseed (H:846 P:124 L:0.39 V:0 TD:124) doing 111 damage
Rabbitseed (H:846 P:124 L:0.39 V:0 TD:185) hit Ogreeast (H:334 P:129 L:0.43 V:0 TD:99) doing 61 damage
Coyotequasar (H:562 P:114 L:0.67 V:0 TD:225) hit Rabbitseed (H:732 P:124 L:0.39 V:0 TD:185) doing 114 damage
Ogreeast (H:334 P:129 L:0.43 V:0 TD:212) hit Coyotequasar (H:449 P:114 L:0.67 V:0 TD:225) doing 113 damage
Ogreeast (H:334 P:129 L:0.43 V:0 TD:318) hit Rabbitseed (H:626 P:124 L:0.39 V:0 TD:185) doing 106 damage
Ogreeast (H:334 P:129 L:0.43 V:0 TD:404) hit Coyotequasar (H:363 P:114 L:0.67 V:0 TD:225) doing 86 damage
Rabbitseed (H:626 P:124 L:0.39 V:0 TD:309) hit Ogreeast (H:210 P:129 L:0.43 V:0 TD:404) doing 124 damage
Rabbitseed (H:626 P:124 L:0.39 V:0 TD:432) hit Coyotequasar (H:240 P:114 L:0.67 V:0 TD:225) doing 123 damage
Coyotequasar (H:240 P:114 L:0.67 V:0 TD:329) hit Rabbitseed (H:522 P:124 L:0.39 V:0 TD:432) doing 104 damage
Ogreeast (H:210 P:129 L:0.43 V:0 TD:485) hit Rabbitseed (H:441 P:124 L:0.39 V:0 TD:432) doing 81 damage
Ogreeast (H:210 P:129 L:0.43 V:0 TD:586) hit Rabbitseed (H:340 P:124 L:0.39 V:0 TD:432) doing 101 damage
Rabbitseed (H:340 P:124 L:0.39 V:0 TD:498) hit Ogreeast (H:144 P:129 L:0.43 V:0 TD:586) doing 66 damage
Rabbitseed (H:340 P:124 L:0.39 V:0 TD:593) hit Coyotequasar (H:145 P:114 L:0.67 V:0 TD:329) doing 95 damage
Coyotequasar (H:145 P:114 L:0.67 V:0 TD:443) hit Ogreeast (H:30 P:129 L:0.43 V:0 TD:586) doing 114 damage
Coyotequasar (H:145 P:114 L:0.67 V:0 TD:557) hit Rabbitseed (H:226 P:124 L:0.39 V:0 TD:593) doing 114 damage
Ogreeast (H:30 P:129 L:0.43 V:0 TD:715) hit Coyotequasar (H:16 P:114 L:0.67 V:0 TD:557) doing 129 damage
Rabbitseed (H:226 P:124 L:0.39 V:1 TD:717) hit Ogreeast (H:-94 P:129 L:0.43 V:0 TD:715) doing 124 damage
    Ogreeast is dead! :(
    2/3 fighters left.
Coyotequasar (H:16 P:114 L:0.67 V:0 TD:671) hit Rabbitseed (H:112 P:124 L:0.39 V:1 TD:717) doing 114 damage
Rabbitseed (H:112 P:124 L:0.39 V:2 TD:804) hit Coyotequasar (H:-71 P:114 L:0.67 V:0 TD:671) doing 87 damage
    Coyotequasar is dead! :(
    1/3 fighters left.
Rabbitseed (H:112 P:124 L:0.39 V:2 TD:804) wins!
```


## License

The [MIT License (MIT)](LICENCE.md)

Copyright (c) 2017 Scott Barr
