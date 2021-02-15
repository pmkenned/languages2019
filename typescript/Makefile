.PHONY: all clean

all: hello.js

%.js: %.ts
	tsc $<

clean:
	rm -f *.js
