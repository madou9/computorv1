Quick Start
Pull the official Go image from Docker Hub:

bash
docker pull golang:latest
Run your Go code directly from your project directory. The command below will start an ephemeral container, mount your current directory, and execute your code:

Method: Interactive Development Shell
For longer development sessions, it's often easier to open an interactive shell inside the container. This allows you to run multiple go commands without typing the long docker run command each time.

bash
# Start an interactive Bash shell in the Go container
docker run -it --rm -v "$(pwd)":/app -w /app golang bash
