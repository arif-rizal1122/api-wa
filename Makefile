# Variabel untuk menyimpan nama file utama
MAIN_FILE := main.go

# Target 'run' untuk mengompilasi dan menjalankan program Go
.PHONY: run
run:
    go run $(MAIN_FILE)

# Target 'build' untuk mengompilasi program Go menjadi executable
.PHONY: build
build:
    go build -o main $(MAIN_FILE)

# Target 'clean' untuk menghapus file executable
.PHONY: clean
clean:
    rm -f main

# Target 'all' sebagai alias untuk target 'build'
.PHONY: all
all: build