# makefile for logger
# 19 January 2018
# Code is licensed under the MIT License
# © 2018 Scott Isenberg

PACKAGE_NAME=logger
DONE=@echo -e $(GREEN)Done.$(NC)
RED='\033[0;31m'
GREEN='\033[0;32m'
WHITE='\033[1;37m'
PURPLE='\e[95m'
CYAN='\e[36m'
YELLOW='\033[1;33m'
ORANGE='\033[38;5;208m'
NC='\033[0m'

all : deps test

deps:
	@echo -e Grabbing dependencies...
	@go get github.com/logrusorgru/aurora
	$(DONE)

test:
	@echo -e Running Tests...
	@go test -v ./... | sed ''/'\(--- PASS\)'/s//$$(printf $(GREEN)---\\x20PASS)/'' | sed ''/PASS/s//$$(printf $(GREEN)PASS)/'' | sed  ''/'\(=== RUN\)'/s//$$(printf $(YELLOW)===\\x20RUN)/'' | sed ''/ok/s//$$(printf $(GREEN)ok)/'' | sed  ''/'\(--- FAIL\)'/s//$$(printf $(RED)---\\x20FAIL)/'' | sed  ''/FAIL/s//$$(printf $(RED)FAIL)/'' | sed ''/RUN/s//$$(printf $(YELLOW)RUN)/'' | sed ''/?/s//$$(printf $(ORANGE)?)/'' | sed ''/'\(^\)'/s//$$(printf $(NC))/''
	$(DONE)


.PHONY: all
