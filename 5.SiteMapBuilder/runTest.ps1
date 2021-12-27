go test .\... -v -coverprofile cover.out

if ($LASTEXITCODE -eq 1) {
    return $false
}

return $true