# Append auto-generated go documentation to readme.md
[string[]]$docString = go doc -all $PSScriptRoot\deck
$docString = $docString.Split("`n")

[string[]]$baseDoc = Get-Content $PSScriptRoot\readme.base.md

$baseDoc += @"
Automatically generated documentation:
``````
"@

$docString | ForEach-Object {$baseDoc += $_}

$baseDoc += @"
``````
"@

$baseDoc > $PSScriptRoot\readme.md