#!/usr/bin/env bash

# Get the highest tag number across ALL branches (not just current commit history)
VERSION=`git tag --sort=-version:refname | head -1`
VERSION=${VERSION:-'0.0.0'}

# Get number parts
MAJOR="${VERSION%%.*}"; VERSION="${VERSION#*.}"
MINOR="${VERSION%%.*}"; VERSION="${VERSION#*.}"
PATCH="${VERSION%%.*}"; VERSION="${VERSION#*.}"
VERSION_TYPE="beta"

# Increase version
PATCH=$((PATCH+1))

# Create new tag.
NEW_TAG="$MAJOR.$MINOR.$PATCH-$VERSION_TYPE"

echo "Updating to $NEW_TAG"
git tag -s -m "$NEW_TAG" "$NEW_TAG"

# Push the new tag to remote.
git push origin "$NEW_TAG"