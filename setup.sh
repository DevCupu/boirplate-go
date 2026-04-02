#!/bin/bash

# Setup script untuk boilerplate Go

set -e

echo "🚀 Boilerplate Go - Setup Script"
echo "=================================="
echo ""

# Check Go installation
echo "✓ Checking Go installation..."
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.22+"
    exit 1
fi

GO_VERSION=$(go version | awk '{print $3}' | cut -d'.' -f2)
if [ "$GO_VERSION" -lt 22 ]; then
    echo "❌ Go version must be 1.22 or higher"
    exit 1
fi

echo "✓ Go version: $(go version | awk '{print $3}')"
echo ""

# Create .env if not exists
echo "✓ Setting up .env file..."
if [ ! -f .env ]; then
    if [ -f .env.example ]; then
        cp .env.example .env
        echo "  Created .env from .env.example"
        echo "  ⚠️  Please update .env with your actual database credentials"
    fi
else
    echo "  .env already exists"
fi
echo ""

# Download dependencies
echo "✓ Downloading dependencies..."
go mod download
go mod tidy
echo "  Dependencies downloaded"
echo ""

# Check Docker (optional)
if command -v docker &> /dev/null; then
    echo "✓ Docker is installed: $(docker --version)"
    echo ""
    
    read -p "Do you want to setup database with Docker Compose? (y/n) " -n 1 -r
    echo ""
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        echo "Starting Docker Compose..."
        docker-compose up -d
        echo "✓ Database started with Docker Compose"
        echo ""
    fi
else
    echo "⚠️  Docker not found (optional)"
    echo "   To use Docker Compose for database setup, install Docker from https://www.docker.com"
    echo ""
fi

# Check air for development
echo "✓ Checking air for hot reload..."
if ! command -v air &> /dev/null; then
    echo "  Installing air..."
    go install github.com/cosmtrek/air@latest
    echo "  ✓ Air installed"
else
    echo "  ✓ Air already installed"
fi
echo ""

# Summary
echo "=================================="
echo "✅ Setup Complete!"
echo ""
echo "Next steps:"
echo "  1. Update .env with your configuration"
echo "  2. Run database migrations: make migrate-up"
echo "  3. Start development: make dev"
echo ""
echo "For more info, see README.md"
