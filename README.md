# kubernetes informer example

This is a minimal example of a kubernetes informer for the [Pod](https://kubernetes.io/docs/concepts/workloads/pods/)
type.

It's purpose is to illustrate a working example of an informer for the purpose of learning how to write effective 
Kubernetes controllers. The informer pattern is used to avoid too many connections. It uses a function called `Watch`
that returns a Go channel, and then that channel receives updates whenever the watched resource changes.
