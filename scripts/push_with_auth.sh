#!/bin/bash

echo "🚀 Pushing to both repositories with authentication..."

# Prompt for credentials
read -p "Enter your Gitea username (cntalouk): " gitea_user
gitea_user=${gitea_user:-cntalouk}

read -s -p "Enter your Gitea password/token: " gitea_pass
echo

read -p "Enter your GitHub username (cntal4): " github_user
github_user=${github_user:-cntal4}

read -s -p "Enter your GitHub password/token: " github_pass
echo

echo ""
echo "📤 Pushing to Gitea..."
if git push https://${gitea_user}:${gitea_pass}@platform.zone01.gr/git/cntalouk/go-reloaded.git main; then
    echo "✅ Successfully pushed to Gitea"
else
    echo "❌ Failed to push to Gitea"
fi

echo ""
echo "📤 Pushing to GitHub..."
if git push https://${github_user}:${github_pass}@github.com/cntal4/go-reloaded.git main; then
    echo "✅ Successfully pushed to GitHub"
else
    echo "❌ Failed to push to GitHub"
fi

echo ""
echo "🎉 Push operation completed!"