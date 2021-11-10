# Quiz me exercise
- Command-line tool to generate quizzes based on csv input files, defaults to problems.csv in this folder.
    - CSV files must not have headers. First column is a question, whereas the second column is an answer.
    - Answers should be integers.
        - This could be revisited, but when implementing argument validation this seemed to be the simplest way.

- Shuffle flag on the tool allows questions to be shuffled.

- Timer flag allows to set the quiz timer. Default is 30 seconds.

- Filepath flag allows you to set the path to a csv file with problems in it. Note that relative pathes will be calculated relative to the caller of the command-line tool.

<br></br>
# Examples
- From same directory, no parameters:
    - go run .\quizgame.go 

- From one directory above, many paramters:
    - go run .\1.QuizGame\quizgame.go -timer 100 -shuffle -filepath .\1.QuizGame\problems.csv