Push-Location $PSScriptRoot\deck
try {
    go test -v
} finally {
    Pop-Location
}

powershell -ExecutionPolicy Bypass -Command "$PSScriptRoot\buildDoc.ps1"