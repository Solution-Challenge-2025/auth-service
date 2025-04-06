# Set working directory to the script's location
$scriptPath = Split-Path -Parent $MyInvocation.MyCommand.Path
Set-Location -Path $scriptPath

# Common headers
$headers = @{"Content-Type" = "application/json"}

# Test registration endpoint
Write-Host "`nTesting registration endpoint..."
$registerData = @{
    username = "testuser"
    password = "testpass123"
    email = "test@example.com"
} | ConvertTo-Json

try {
    $response = Invoke-WebRequest -Uri "http://localhost:8080/api/v1/auth/register" -Method Post -Headers $headers -Body $registerData -TimeoutSec 5
    Write-Host "Registration successful: $($response.Content)"
} catch {
    Write-Host "Registration failed: $($_.Exception.Message)"
}

# Test login endpoint
Write-Host "`nTesting login endpoint..."
$loginData = @{
    username = "testuser"
    password = "testpass123"
} | ConvertTo-Json

try {
    $response = Invoke-WebRequest -Uri "http://localhost:8080/api/v1/auth/login" -Method Post -Headers $headers -Body $loginData -TimeoutSec 5
    Write-Host "Login successful: $($response.Content)"
    
    # Extract token from response
    $token = ($response.Content | ConvertFrom-Json).token
    if ($token) {
        $headers["Authorization"] = "Bearer $token"
    }
} catch {
    Write-Host "Login failed: $($_.Exception.Message)"
}

# Test token validation endpoint
Write-Host "`nTesting token validation endpoint..."
try {
    $response = Invoke-WebRequest -Uri "http://localhost:8080/api/v1/auth/validate" -Method Get -Headers $headers -TimeoutSec 5
    Write-Host "Token validation successful: $($response.Content)"
} catch {
    Write-Host "Token validation failed: $($_.Exception.Message)"
} 