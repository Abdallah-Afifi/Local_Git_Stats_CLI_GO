To use `gogitlocalstats`, follow these steps:

1. **Install the tool:**
    ```sh
    go install
    ```

2. **Scan a folder and its subdirectories for repositories:**
    ```sh
    gogitlocalstats -add /path/to/folder
    ```

3. **Generate a CLI stats graph for a specific email:**
    ```sh
    gogitlocalstats -email your@email.com
    ```
    This command will generate a graph representing the last 6 months of activity for the specified email. 

You can configure the default email in `main.go`, allowing you to run `gogitlocalstats` without parameters. Passing an email as a parameter also enables scanning repositories for collaborators' activity.