build:
	if [ -d "./build" ]; then \
		rm -rf ./build; \
	fi
	go build -o build/gols
