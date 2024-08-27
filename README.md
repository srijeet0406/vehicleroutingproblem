# Vehicle Routing Problem

This document gives the general description of the project, along with steps to run and test the code.

The Vehicle Routing Problem (VRP) aims to route an unlimited number of vehicles to complete a list of loads in an optimized manner.
Optimization is necessary because for each additional vehicle, we induce an extra cost (500 in this case). Optimization is also necessary
because for each additional trip that a vehicle makes, we induce the time to drive from the previous location to the new location, plus the 
time to come back to base (0,0).

*Note* All vehicles must start and end at the base location.

## Compiling the code

After you clone the repo, run the following command from inside the directory where you have cloned the repository.
```
go build
```

## Testing the code
Run the following command from inside the directory where you have cloned the repository.
```
go test
```

## Running the application
After you've compiled the code, run the following command to run the application.
```
./vehicleroutingproblem <path to the input file name>
```

## Testing the application against the given test data set
```
python3 evaluateShared.py --cmd "./vehicleroutingproblem" --problemDir <Path to Training Problems directory>
```
Here, ``vehicleroutingproblem`` is the binary that the application produces as a result of compiling the code.

The input file has a list of loads, like this:
```text
loadNumber pickup dropoff
1 (-50.1,80.0) (90.1,12.2)
2 (-24.5,-19.2) (98.5,1.8)
3 (0.3,8.9) (40.9,55.0)
4 (5.3,-61.1) (77.8,-5.4)
```

## Explanation of the approach used in the algorithm
A "greedy" approach (also known as nearest neighbor) is used in the presented algorithm to solve the VRP. What this means is that we greedily pick whichever load is closest for the
current driver, and then repeat the process until either all loads have been picked up or the driver has reached the end of their shift (12 hours).
This approach, while sufficient for the current use cases, may not prove to be the best when considering more complex problems.
For example, consider the following use case:
Say we have a load ``load1`` at ``(100,100)`` that needs to be picked up, and all the free drivers are > 1 hour away.
However, there is a driver ``driver1`` at ``(99,99)`` who is about to finish up a load drop in the next minute.
Our current algorithm would assign ``load1`` to one of the free drivers who is more than an hour away from the pickup location, instead of waiting for
``driver1`` to finish up their load and pickup ``load1`` once they are done.

## References
[Wikipedia](https://en.wikipedia.org/wiki/Vehicle_routing_problem)

[Wikipedia Metaheuristic](https://en.wikipedia.org/wiki/Metaheuristic)

[Travelling Salesman Problem](https://www.routific.com/blog/travelling-salesman-problem)