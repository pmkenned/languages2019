.PHONY: all run clean

JAVA_HOME=/usr/lib/jvm/java-11-openjdk-amd64

all: ex_TTYUtil.h ex/TTYUtil.class libttyutil.so Test.class

run: all
	LD_LIBRARY_PATH=. java Test

clean:
	rm -f *.o *.so *.class ex_TTYUtil.h

ex_TTYUtil.h ex/TTYUtil.class: ex/TTYUtil.java
	javac -cp . -h . ex/TTYUtil.java

ex_TTYUtil.o: ex_TTYUtil.c
	gcc -fPIC -c ex_TTYUtil.c -I $(JAVA_HOME)/include -I $(JAVA_HOME)/include/linux

libttyutil.so: ex_TTYUtil.o
	gcc ex_TTYUtil.o -shared -o libttyutil.so -Wl,-soname,ttyutil

Test.class: Test.java
	javac Test.java
