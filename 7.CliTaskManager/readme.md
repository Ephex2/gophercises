# CLI Task Manager - Gophercise 7

The purpose of this mini-project is to have us working with bbolt (boltdb) as well as cobra, to create a command-line TODO list that stores data locally.

## Installation

I wrote a quick install script for those running Windows - the install.ps1 script should build the task executable, install it into a folder called clitaskmanager in your $env:Home directory, and also adds the executable to the PATH. Uninstall.ps1 will remove it from the PATH as well as delete the folder that was created. The scripts should be run with PowerShell as administrator, no arguments necessary:
```powershell
# From project directory (cd into it if not already in it)
.\install.ps1
# OR
.\uninstall.ps1
```
**NOTE**: You must run PowerShell as administrator, otherwise the install will fail.

## Commands
Here are the commands available to the task executable:

- add: Adds a task to the task list. The task list is maintained in a local instance of FireboltDB. Accepts a set of strings as an argument.
- clear: Clear is used to remove all tasks from the task list. It should also crate a new, empty task list for future use.
- do: Removes a task from the task list. If no task is found with the same name, will return a message saying that the task could not be found.
- list: Lists all tasks currently on the task list. Does not accept any arguments.
- completion:  Automatically provided by cobra. Generate the autocompletion script for the specified shell

## Usage
```
> task add test
Adding the following task to the task list: test
> task list
test
> task do test
Removing the following task from task list: test
> task add test1
Adding the following task to the task list: test1
> task add test2
Adding the following task to the task list: test2
> task list
test1
test2
> task clear
Clear called - removing all tasks from task list.
> task list
```
