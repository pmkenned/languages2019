.PHONY: all run clean

all: RVSim.class

%.class: %.scala
	scalac $<

run: all
	scala RVSim insert_sort.hex

clean:
	rm -f *.class
