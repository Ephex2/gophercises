# Choose your Own Adventure - Gophercise 3

This project is meant to be a choose your own adventure which is based on a json file which much be fed to main.go.
By default, this will start a local webserver that you can navigate to by browsing to http://localhost:8080 .
The default adventures are stored in the gopher.json file in the same directory as this readme file. If another one should be specified, the -filepath flag should be used.
<br></br>
## Optional Flags
Here's an overview of the different flags available for main.go:
- -cli: This flag indicates that you would like to navigate through the adventures as a command line game rather than a browser based one. Text will be output to the terminal, and you will be offered choices near the end. You can make a choice by typing in the choice (spelling and casing matters).

- -filepath: Specifies the path to a json file that contains adventures that can make up a choose your own adventure story. The format should match the existing json called gopher.json found [here](.\gopher.json). Note that the strings that name each object are considered the **ArcTitles**.

- -defaultarc: Specifies the first arc to start on for the choose your own adventure story. In the [sample](.\gopher.json), the default arc is intro. Because of this, if the -defaultarc flag is not specified, it will have the default value of intro.

If the program is run with no flags, it will start a webserver on port 8080 and use the default story found in [gopher.json](.\gopher.json)

## Notes on choose your own adventure json files
There isn't a jsonschema used in this project, but using one could potentially help with validating new CYOA json files.
Keep in mind:
- Each arc should have at least one option pointing to another arc.
- The story elements describe what happened that arc.
- If no other arc has an option to point to an arc that you are writing, it will be unavailable in the CYOA game.