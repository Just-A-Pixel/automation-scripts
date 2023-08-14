WORK_DIR = $(shell pwd)

TARGET ?= ""
SCRIPT_NAME ?= ""

build:
	echo $(WORK_DIR)/$(TARGET)
	go build $(TARGET)
	mv script script_$(SCRIPT_NAME)
	go run cron/cron.go $(WORK_DIR)/script_$(SCRIPT_NAME)

default: build