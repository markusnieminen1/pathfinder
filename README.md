# pathfinder

![alt text](output.gif)
Project part of Hive Helsinki curriculum

## Project 

In this project our task was to implement pathfinding algorithm to route trains from one station to another. There are multiple trains leaving the same start station and will head towards same end station.  **The GOAL is to find combination of routes that minimises the movements needed in total**. 

Main restriction: **Only one train at any station at any given time (except end and start)** 

The program treats all routes to be equal lenght. So connection from x,y 0, 0 -> 100,100 will travel equally fast as 0,0 -> 0,1 

## How to use the pathfinder

- Clone the project from github or gitea.
- Go to the cloned directory. 

Run the project using the following syntax:  
```bash 
go run . [path to file containing network map] [start station] [end station] [number of trains]
```
Example: 
```bash 
go run . test_files/network_example.map waterloo st_pancras 4
```

Argument flags: 

-v for visualisation. The project can draw the grid to the browser and show how the algorithm works. 

-aid for Allow Invalid Data. The program tries to create the map even with overlapping connections and stations. THIS CAN CAUSE UNEXPECTED BEHAVIOUR!!!

### Input file format
#### Stations (The nodes)

This section lists the name of each station followed by two numbers (integers). These numbers are X and Y coordinates that show exactly where the station is located on a flat grid or map.

Example: waterloo, 3, 1 means the station is named Waterloo and sits at grid position (3, 1).

Normally only one station per coordinate combination is allowed. 

#### Connections

This section shows which stations are linked together. It uses a dash to show that a train track runs directly between two places.

Example: waterloo - victoria means you can travel directly between Waterloo and Victoria.

Normally connection between 2 stations can be declared only once! 

Example of an error case: 
...
Connections: 
waterloo - victoria
victoria - waterloo
...


#### General 
Everything after \# is a comment and the computer ignores it! Whitespaces are ignored. Stations or Connections can only be written on the same line! 

The example file has "Stations:" and "Connections:" prefixes for sections for the file. The program does not need prefixes in order to work. Lines and connections can also be mixed.

## What we learned

to be discussed! 


## Initial planning notes 
Main components identified: 
- Path finding algorithm
- Routing (which can utilize path finding if needed)
- Grid creation
- Grid validation
- Predetermining strategy
- Errors
- Benchmarking 

### Path finding algorithm

- Fastest route
- Routes that optimize concurrency
- Needs to support dynamic maps based on the tics
- e.g. Many trains out - who leads shortest paths based on unblocked routes 


### Routing
- Aknowledges routes unavailable at any tics due to other trains (part of the path finding)
- Takes input (to be determined) for optimizations. e.g. Which kind of algorithm to prefer. E.g. Only 1 route in or out -> all trains after one other fastest route. 


### Validation 
- map is valid (routes, stations, no overlaps, )


### Errors 
- Check the extensive edge case list

### Grid 
- preliminary format - Linked list (tbd)
- Metadata to save for decision making (tbd)

### Generic structure and program flow 
- Create grid
- Validate grid
- Predetrmine strategy
- Routing (generate ticks for the train) - This section controls all trains  
- Find paths (use ticks to see if certain routes are available at any given time.) - This section finds routes for trains based on constraints 

### Potential optimizations 
- preprocessing grid to leave out deadends etc.


### Guidelines for the code 
- Use packages 
- Small, reusable functions
- 1 - 5 functions per file ideally 
- filename should reflect what kind of functions it contains 
- Carry the err to main, so almost all files should return (something, error)


#### Input format

The input file needs to have connections and stations marked properly. 

#### Stations 

Every station will have a name, x coordinate and y coordinate 

For a row to be valid, the station will have 2 commas in total. All spaces will be truncated. 

name,x,y