.PHONY: all run clean

all: hello

hello: Hello.class

%.class: %.scala
	scalac $<

run: all
	scala Hello

clean:
	rm -f *.class
