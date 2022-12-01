As the output shows, with signals sent to channel and then closed, the channel will drain itself and finally output the `close` signal (`!ok` boolean type return value of `<-` operator)

```
@./main.go 45 | 3acd02fd-11ee-4f63-a4e9-1dae51a81f2f | *snipets.CloseChanTstr starts at: 1669860998s
child routine received! rcv_cnt: 1 
complete sending inputs; start closing channel; 
complete closing channel; 
child waiting; buffured len: 4 
child routine received! rcv_cnt: 2 
child waiting; buffured len: 3 
child routine received! rcv_cnt: 3 
child waiting; buffured len: 2 
child routine received! rcv_cnt: 4 
child waiting; buffured len: 1 
child routine received! rcv_cnt: 5 
child waiting; buffured len: 0 
discard 0 inputs; quit goroutine; 
@./main.go 60 | 3acd02fd-11ee-4f63-a4e9-1dae51a81f2f | *snipets.CloseChanTstr ends at 1669861005s; elapse: 7503933500ns
```