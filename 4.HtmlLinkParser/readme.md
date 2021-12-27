# Html Link Parser - Gophercise 4

This project is meant to parse an html file and output all of the links found within the file in the following format:
```json
[
    {
        "Href": "https://www.twitter.com/joncalhoun",
        "Text": "Check me out on twitter"
    },
    {
        "Href": "https://github.com/gophercises",
        "Text": "Gophercises is on Github!"
    }
]
```
<br></br>
## Optional Flags
Here's an overview of the different flags available for main.go:
- -path: This flag is used to provide the path towards an html file that must be parsed.

## Note
The fifth exercise makes use of the linkparser package developped in this exercise.
	