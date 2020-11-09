CXX=g++
CFLAGS=-Wall -Wextra -Werror -O2
TARGET=markov

SRC = markov.cpp
OBJ = $(SRC:%.cpp=%.o)

%.o: %.cpp
	$(CXX) $(CFLAGS) -c $< -o $@

$(TARGET): $(OBJ)
	$(CXX) $(CFLAGS) $^ -o $@

all: $(TARGET)

.PHONY: perf

perf: all
	./do_n.sh 10

.PHONY: clean

clean:
	rm -f *.o $(TARGET)
