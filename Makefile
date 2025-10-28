# Define Go command and flags
GO = go

# Define the target executable
TARGET = computor

# Define target: build the executable
all : $(TARGET)

# Rule to build the target executable
$(TARGET): computor.go
	$(GO) build -o $(TARGET) computor.go

re: fclean all

# clean target: remove the target executable
clean:
	rm -f $(TARGET)

fclean: clean
	rm -f $(TARGET)

.PHONY: all clean fclean re
