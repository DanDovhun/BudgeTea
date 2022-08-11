build:
	go build -o ~/Desktop/BudgetManager/bin

run:
	go run main.go

build-and-run:
	make build
	./BudgetManager