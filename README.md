# Gokedex - a command line Pokedex

Gokedex is a small project created to demonstrate understanding of HTTP clients, caching and REPL in Golang. It is a command-line application with various features. Under the hood, it utilises two internal libraries to faciliate API calls to the [PokéAPI](https://pokeapi.co/) and implement caching. 

## Getting Started

### Prerequisites

- Gokedex is built with go v1.23.5 so you must have, at minimum this version installed on your machine. The install guide can be found on the [Go website](https://go.dev/doc/install).

### Installing 

1. Download this repository to your machine using the command line method of your choice.

    - HTTP: `git clone https://github.com/SamW94/gokedex.git`
    - SSH: `git@github.com:SamW94/gokedex.git`
    - GitHub CLI: `gh repo clone SamW94/gokedex`

2. From the root of the repository, run the `go build` command to compile the source code into a binary file. You can also use `go install` from the same location to install all packages to your machine so you can access gokedex from anywhere with the `gokedex` command. 

## Command Reference

If you used `go install` to install Gokedex on your machine, simply run the `gokedex` command in your Terminal. If you used `go build` run the `./gokedex` command. 

Usage:

- *help*: Displays this help message
- *map*: Get the next page of locations
- *mapb*: Get the previous page of locations
- *explore*: Explore a location
- *catch*: Catch a Pokemon!
- *inspect*: See the attributes of Pokemon you've caught.
- *pokedex*: Display all of the Pokemon you have caught.
- *exit*: Exit the Pokedex

### help

Displays the help message on the command line, showing the above command reference.

```
Pokedex > help

Welcome to the Pokedex!
Usage:

explore <location_name>: Explore a location
map: Get the next page of locations
mapb: Get the previous page of locations
exit: Exit the Pokedex
help: Displays a help message
pokedex: Display all of the Pokemon you have caught.
inspect <pokemon>: See the attributes of Pokemon you've caught.
catch <pokemon>: Catch a Pokemon!
```

### map

Displays the next 20 location areas available to explore. 

```
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
oreburgh-mine-1f
oreburgh-mine-b1f
valley-windworks-area
eterna-forest-area
fuego-ironworks-area
mt-coronet-1f-route-207
mt-coronet-2f
mt-coronet-3f
mt-coronet-exterior-snowfall
mt-coronet-exterior-blizzard
mt-coronet-4f
mt-coronet-4f-small-room
mt-coronet-5f
mt-coronet-6f
mt-coronet-1f-from-exterior
```

### mapb 

Displays the previous 20 location areas available to explore. If you are on the first page of locations, this message will be displayed.

```
Pokedex > mapb
you're on the first page
```

If you are on a later page, you will see the previous 20 locations displayed by the `map` command. 

```
Pokedex > map
mt-coronet-1f-route-216
mt-coronet-1f-route-211
mt-coronet-b1f
great-marsh-area-1
great-marsh-area-2
great-marsh-area-3
great-marsh-area-4
great-marsh-area-5
great-marsh-area-6
solaceon-ruins-2f
solaceon-ruins-1f
solaceon-ruins-b1f-a
solaceon-ruins-b1f-b
solaceon-ruins-b1f-c
solaceon-ruins-b2f-a
solaceon-ruins-b2f-b
solaceon-ruins-b2f-c
solaceon-ruins-b3f-a
solaceon-ruins-b3f-b
solaceon-ruins-b3f-c

Pokedex > mapb
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
oreburgh-mine-1f
oreburgh-mine-b1f
valley-windworks-area
eterna-forest-area
fuego-ironworks-area
mt-coronet-1f-route-207
mt-coronet-2f
mt-coronet-3f
mt-coronet-exterior-snowfall
mt-coronet-exterior-blizzard
mt-coronet-4f
mt-coronet-4f-small-room
mt-coronet-5f
mt-coronet-6f
mt-coronet-1f-from-exterior
```

### explore

Run the `explore <location>` command to explore a location. This will return the Pokemon that live in that area.

```
Pokedex > explore canalave-city-area
Exploring canalave-city-area...
Found Pokemon: 
 - tentacool
 - tentacruel
 - staryu
 - magikarp
 - gyarados
 - wingull
 - pelipper
 - shellos
 - gastrodon
 - finneon
 - lumineon
 ```

### catch

Use the `catch` command to catch a Pokemon. The higher the Pokemon's base experience, the more difficult it is to catch! 

**Successful Catch**

```
Pokedex > catch sunkern
Throwing a Pokeball at sunkern...
sunkern was caught!
```

**Unsuccessful Catch**

```
Pokedex > catch mewtwo
Throwing a Pokeball at mewtwo...
mewtwo escaped!
```

If you attempt to catch a Pokemon you already have in your Pokedex, you will see this message:

```
Pokedex > catch pidgey
You already have a pidgey in your Pokedex
``` 

### inspect 

Use the `inspect` command to inspect any Pokemon you've caught. This will return the Pokemon's name, height, weight, stats and types like below. 

```
Pokedex > inspect pidgey
Name: pidgey
Height: 3
Weight: 18
Stats:
  -hp: 40
  -attack: 45
  -defense: 40
  -special-attack: 35
  -special-defense: 35
  -speed: 56
Types:
  - normal
  - flying
```

If you have not caught one of the Pokemon you're trying to inspect, you will see this error message:

```
Pokedex > inspect mewtwo
you have not caught that pokemon
```

### pokedex

Lists all of the Pokemon you've caught during this play session. 

```
Pokedex > pokedex
Your Pokedex:
  - pidgey
  - sunkern
```

### exit 

Closes the Gokedex application. 

```
Pokedex > exit
Closing the Pokedex... Goodbye!
```

## Internal Libraries

Gokedex makes use of 2 internal packages: `pokeapi` and `pokecache`. The source code is availabile in the respective package directories under `./internal`

### pokeapi

The `pokeapi` package provides methods for interacting with various endpoints provided by [PokéAPI](https://pokeapi.co/). It provides types for unmarshalling the responses from these endpoints to Go structs, so that you can more easily interact with the JSON data returned in Golang programs.

### pokecache

The `pokecache` package implements a cache for the data provided by [PokéAPI](https://pokeapi.co/). Each cache object is reaped after 1 minute of time to ensure it manages memory usage itself. This makes the program faster, as an API call is not required for every single attempt to retrieve information about a location or Pokemon. 