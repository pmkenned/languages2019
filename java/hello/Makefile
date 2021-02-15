.PHONY: all perf clean

all: Hello.java
	javac Hello.java

perf: all
	./do_n.sh 10

clean:
	rm *.class
