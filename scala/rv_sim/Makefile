.PHONY: all run clean

all: RVSim.class

%.class: %.scala
	scalac $<

run: all
	scala RVSim

clean:
	rm -f *.class
