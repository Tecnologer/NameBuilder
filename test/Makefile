COUNT=100

.SILECT: all

bench:
	go test -bench=GetRandom -benchtime=$(COUNT)x

all:
	@make -s --no-print-directory 100
	@make -s --no-print-directory 1000
	@make -s --no-print-directory 10000
	@make -s --no-print-directory 100000
	@make -s --no-print-directory 1000000

100:
	@make -s --no-print-directory bench COUNT=100

1000:
	@make -s --no-print-directory bench COUNT=1000

10000:
	@make -s --no-print-directory bench COUNT=10000

100000:
	@make -s --no-print-directory bench COUNT=100000

1000000:
	@make -s --no-print-directory bench COUNT=1000000