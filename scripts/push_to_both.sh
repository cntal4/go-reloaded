#!/bin/bash

# Push to both Gitea and GitHub repositories
# This script handles authentication for both remotes

set -e

echo "🚀 Pushing to both Gitea and GitHub repositories..."

# Check if we have uncommitted changes
if ! git diff-index --quiet HEAD --; then
    echo "❌ You have uncommitted changes. Please commit them first."
    exit 1
fi

# Push to Gitea (origin)
echo "📤 Pushing to Gitea (origin)..."
if git push origin main; then
    echo "✅ Successfully pushed to Gitea"
else
    echo "❌ Failed to push to Gitea. Please check your credentials."
    echo "💡 You may need to configure your Gitea credentials:"
    echo "   git config credential.helper store"
    echo "   Or use SSH keys for authentication"
fi

echo ""

# Push to GitHub
echo "📤 Pushing to GitHub..."
if git push github main; then
    echo "✅ Successfully pushed to GitHub"
else
    echo "❌ Failed to push to GitHub. Please check your credentials."
    echo "💡 You may need to configure your GitHub credentials:"
    echo "   - Use a Personal Access Token"
    echo "   - Or configure SSH keys"
    echo "   - Or run: gh auth login (if you have GitHub CLI)"
fi

echo ""
echo "🎉 Push operation completed!"
echo "📋 Summary:"
echo "   - Gitea: https://platform.zone01.gr/git/cntalouk/go-reloaded"
echo "   - GitHub: https://github.com/cntal4/go-reloaded"