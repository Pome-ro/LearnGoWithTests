# Notes

Following allong with: https://quii.gitbook.io/learn-go-with-tests/

## 11-22-23

* gopls was angry that I created a new module inside of the integers folder. Typically your working directory should be the root of your module where your go.mod file lives.
* However, you can work on multiple modules at once by creating a go.work file.
* You can create this file using the command `go work init` and then `go work use ./mod1/ ./mod2/`, replacing mod1 and mod2 with the paths of your modules.
* Relevent Doc: <https://github.com/golang/tools/blob/master/gopls/doc/workspace.md>
* More likely what needs to change though is I should have a single go.mod file at the root of this directory, that way it is all one module. I'll fix this now.
* that's correced now and the workspace is happy again.
