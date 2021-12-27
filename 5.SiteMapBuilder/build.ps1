param (
    [string]$ExecutablePath = "./siteMapBuilder.exe"
)

$passed = .\runTest.ps1

if ($passed) {
    # TODO: Publish code coverage report to ???
    # TODO: Write actual doc strings and run go doc to auto-generate documentation. Publish documentation.
    go build -o "$ExecutablePath"
}

