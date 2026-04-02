@echo off
REM Setup script untuk boilerplate Go (Windows)

echo.
echo 🚀 Boilerplate Go - Setup Script (Windows)
echo ============================================
echo.

REM Check Go installation
echo Checking Go installation...
go version >nul 2>&1
if errorlevel 1 (
    echo ❌ Go is not installed. Please install Go 1.22+
    exit /b 1
)

for /f "tokens=3" %%i in ('go version') do set GO_VERSION=%%i
echo ✓ Go version: %GO_VERSION%
echo.

REM Create .env if not exists
echo Setting up .env file...
if not exist .env (
    if exist .env.example (
        copy .env.example .env
        echo ✓ Created .env from .env.example
        echo ⚠️  Please update .env with your database credentials
    )
) else (
    echo ✓ .env already exists
)
echo.

REM Download dependencies
echo Downloading dependencies...
call go mod download
call go mod tidy
echo ✓ Dependencies downloaded
echo.

REM Check Docker (optional)
docker --version >nul 2>&1
if errorlevel 0 (
    echo ✓ Docker is installed
    echo.
    
    set /p DOCKER_CHOICE="Do you want to setup database with Docker Compose? (y/n): "
    if /i "%DOCKER_CHOICE%"=="y" (
        echo Starting Docker Compose...
        call docker-compose up -d
        echo ✓ Database started
        echo.
    )
) else (
    echo ⚠️  Docker not found (optional)
    echo   To use Docker Compose, install from https://www.docker.com
    echo.
)

REM Summary
echo ============================================
echo ✅ Setup Complete!
echo.
echo Next steps:
echo   1. Update .env with your configuration
echo   2. Run database migrations: make migrate-up
echo   3. Start development: make run
echo.
echo For more info, see README.md
echo.
pause
