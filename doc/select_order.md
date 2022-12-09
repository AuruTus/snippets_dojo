
c3 first closed (following the order of `c3`, `c1`, `c2`)
```
	func() {
		c3 <- struct{}{}
		c1 <- struct{}{}
		c2 <- struct{}{}
	}
```

main thread (following the order of `c1`, `c2`, `c3`)
```
	select {
	case <-c1:
		cfmt.Printf(ctx, "c1 done\n")
	case <-c2:
		cfmt.Printf(ctx, "c2 done\n")
	case <-c3:
		cfmt.Printf(ctx, "c3 done\n")
	}
```

output (`c3 case` is chosen)
```
@./main.go 45 | 82a176c6-f1c8-453d-92a4-3ae5d3626620 | *snipets.SelectOrderTstr starts at: 1669861143s
./src/test_select_order.go 54: enter channel
./src/test_select_order.go 44: start cnt
./src/test_select_order.go 61: c3 done
./src/test_select_order.go 63: finish test
@./main.go 60 | 82a176c6-f1c8-453d-92a4-3ae5d3626620 | *snipets.SelectOrderTstr ends at 1669861145s; elapse: 2001492500ns
```