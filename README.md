# drag
A barebones tool to aggregate Go style duration strings via stdin.

```sh
echo "1ns 2us 4ms 8s 16m 32h" | tr " " "\n" | drag
```
