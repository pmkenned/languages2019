.PHONY: all run clean

JAVA_HOME=/usr/lib/jvm/java-11-openjdk-amd64

all: tty_TTYUtil.h tty/TTYUtil.class libttyutil.so JEdit.class

run: all
	LD_LIBRARY_PATH=. java JEdit

clean:
	rm -f tty_TTYUtil.h
	find . \( -name "*.class" -o -name "*.o" -o -name "*.so" \) -exec rm {} \;

tty_TTYUtil.h tty/TTYUtil.class: tty/TTYUtil.java
	javac -cp . -h . tty/TTYUtil.java

tty/TTYUtil.class: tty_TTYUtil.h

tty_TTYUtil.o: tty_TTYUtil.cpp
	g++ -fPIC -c tty_TTYUtil.cpp -I $(JAVA_HOME)/include -I $(JAVA_HOME)/include/linux

libttyutil.so: tty_TTYUtil.o
	g++ tty_TTYUtil.o -shared -o libttyutil.so -Wl,-soname,ttyutil

JEdit.class: JEdit.java
	javac JEdit.java
