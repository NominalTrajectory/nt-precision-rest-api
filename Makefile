ifneq ($(DBADDR),)
ADDRFLAG = -dbaddr=$(DBADDR)
endif

.PHONY: start
start:
	@go build
	@./gorm '$(ADDRFLAG)'