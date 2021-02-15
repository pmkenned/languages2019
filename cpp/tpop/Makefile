CXX=g++
CXXFLAGS=-Wall -Wextra -Werror -O2
TARGET=markov
BUILD_DIR=./build

SRC = markov.cpp
OBJ = $(SRC:%.cpp=$(BUILD_DIR)/%.o)

.PHONY: all run perf clean

all: $(BUILD_DIR)/$(TARGET)

run: all
	$(BUILD_DIR)/$(TARGET) < Genesis.txt

$(BUILD_DIR)/%.o: %.cpp
	mkdir -p $(BUILD_DIR)
	$(CXX) $(CXXFLAGS) -c $< -o $@

$(BUILD_DIR)/$(TARGET): $(OBJ)
	mkdir -p $(BUILD_DIR)
	$(CXX) $(CXXFLAGS) $^ -o $@

perf: all
	./do_n.sh 10

clean:
	rm -rf $(BUILD_DIR)
