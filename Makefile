.PHONY: help
help:
	@echo "Usage:"
	@echo "    install         Install protoc programs"
	@echo "    generate        Generate proto files"
	@echo "    update          Perform go mod init or update"
	@echo "    clean           Clean all generated proto files"

.PHONY: install
install:
	./proto-install.sh

.PHONY: generate
generate:
	./proto-gen.sh

.PHONY: clean
clean:
	rm -rf lib/*
