# Deck of Cards - Gophercise 9

We were meant to create a package that would allow us to interact with a deck of cards, with the following requirements:

- An option to sort the cards with a user-defined comparison function. The sort package in the standard library can be used here, and expects a less(i, j int) bool function.
- A default comparison function that can be used with the sorting option.
- An option to shuffle the cards.
- An option to add an arbitrary number of jokers to the deck.
- An option to filter out specific cards. Many card games are played without 2s and 3s, while others might filter out other cards. We can provide a generic way to handle this as an option.
- An option to construct a single deck composed of multiple decks. This is used often enough in games like blackjack that having an option to build a deck of cards with say 3 standard decks can be useful.

Also, we must be able to generate new decks.

*Note*: We were not meant to create a Deck type but I could really wrap my head around how it should work better with such a type present, so I chose to implement one for my personal attempt at the task.

The package has a fair amount of tests as well.

## Build
PowerShell scripts are included to run the go tests and to append the automatically generated go doc to readme.md.
To run the build, simply cd into this directory and enter the following:
```powershell
.\build.ps1
```

## Documentation
I made an effort to provide the comments necessary for go doc to be helpful. The documentation can be output into your terminal with the following command:
```
go doc .\deck
```
OR
```
go doc ./deck
```


