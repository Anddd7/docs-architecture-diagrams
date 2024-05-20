# D2 library

## VSM

It contains a color theme to build a vsm diagram with d2. You can use import to compose multiple 
stages vsm into an "overview" vsm.

```plain
...@vsm

main: {
  ...@vsm-demo-main
}

branch: {
  ...@vsm-demo-branch
}
branch2: {
  ...@vsm-demo-branch
}

main -> branch
main -> branch2
```
