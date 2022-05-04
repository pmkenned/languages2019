.PHONY: all run perf clean

TARGET = Hello

all: $(TARGET).class

%.class: %.java
	javac $<

run: all
	java $(TARGET)

perf: all
	./do_n.sh 10

clean:
	rm *.class
