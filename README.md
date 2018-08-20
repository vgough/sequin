# Workflow Engine

The Workflow Engine is a library for constructing, testing, and running workflows.

Workflows are written in Go and look a lot like any other Go program.  The
primary difference is that a workflow is durable and can live beyond the life
of a single worker.  Durability impacts code by requiring that workflow state
be serialized and have well-defined synchronization points.

## Status

WIP!

Working on developer experience and core functionality.


## TODO

* Status reporting / UI
  Workflows should have a way to report how to represent themselves visually.
  This could look like a graph language, specifying how to logically connect
  the components.  Use case, workflow with backup cases (switch-style).
* Progress estimation
* Testing infrastructure
  Test alternative paths through code.  Eg with/without cache.

