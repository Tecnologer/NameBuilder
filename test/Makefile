count = 100

bench:
	go test -bench=GetRandom -benchtime=$(count)x

100:
	count=100
	make bench

1000:
	count=1000
	make bench

10000:
	count=10000
	make bench

100000:
	count=100000
	make bench

1000000:
	count=1000000
	make bench