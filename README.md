Compile the program:

```bash
go build -o gocd
```

Create a shell function in your .bashrc or .zshrc:

```bash
sonar() {
	target_dir="$(./sonar 2>/dev/null)"
	if [ -n "$target_dir" ]; then
		cd "$target_dir"
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

