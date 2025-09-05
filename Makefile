MAKEFLAGS += -s

NAME = g2lint

all: clean
	@ printf "\033[0;34m[*]\033[0m Building g2lint...\n"

	@ mkdir -p build/

	@ go mod tidy
	@ go build -ldflags="-s -w" -o build/g2lint

	@ printf "\033[0;32m[*]\033[0m Built g2lint!\n"

clean:
	@ rm -rf build/

	@ go mod tidy