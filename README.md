# kubernetes informer example

This is a minimal example of a kubernetes informer for the [Pod](https://kubernetes.io/docs/concepts/workloads/pods/)
type.

It's purpose is to illustrate a working example of an informer for the purpose of learning how to write effective 
Kubernetes controllers. The informer pattern is used to avoid too many connections. It uses a function called `Watch`
that returns a Go channel, and then that channel receives updates whenever the watched resource changes.

`watcher.ResultChan()` returns a `<-chan watch.Event`

Where watch is the `k8s.io/apimachinery/pkg/watch` module.

The `context.Background()` returns a `context.Context` that is basically empty. It's there to make the API happy.

