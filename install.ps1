# Download URL
$binaryUrl = "https://github.com/Tao-Toxicboy19/imotif-tools/releases/latest/download/itgc.exe"

# Default install path (local user space, no admin needed)
$defaultInstallPath = "$env:USERPROFILE\.itgc"
Write-Host "Default install path is: $defaultInstallPath"

# Ask for custom path
$customPath = Read-Host "Enter custom install path or press [Enter] to use default"
$installPath = if ([string]::IsNullOrWhiteSpace($customPath)) { $defaultInstallPath } else { $customPath }

# Ensure directory exists
New-Item -ItemType Directory -Force -Path $installPath | Out-Null

# Path to binary
$binaryPath = Join-Path $installPath "itgc.exe"

# Download the binary
Write-Host "Downloading itgc.exe to $binaryPath..."
Invoke-WebRequest -Uri $binaryUrl -OutFile $binaryPath

# Add installPath to PATH if not already in it
$envPath = [Environment]::GetEnvironmentVariable("Path", "User")

if (-not $envPath.Split(";") -contains $installPath) {
    Write-Host "Adding $installPath to PATH (User scope)..."
    [Environment]::SetEnvironmentVariable("Path", "$envPath;$installPath", "User")
    $pathUpdated = $true
} else {
    $pathUpdated = $false
}

Write-Host "itgc installed at: $binaryPath"

if ($pathUpdated) {
    Write-Host "You may need to restart your terminal or log out/in for PATH to update."
} else {
    Write-Host "You can now run 'itgc' from any terminal!"
}
