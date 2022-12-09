# Snipets Dojo

A coarse playground for snipets testing.

## testing utils

Some small tools to help construct the dojo. (under `./utils`)

## snipets

Main test codes here. (under `./src`)

* [arr_partition](./doc/arr_partition.md)
	> 1. Test different array partition implementations with swap and assignment;
	> 2. Test go generics and some function syntaxes;
* [close_chan](./doc/close_chan.md)
	> 1. Test if a channel will be drained after calling its `close()` function.
* [select_order](./doc/select_order.md)
	> 1. Test how the `case` expression order in `select` block affects the excecution order when all cases are true.