# pathfinder
Project part of Hive Helsinki curriculum

- Many trains moving at the tracks at the same time.
- Only one train can be on a single station at a time.

Main components identified: 
- Path finding algorithm
- Routing (which can utilize path finding if needed)
- Grid creation
- Grid validation
- Predetermining strategy
- Errors
- Benchmarking 

## Path finding algorithm

- Fastest route
- Routes that optimize concurrency
- Needs to support dynamic maps based on the tics
- e.g. Many trains out - who leads shortest paths based on unblocked routes 


## Routing
- Aknowledges routes unavailable at any tics due to other trains (part of the path finding)
- Takes input (to be determined) for optimizations. e.g. Which kind of algorithm to prefer. E.g. Only 1 route in or out -> all trains after one other fastest route. 


## Validation 
- map is valid (routes, stations, no overlaps, )


## Errors 
- Check the extensive edge case list

## Grid 
- preliminary format - Linked list (tbd)
- Metadata to save for decision making (tbd)

## Generic structure and program flow 
- Create grid
- Validate grid
- Predetrmine strategy
- Routing (generate ticks for the train) - This section controls all trains  
- Find paths (use ticks to see if certain routes are available at any given time.) - This section finds routes for trains based on constraints 

## Potential optimizations 
- preprocessing grid to leave out deadends etc.
