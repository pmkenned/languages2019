all: Markov.java
	javac Markov.java

.PHONY: perf

perf: all
	./do_n.sh 10

.PHONY: clean

clean:
	rm *.class
