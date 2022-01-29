# remove path variable, delete folder
$installDir = $env:Home + "\clitaskmanager"
[Environment]::SetEnvironmentVariable("Path", $env:Path.Replace(";$installDir",""), [System.EnvironmentVariableTarget]::Machine)

if (Test-Path -Path $installDir) {
    Remove-Item -Path $installDir -Force -Recurse
}
