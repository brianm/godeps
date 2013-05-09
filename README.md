# Go Dependency Helper

Tool which looks at a DEPS file in current directory and recursively
checks out the right versions of thing. 

It follows the same rules as "go get" except that it has versions. A
DEPS file looks like

```
github.com/foo/bar         v0.1.2
git.example.com/hello/     master      
foo.com/repo/blah/bling    bling-1.2.3  svn
```


