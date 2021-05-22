.PHONY: all
all: benchmark

.PHONY: benchmark
benchmark: testdata/canterbury
	go test -benchmem -bench .

testdata/canterbury: testdata/cantrbry.tar.gz
	mkdir -p $@
	tar -xvf $^ -C $@

.PHONY: clean
clean:
	rm -rf testdata/canterbury
