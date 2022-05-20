.PHONY: all run clean

BUILD_DIR = ./build
TARGET = prog
SRC = $(wildcard *.rs)

all: $(BUILD_DIR)/$(TARGET)

$(BUILD_DIR)/$(TARGET): $(SRC)
	mkdir -p $(BUILD_DIR)
	rustc -o $@ $<

run: all
	$(BUILD_DIR)/$(TARGET)

clean:
	rm -rf $(BUILD_DIR)
