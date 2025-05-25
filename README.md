Compile the program:

```bash
go build -o sonar
```

Create a shell function in your .bashrc or .zshrc:

```bash
sonar() {
    target_dir="$(<target_dir> "$@" 2>/dev/null)"

    exit_code=$?

	if [[ -n "$target_dir" && "$exit_code" -eq 100 ]] ; then
		cd "$target_dir"  || echo "Failed to cd into $target_dir"
    else 
       echo "$target_dir"
	fi

}

```
Source your shell configuration:

```bash
source ~/.bashrc  # or ~/.zshrc
```
Now you can run:

```bash
sonar
```

