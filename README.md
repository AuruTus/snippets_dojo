# Snippets Dojo

A coarse playground for snippets testing.

## testing utils

Some small tools to help construct the dojo. (under `./utils`)

## snippets

Main test codes here. (under `./src/test_snippets`)

* [arr_partition](./doc/arr_partition.md)
	> 1. Test different array partition implementations with swap and assignment;
	> 2. Test go generics and some function syntaxes;
* [close_chan](./doc/close_chan.md)
	> 1. Test if a channel will be drained after calling its `close()` function.
* [select_order](./doc/select_order.md)
	> 1. Test how the `case` expression order in `select` block affects the excecution order when all cases are true.
* [rw_lock](./doc/rw_lock.md)
	> 1. Test rw lock with writer priority
	> 2. Unfinished. The former thoughts of an implementation of channel have some bugs. 
	Can refer to the `sync.RWMutex`, which seems just use the mutex.
* [chan_range](./doc/chan_range.md)
	> 1. Test when the range on a channel will complete (channel is closed or the scenario like program is end)
	> 2. TODO: md record
* [type_constraints](./doc/type_constraints_and_interface.md)
	> 1. Test the differences between type constraints and composite interface. And try to use them when mixing them together.
	> 2. TODO: md record