# D2 library

## VSM

It contains a color theme to build a vsm diagram with d2. You can use import to compose multiple 
stages vsm into an "overview" vsm.

```plain
# import styles 
...@vsm

# import a vsm as a block
main: {
  ...@vsm-demo-main
}

# import a vsm as a block
branch: {
  ...@vsm-demo-branch
}

# connect 
main.stage1 -> branch.stage
```

## 