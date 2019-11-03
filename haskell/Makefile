TARGET_EXEC	    := hello

BUILD_DIR ?= ./build
SRC_DIRS ?= ./src

SRCS := $(shell find $(SRC_DIRS) -name *.hs)

$(BUILD_DIR)/$(TARGET_EXEC): $(SRCS)
	$(MKDIR_P) $(dir $@)
	ghc -o $@ $(SRCS)

.PHONY: clean

clean:
	$(RM) -r $(BUILD_DIR)

MKDIR_P ?= mkdir -p
