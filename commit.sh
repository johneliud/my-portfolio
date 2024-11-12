#!/bin/bash

# Run gofmt to format the code
echo "Running gofmt..."
gofmt -w -s .

# Prompt user to choose files to add
echo "Please specify the files you want to add (space-separated):"
read -r files

# Check if files were provided
if [ -z "$files" ]; then
  echo "No files specified. Exiting."
  exit 1
fi

# Add specified files to git
git add $files

# Commit changes with a message
if [ -z "$1" ]; then
  echo "Please provide a commit message."
  exit 1
fi

git commit -m "$1"

# Push changes
git push

echo "Selected files committed and pushed successfully."
