# Theories

This repo contains theories that I have sought an answer to in go by writing tests or benchmarks.

These can be executed by running the `run` script in the root of this repository.

## Tested

These theories have tests or benchmarks that confirm the theory as True or reject it as False.

### True

- It is faster to assign to a variable instead of calling the same function multiple times.
- If a value returned from a function is only used once then it should not be stored in a variable but instead the function called where it is used.

### False

## Untested

- SQLite on an ssd is faster than Redis using AOF
