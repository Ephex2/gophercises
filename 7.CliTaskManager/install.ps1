go build .\main.go
Rename-Item -Path "$PSScriptRoot\main.exe" -NewName "task.exe" -Force 

$installDir = $env:home + "\clitaskmanager"

if (!(Test-Path -Path $installDir)) {
    mkdir $installDir | Out-Null
}

Move-Item -Path "$PSScriptRoot\task.exe" -destination $installDir -Force

# Create New Path variable 
[Environment]::SetEnvironmentVariable("Path", "$($env:Path);$installDir", [System.EnvironmentVariableTarget]::Machine)