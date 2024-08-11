#!/bin/bash

# Define the target directory where the binaries will be placed
TARGET_DIR="/Users/romittajale/Documents/2024/tajalevegsuppliers"

# Navigate to the directory containing both applications
cd /Users/romittajale/Documents/2024/tajalevegsuppliers

# Build the dashboard application and place the binary in the target directory
echo "Building dashboard..."
cd dashboard
go build -o $TARGET_DIR/dashboard_app
if [ $? -eq 0 ]; then
    echo "Dashboard built successfully."
else
    echo "Failed to build dashboard."
    exit 1
fi

# Build the myoneapp application and place the binary in the target directory
echo "Building myoneapp..."
cd ../myoneapp
go build -o $TARGET_DIR/myoneapp_app
if [ $? -eq 0 ]; then
    echo "myoneapp built successfully."
else
    echo "Failed to build myoneapp."
    exit 1
fi

# Run both applications in parallel from the target directory
echo "Running both applications..."
$TARGET_DIR/dashboard_app &   # Running the dashboard in the background
$TARGET_DIR/myoneapp_app &     # Running myoneapp in the background

# Wait for both applications to finish
wait

echo "Both applications have stopped."
