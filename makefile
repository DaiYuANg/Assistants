all:build

build:
	poetry run nuitka --standalone --output-dir=dist main.py

clean:
	rm -rf build
	rm -rf dist

format:
	poetry run black .