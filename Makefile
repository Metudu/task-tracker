build:
	@echo "Building the application..."
	@go build -o task-tracker
	@rm ./task/task.json
	@touch ./task/task.json
	@echo "Builded."